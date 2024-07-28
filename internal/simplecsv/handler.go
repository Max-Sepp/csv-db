package simplecsv

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

type CsvHandler struct {
	file        *os.File
	reader      *bufio.Reader
	FileName    string
	Offset      int // the byte offset of the previous read if unknown / possible edge case will be -1
	WriteOffset int
}

func NewHandler(name string) (*CsvHandler, error) {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_RDWR, 0755) // uncertain what filemode to use
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)

	return &CsvHandler{
		file:        f,
		reader:      r,
		FileName:    name,
		Offset:      0,
		WriteOffset: 0,
	}, nil
}

// ReadLineAt returns the row that starts at the byte offset from the start and checks that the byte offset is at the start of a row
func (reader *CsvHandler) ReadLineAt(offset int64) ([]string, error) {
	if offset != 0 {
		// check if previous char is endline so this is a complete line
		reader.setReadOffset(offset - 1)

		// check start of line
		prevLineChar, err := reader.reader.ReadString('\n')
		if err != nil {
			reader.ResetReaderOffset()
			return nil, err
		}
		if prevLineChar != "\n" {
			reader.ResetReaderOffset()
			return nil, errors.New("the byte offset does not begin with a row")
		}
	}

	ret, err := reader.reader.ReadString('\n')

	if err != nil {
		reader.ResetReaderOffset()
		return nil, err
	}

	reader.Offset = int(offset) + len(ret)

	ret = strings.TrimSuffix(ret, "\n")

	return strings.Split(ret, ","), nil
}

func (reader *CsvHandler) ResetReaderOffset() {
	reader.setReadOffset(0)
}

func (reader *CsvHandler) setReadOffset(offset int64) error {
	_, err := reader.file.Seek(offset, io.SeekStart)

	if err != nil {
		reader.Offset = -1
		return err
	}

	reader.reader.Reset(reader.file)

	reader.Offset = int(offset)

	return nil
}

// Read returns the the next line
// if Read fails it does not reset Offset and leaves it unchanged
func (reader *CsvHandler) Read() ([]string, error) {
	ret, err := reader.reader.ReadString('\n')

	if err != nil && err != io.EOF {
		return nil, err
	}
	reader.Offset = reader.Offset + len(ret)

	ret = strings.TrimSuffix(ret, "\n")

	return strings.Split(ret, ","), err
}

func (handler *CsvHandler) Append(input []string) {
	data := []byte{}

	for i, v := range input {
		data = append(data, v...)
		if i != len(input)-1 {
			data = append(data, ',')
		}
	}

	data = append(data, '\n')

	handler.WriteOffset = handler.WriteOffset + len(data)

	_, err := handler.file.Write(data)

	if err != nil {
		log.Fatalf("Write operation failed: %v", err)
	}
}

func (handler *CsvHandler) Close() {
	handler.file.Close()
}
