package simplecsv_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/Max-Sepp/csv-indexing/src/internal/simplecsv"
)

func TestAppend(t *testing.T) {
	const fileName = "test_data_write_append.csv"

	// setting up the test
	data := []byte("id,first_name,second_name\n6,Don,Sampson\n5,Rufus,Arias\n23,Joan,Morgan\n19,Joesph,Summers\n3,Despina,Coppola\n13,Karina,Everett\n17,Helen,Holman\n7,Shonta,Davis\n10,Howard,Elizondo\n9,John,Fennell\n")

	err := os.WriteFile(fileName, data, 0755)

	defer os.Remove(fileName)

	if err != nil {
		t.Fatal(err)
	}

	// actual test
	w, err := simplecsv.NewHandler(fileName)

	if err != nil {
		t.Fatal(err)
	}

	w.Append([]string{"2", "Bob", "Hamilton"})

	// checking
	content, err := os.ReadFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(content, []byte("id,first_name,second_name\n6,Don,Sampson\n5,Rufus,Arias\n23,Joan,Morgan\n19,Joesph,Summers\n3,Despina,Coppola\n13,Karina,Everett\n17,Helen,Holman\n7,Shonta,Davis\n10,Howard,Elizondo\n9,John,Fennell\n2,Bob,Hamilton\n")) {
		t.Error("did not write correctly")
	}

}
