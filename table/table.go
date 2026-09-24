package db

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/Max-Sepp/csv-indexing/internal/btree"
	"github.com/Max-Sepp/csv-indexing/internal/simplecsv"
)

type Table struct {
	// This Table implementation is intentionally simple for demonstration purposes.
	// A proper database would use more sophisticated file handling.
	// For example using memory-mapped files or buffering writes to reduce disk I/O.
	csvHandler    *simplecsv.CsvHandler
	fields        []string
	indexedFields []int // holds the indexs of the fields that are indexed
	btrees        []*btree.Btree
	deleted       map[int]bool // the byte offset of deleted rows
	mutex         sync.Mutex
}

func NewTable(fileName string, fieldsToIndex []string) (table *Table, err error) {
	csvHandler, err := simplecsv.NewHandler(fileName)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			csvHandler.Close()
		}
	}()

	// makes sure rows inserted later start on their own line
	if err := csvHandler.EnsureTrailingNewline(); err != nil {
		return nil, err
	}
	if err := csvHandler.ResetReaderOffset(); err != nil {
		return nil, err
	}

	fields, err := csvHandler.Read()

	if err == io.EOF {
		return nil, errors.New("empty csv file")
	}

	if err != nil {
		return nil, err
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

		if len(record) != len(fields) {
			return nil, fmt.Errorf("row at byte offset %d has %d fields but the header has %d", offset, len(record), len(fields))
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
		deleted:       map[int]bool{},
	}, nil
}

// Find first returns the first element that matches the key and the field.
// if no record is found returns a nil string
func (table *Table) FindFirst(field string, key string) ([]string, error) {
	table.mutex.Lock()

	defer table.mutex.Unlock()

	i := 0

	for i < len(table.fields) && field != table.fields[i] {
		i++
	}
	if i == len(table.fields) {
		return nil, errors.New("field is not a field in the table")
	}

	indexOfBtree := 0
	for indexOfBtree < len(table.indexedFields) && i != table.indexedFields[indexOfBtree] {
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
	_, record, err := table.scanForRecord(fieldIndex, key)
	return record, err
}

// scanForRecord reads through the file for the first row that has not been deleted and has key in the field fieldIndex.
// returns the byte offset of the row and the row, the row is nil if nothing is found
func (table *Table) scanForRecord(fieldIndex int, key string) (int, []string, error) {
	if err := table.csvHandler.ResetReaderOffset(); err != nil {
		return 0, nil, err
	}

	// skip the header row
	if _, err := table.csvHandler.Read(); err != nil && err != io.EOF {
		return 0, nil, err
	}

	for {
		offset := table.csvHandler.Offset
		record, err := table.csvHandler.Read()

		if err == io.EOF {
			// we have reached the end of the file and not found the item in the field
			return 0, nil, nil
		}

		if err != nil {
			return 0, nil, err
		}

		if table.deleted[offset] {
			continue
		}

		if record[fieldIndex] == key {
			return offset, record, nil
		}
	}
}

func (table *Table) Insert(record []string) error {
	table.mutex.Lock()

	defer table.mutex.Unlock()

	if len(record) != len(table.fields) {
		return fmt.Errorf("record has %d fields but the table has %d", len(record), len(table.fields))
	}

	// the csv format used does not support quoting so these would corrupt the file
	for _, v := range record {
		if strings.ContainsAny(v, ",\n") {
			return errors.New("values can not contain commas or newlines")
		}
	}

	offset := table.csvHandler.WriteOffset

	if err := table.csvHandler.Append(record); err != nil {
		return err
	}

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

	for i < len(table.fields) && field != table.fields[i] {
		i++
	}
	if i == len(table.fields) {
		return errors.New("field is not a field in the table")
	}

	indexOfBtree := 0
	for indexOfBtree < len(table.indexedFields) && i != table.indexedFields[indexOfBtree] {
		indexOfBtree++
	}

	var offset int
	var record []string

	if indexOfBtree == len(table.indexedFields) {
		var err error
		offset, record, err = table.scanForRecord(i, key)

		if err != nil {
			return err
		}

		if record == nil {
			// we have reached the end of the file and not found the item in the field
			return nil
		}
	} else {
		found, _, recordByteOffset := table.btrees[indexOfBtree].Find(key)

		if !found {
			return errors.New("record could not be found with that key")
		}

		var err error
		record, err = table.csvHandler.ReadLineAt(recordByteOffset)

		if err != nil {
			return err
		}

		offset = int(recordByteOffset)
	}

	// remove the row from every index so it can not be found through another field
	for j, v := range table.indexedFields {
		if err := table.btrees[j].DeleteEntry(record[v], int64(offset)); err != nil {
			return err
		}
	}

	table.deleted[offset] = true

	return nil
}

func (table *Table) Close() error {
	table.mutex.Lock()

	defer table.mutex.Unlock()

	fileName := table.csvHandler.FileName

	tempFile := fileName + "_temp"

	table.csvHandler.Close()

	// the rows that are kept are written to a temporary file which then replaces the original,
	// so the original file is never left partially written
	if err := table.writeRowsNotDeleted(fileName, tempFile); err != nil {
		os.Remove(tempFile)
		return err
	}

	if err := os.Rename(tempFile, fileName); err != nil {
		return errors.New("Failed to rename: " + err.Error())
	}

	return nil
}

func (table *Table) writeRowsNotDeleted(fileName string, tempFile string) error {
	info, err := os.Stat(fileName)
	if err != nil {
		return err
	}

	if err := os.WriteFile(tempFile, nil, info.Mode().Perm()); err != nil {
		return errors.New("Failed to create file: " + err.Error())
	}

	writeHandler, err := simplecsv.NewHandler(tempFile)
	if err != nil {
		return errors.New("Failed to open file: " + err.Error())
	}
	defer writeHandler.Close()

	readHandler, err := simplecsv.NewHandler(fileName)
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

		if err != nil {
			return err
		}

		if !table.deleted[offset] {
			if err := writeHandler.Append(record); err != nil {
				return err
			}
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
