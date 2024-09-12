package db

import (
	"errors"
	"io"
	"log"
	"os"
	"sync"

	"github.com/Max-Sepp/csv-indexing/internal/btree"
	"github.com/Max-Sepp/csv-indexing/internal/simplecsv"
)

type Table struct {
	csvHandler    *simplecsv.CsvHandler
	fields        []string
	indexedFields []int // holds the indexs of the fields that are indexed
	btrees        []*btree.Btree
	deleted       map[int]bool // the byte offset of deleted rows
	mutex         sync.Mutex
}

func NewTable(fileName string, fieldsToIndex []string) (*Table, error) {
	csvHandler, err := simplecsv.NewHandler(fileName)

	if err != nil {
		return nil, err
	}

	fields, err := csvHandler.Read()

	if err == io.EOF {
		return nil, errors.New("empty csv file")
	}

	indexedFields := make([]int, 0, len(fieldsToIndex))

	for _, v := range fieldsToIndex {
		index, err := findInSlice(fields, v)

		if err != nil {
			return nil, errors.New("cannot find the element in the fields")
		}

		indexedFields = append(indexedFields, index)
	}

	btrees := generateBtrees(len(fieldsToIndex))

	for {
		offset := csvHandler.Offset
		record, err := csvHandler.Read()

		if err == io.EOF {
			csvHandler.WriteOffset = offset
			break
		}

		if err != nil {
			return nil, err
		}

		for i, v := range indexedFields {
			btrees[i].Insert(record[v], int64(offset))
		}
	}

	return &Table{
		csvHandler:    csvHandler,
		fields:        fields,
		indexedFields: indexedFields,
		btrees:        btrees,
	}, err
}

// Find first returns the first element that matches the key and the field.
// if no record is found returns a nil string
func (table *Table) FindFirst(field string, key string) ([]string, error) {
	table.mutex.Lock()

	defer table.mutex.Unlock()

	i := 0

	for field != table.fields[i] && i < len(table.fields) {
		i++
	}
	if i == len(table.fields) {
		return nil, errors.New("field is not a field in the table")
	}

	indexOfBtree := 0
	for i != table.indexedFields[indexOfBtree] && indexOfBtree < len(table.indexedFields) {
		indexOfBtree++
	}
	if indexOfBtree == len(table.indexedFields) {
		return table.findFirstUnindexed(i, key)
	}
	return table.findIndexed(indexOfBtree, key)

}

func (table *Table) findIndexed(indexOfBtree int, key string) ([]string, error) {
	found, _, recordByteOffset := table.btrees[indexOfBtree].Find(key)

	if !found {
		return nil, errors.New("record could not be found with that key")
	}

	data, err := table.csvHandler.ReadLineAt(recordByteOffset)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// fieldIndex is the index of the item we are searching for in the record i.e. is it the first second or third field in the record
// returns nil if nothing is found
func (table *Table) findFirstUnindexed(fieldIndex int, key string) ([]string, error) {
	table.csvHandler.ResetReaderOffset()

	for {
		record, err := table.csvHandler.Read()

		if err == io.EOF {
			// we have reached the end of the file and not found the item in the field
			return nil, nil
		}

		if err != nil {
			return nil, err
		}

		if record[fieldIndex] == key {
			return record, nil
		}
	}
}

func (table *Table) Insert(record []string) error {
	table.mutex.Lock()

	defer table.mutex.Unlock()

	offset := table.csvHandler.WriteOffset

	table.csvHandler.Append(record)

	for i, v := range table.indexedFields {
		table.btrees[i].Insert(record[v], int64(offset))
	}

	return nil
}

func (table *Table) Remove(field string, key string) error {
	table.mutex.Lock()

	defer table.mutex.Unlock()

	// find the byteoffset to delete

	i := 0

	for field != table.fields[i] && i < len(table.fields) {
		i++
	}
	if i == len(table.fields) {
		return errors.New("field is not a field in the table")
	}

	indexOfBtree := 0
	for i != table.indexedFields[indexOfBtree] && indexOfBtree < len(table.indexedFields) {
		indexOfBtree++
	}

	offset := 0

	if indexOfBtree == len(table.indexedFields) {
		fieldIndex := i

		table.csvHandler.ResetReaderOffset()

		for {
			offset = table.csvHandler.Offset

			record, err := table.csvHandler.Read()

			if err == io.EOF {
				// we have reached the end of the file and not found the item in the field
				return nil
			}

			if err != nil {
				return err
			}

			if record[fieldIndex] == key {
				break
			}
		}
	} else {
		returnedOffset, err := table.btrees[indexOfBtree].Delete(key)
		if err != nil {
			log.Fatal(err)
		}

		offset = int(returnedOffset)
	}

	table.deleted[offset] = true

	return nil
}

func (table *Table) Close() error {
	fileName := table.csvHandler.FileName

	tempFile := fileName + "_temp"

	table.csvHandler.Close()

	if err := os.Rename(fileName, tempFile); err != nil {
		// highly unlikely to ever occur
		return errors.New("Failed to rename: " + err.Error())
	}

	if err := os.Truncate(fileName, 0); err != nil {
		return errors.New("Failed to truncate: " + err.Error())
	}

	writeHandler, err := simplecsv.NewHandler(fileName)
	if err != nil {
		return errors.New("Failed to open file: " + err.Error())
	}
	defer writeHandler.Close()

	readHandler, err := simplecsv.NewHandler(tempFile)
	if err != nil {
		return errors.New("Failed to open file: " + err.Error())
	}
	defer readHandler.Close()

	for {
		offset := readHandler.Offset
		record, err := readHandler.Read()

		if err == io.EOF {
			break
		}

		if _, present := table.deleted[offset]; !present {
			writeHandler.Append(record)
		}
	}
	return nil
}

func findInSlice[T comparable](slice []T, value T) (int, error) {
	for i, v := range slice {
		if v == value {
			return i, nil
		}
	}

	return -1, errors.New("can not find item in the slice")
}

func generateBtrees(numBtrees int) []*btree.Btree {
	btrees := make([]*btree.Btree, 0, numBtrees)

	for i := 0; i < numBtrees; i++ {
		btrees = append(btrees, btree.New(5))
	}

	return btrees
}
