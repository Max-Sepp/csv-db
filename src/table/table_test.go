package db

import (
	"testing"
)

func TestTable(t *testing.T) {
	db, err := NewTable("test_data.csv", []string{"id", "second_name"})

	if err != nil {
		t.Fatal(err)
	}

	out, err := db.FindFirst("id", "3")

	if err != nil {
		t.Fatal(err)
	}

	if !EqualStringSlice(out, []string{"3", "Despina", "Coppola"}) {
		t.Error("Did not get correct output")
	}

	out, err = db.FindFirst("second_name", "Coppola")

	if err != nil {
		t.Fatal(err)
	}

	if !EqualStringSlice(out, []string{"3", "Despina", "Coppola"}) {
		t.Error("Did not get correct output")
	}
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
