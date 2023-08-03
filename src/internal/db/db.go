package db

import (
	"encoding/csv"
	"errors"
	"os"
)

type db struct {
	file      *os.File
	csvReader *csv.Reader
}

// creates a instance of db
func NewDb(name string) (*db, error) {
	file, err := os.Open(name)

	if err != nil {
		return &db{}, err
	}

	csvReader := csv.NewReader(file)

	db := &db{
		file:      file,
		csvReader: csvReader,
	}

	return db, nil
}

func (db *db) GetDataAt(offset int64) ([]string, error) {
	f := db.file
	r := db.csvReader

	ret, err := f.Seek(offset, 0)

	if err != nil {
		return nil, err
	}

	if ret != offset {
		return nil, errors.New("offset was not set correctly")
	}

	record, err := r.Read()

	if err != nil {
		return nil, err
	}

	return record, nil
}
