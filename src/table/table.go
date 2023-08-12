package db

import (
	"errors"
	"io"

	"github.com/Max-Sepp/csv-indexing/src/internal/btree"
	"github.com/Max-Sepp/csv-indexing/src/internal/simplecsv"
)

type Table struct {
	handler       *simplecsv.CsvHandler
	fields        []string
	indexedFields []int // holds the indexs of the fields that are indexed
	btrees        []*btree.Btree
}

func NewTable(fileName string, fieldsToIndex []string) (*Table, error) {
	handler, err := simplecsv.NewHandler(fileName)

	if err != nil {
		return nil, err
	}

	fields, err := handler.Read()

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

	// generate btrees
	btrees := make([]*btree.Btree, 0, len(fieldsToIndex))

	for i := 0; i < len(fieldsToIndex); i++ {
		btrees = append(btrees, btree.New(5))
	}

	for {
		offset := handler.Offset
		record, err := handler.Read()

		if err == io.EOF {
			handler.WriteOffset = offset
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
		handler:       handler,
		fields:        fields,
		indexedFields: indexedFields,
		btrees:        btrees,
	}, err
}

// Find first returns the first element that matches the key and the field.
// if no record is found returns a nil string
func (table *Table) FindFirst(field string, key string) ([]string, error) {
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

	data, err := table.handler.ReadLineAt(recordByteOffset)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// fieldIndex is the index of the item we are searching for in the record i.e. is it the first second or third field in the record
// returns nil if nothing is found
func (table *Table) findFirstUnindexed(fieldIndex int, key string) ([]string, error) {
	table.handler.Reset()

	for {
		record, err := table.handler.Read()

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

	offset := table.handler.WriteOffset

	table.handler.Append(record)

	for i, v := range table.indexedFields {
		table.btrees[i].Insert(record[v], int64(offset))
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
