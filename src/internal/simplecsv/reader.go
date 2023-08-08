package simplecsv

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type CsvReader struct {
	file   *os.File
	reader *bufio.Reader
}

// NewReader retuns a new CsvReader. It returns the errors only of the opening of the file
func NewReader(name string) (*CsvReader, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)

	return &CsvReader{file: f, reader: r}, nil
}

// ReadLineAt returns the row that starts at the byte offset from the start and checks that the byte offset is at the start of a row
func (reader *CsvReader) ReadLineAt(offset int64) ([]string, error) {
	if offset != 0 {
		// check if previous char is endline so this is a complete line
		reader.setReadOffset(offset - 1)

		// check start of line
		prevLineChar, err := reader.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if prevLineChar != "\n" {
			return nil, errors.New("the byte offset does not begin with a row")
		}
	}

	ret, err := reader.reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	ret = strings.TrimSuffix(ret, "\n")

	return strings.Split(ret, ","), nil
}

// Reset sets the offset of the reader to 0
func (reader *CsvReader) Reset() {
	reader.setReadOffset(0)
}

// sets the read offset
func (reader *CsvReader) setReadOffset(offset int64) error {
	_, err := reader.file.Seek(offset, io.SeekStart)

	if err != nil {
		return err
	}

	reader.reader.Reset(reader.file)

	return nil
}

// Read returns the the next line
func (reader *CsvReader) Read() ([]string, error) {
	ret, err := reader.reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	ret = strings.TrimSuffix(ret, "\n")

	return strings.Split(ret, ","), nil
}