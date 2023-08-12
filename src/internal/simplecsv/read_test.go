package simplecsv_test

import (
	"io"
	"testing"

	"github.com/Max-Sepp/csv-indexing/src/internal/simplecsv"
)

func TestReadLineAtNormalData(t *testing.T) {
	r, err := simplecsv.NewHandler("test_data_read.csv")

	if err != nil {
		t.Error(err)
	}

	out, err := r.ReadLineAt(40)

	if err != nil {
		t.Error(err)
	}

	if !EqualStringSlice(out, []string{"5", "Rufus", "Arias"}) {
		t.Error("Failed: Reading at byte 40")
	}

	out, err = r.ReadLineAt(68)

	if err != nil {
		t.Error(err)
	}

	if !EqualStringSlice(out, []string{"19", "Joesph", "Summers"}) {
		t.Error("Failed: Reading at byte 68")
	}

	out, err = r.ReadLineAt(54)

	if err != nil {
		t.Error(err)
	}

	if !EqualStringSlice(out, []string{"3", "Joan", "Morgan"}) {
		t.Error("Failed: Reading at byte 54")
	}
}

func TestReadLineAtErroneousData(t *testing.T) {
	r, err := simplecsv.NewHandler("test_data_read.csv")

	if err != nil {
		t.Error(err)
	}

	_, err = r.ReadLineAt(41)

	if err == nil {
		t.Error("Failed does not return the byte offset error")
	}
}

func TestConsecutiveReads(t *testing.T) {
	// creating correct answer
	correctOutput := [][]string{
		{"id", "first_name", "second_name"},
		{"6", "Don", "Sampson"},
		{"5", "Rufus", "Arias"},
		{"3", "Joan", "Morgan"},
		{"19", "Joesph", "Summers"},
		{"3", "Despina", "Coppola"},
		{"13", "Karina", "Everett"},
		{"17", "Helen", "Holman"},
		{"7", "Shonta", "Davis"},
		{"10", "Howard", "Elizondo"},
		{"9", "John", "Fennell"},
	}

	correctOffset := []int{0, 26, 40, 54, 68, 86, 104, 122, 138, 153, 172, 187}

	output := [][]string{}

	r, err := simplecsv.NewHandler("test_data_read.csv")

	if err != nil {
		t.Error(err)
	}

	r.Reset()

	for i := 0; true; i++ {

		if r.Offset != correctOffset[i] {
			t.Error("Failed: did not have the correct offset")
		}

		out, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Error(out)
		}
		output = append(output, out)
	}

	if !Equal2DStringSlice(output, correctOutput) {
		t.Error("Failed: did not give the correct return value")
	}
}

func Equal2DStringSlice(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !EqualStringSlice(v, b[i]) {
			return false
		}
	}
	return true
}

func EqualStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
