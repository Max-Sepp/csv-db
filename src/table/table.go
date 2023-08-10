package db

import (
	"errors"
	"io"

	"github.com/Max-Sepp/csv-indexing/src/internal/btree"
	"github.com/Max-Sepp/csv-indexing/src/internal/simplecsv"
)

type Table struct {
	reader        *simplecsv.CsvReader
	fields        []string
	indexedFields []int // holds the indexs of the fields that are indexed
	btrees        []*btree.Btree
}

func NewTable(fileName string, fieldsToIndex []string) (*Table, error) {
	reader, err := simplecsv.NewReader(fileName)

	if err != nil {
		return nil, err
	}

	fields, err := reader.Read()

	if err == io.EOF {
		return nil, errors.New("start of csv should contain field names of table")
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
		offset := reader.Offset
		record, err := reader.Read()

		if err == io.EOF {
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
		reader:        reader,
		fields:        fields,
		indexedFields: indexedFields,
		btrees:        btrees,
	}, err
}

func (table *Table) FindByUnique(field string, key string) ([]string, error) {
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
		return nil, errors.New("field is not indexed in the table")
	}

	found, _, recordByteOffset := table.btrees[indexOfBtree].Find(key)

	if !found {
		return nil, errors.New("record could not be found with that key")
	}

	data, err := table.reader.ReadLineAt(recordByteOffset)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func findInSlice[T comparable](slice []T, value T) (int, error) {
	for i, v := range slice {
		if v == value {
			return i, nil
		}
	}

	return -1, errors.New("can not find item in the slice")
}
