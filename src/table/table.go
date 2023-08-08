package db

import (
	"github.com/Max-Sepp/csv-indexing/src/internal/btree"
	"github.com/Max-Sepp/csv-indexing/src/internal/simplecsv"
)

type table struct {
	reader        *simplecsv.CsvReader
	fields        []string
	indexedFields []string
	btrees        []*btree.Btree
}

func NewTable(fileName string, fieldsToIndex []string) (*table, error) {

}
