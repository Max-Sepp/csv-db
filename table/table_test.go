package db

import (
	"bytes"
	"os"
	"testing"
)

func TestTableFindFirst(t *testing.T) {
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

func TestTableInsert(t *testing.T) {
	const fileName = "test_data_write_insert.csv"

	// setting up the test
	data := []byte("id,first_name,second_name\n6,Don,Sampson\n5,Rufus,Arias\n23,Joan,Morgan\n19,Joesph,Summers\n3,Despina,Coppola\n13,Karina,Everett\n17,Helen,Holman\n7,Shonta,Davis\n10,Howard,Elizondo\n9,John,Fennell\n")

	err := os.WriteFile(fileName, data, 0755)

	defer os.Remove(fileName)

	if err != nil {
		t.Fatal(err)
	}

	// actual test
	db, err := NewTable(fileName, []string{"id", "second_name"})

	if err != nil {
		t.Fatal(err)
	}

	db.Insert([]string{"2", "Bob", "Hamilton"})

	// checking
	content, err := os.ReadFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(content, []byte("id,first_name,second_name\n6,Don,Sampson\n5,Rufus,Arias\n23,Joan,Morgan\n19,Joesph,Summers\n3,Despina,Coppola\n13,Karina,Everett\n17,Helen,Holman\n7,Shonta,Davis\n10,Howard,Elizondo\n9,John,Fennell\n2,Bob,Hamilton\n")) {
		t.Error("did not write correctly")
	}

	out, err := db.FindFirst("id", "2")

	if err != nil {
		t.Fatal(err)
	}

	if !EqualStringSlice(out, []string{"2", "Bob", "Hamilton"}) {
		t.Error("Did not get correct output")
	}

	out, err = db.FindFirst("second_name", "Hamilton")

	if err != nil {
		t.Fatal(err)
	}

	if !EqualStringSlice(out, []string{"2", "Bob", "Hamilton"}) {
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
