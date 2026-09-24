package db

import (
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var fuzzFields = []string{"id", "a", "b"}

// few values so there are lots of duplicates, "a" is also the name of a field so it matches the header row
var fuzzValues = []string{"x", "y", "", "a"}

type fuzzRow struct {
	fields  []string
	deleted bool
}

// fuzzInput hands out the fuzz data one byte at a time, returning 0 once it runs out
type fuzzInput struct {
	data []byte
}

func (in *fuzzInput) next() int {
	if len(in.data) == 0 {
		return 0
	}
	b := in.data[0]
	in.data = in.data[1:]
	return int(b)
}

func (in *fuzzInput) choose(options []string) string {
	return options[in.next()%len(options)]
}

// FuzzTable creates a random csv file, opens it as a Table with random indexed fields and runs a
// random sequence of Insert and Remove calls. After every call FindFirst is checked for every field
// and value against a model of the rows. Finally the table is closed and the file is checked to
// contain exactly the rows that were not removed.
func FuzzTable(f *testing.F) {
	f.Add([]byte{3, 0, 1, 1, 2, 3, 0, 1, 7, 0, 0, 1, 1, 1, 2, 1, 0, 1})
	f.Add([]byte{2, 0, 0, 1, 1, 1, 1, 0, 2, 1, 1, 1, 0, 1, 0, 1})
	f.Add([]byte{0, 0, 1})

	f.Fuzz(func(t *testing.T, data []byte) {
		in := &fuzzInput{data: data}
		fileName := filepath.Join(t.TempDir(), "fuzz.csv")

		// build the starting file, sometimes without a newline at the end
		content := strings.Join(fuzzFields, ",") + "\n"
		rows := []*fuzzRow{}

		numRows := in.next() % 6
		for i := 0; i < numRows; i++ {
			fields := []string{strconv.Itoa(i), in.choose(fuzzValues), in.choose(fuzzValues)}
			rows = append(rows, &fuzzRow{fields: fields})
			content += strings.Join(fields, ",") + "\n"
		}

		if in.next()%2 == 0 {
			content = strings.TrimSuffix(content, "\n")
		}

		if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}

		indexMask := in.next()
		indexed := map[string]bool{}
		fieldsToIndex := []string{}
		for i, field := range fuzzFields {
			if indexMask&(1<<i) != 0 {
				indexed[field] = true
				fieldsToIndex = append(fieldsToIndex, field)
			}
		}

		table, err := NewTable(fileName, fieldsToIndex)
		if err != nil {
			t.Fatalf("NewTable: %v\nfile: %q", err, content)
		}

		checkAllFinds(t, table, rows, indexed)

		for op := 0; len(in.data) > 0 && op < 100; op++ {
			if in.next()%2 == 0 {
				record := []string{strconv.Itoa(len(rows)), in.choose(fuzzValues), in.choose(fuzzValues)}

				if err := table.Insert(record); err != nil {
					t.Fatalf("Insert(%q): %v", record, err)
				}
				rows = append(rows, &fuzzRow{fields: record})
			} else {
				field, key := chooseRemove(in, rows, indexed)
				checkRemove(t, table, rows, indexed, field, key)
			}

			checkAllFinds(t, table, rows, indexed)
		}

		if err := table.Close(); err != nil {
			t.Fatalf("Close: %v", err)
		}

		want := strings.Join(fuzzFields, ",") + "\n"
		for _, row := range rows {
			if !row.deleted {
				want += strings.Join(row.fields, ",") + "\n"
			}
		}

		got, err := os.ReadFile(fileName)
		if err != nil {
			t.Fatalf("reading file after Close: %v", err)
		}
		if string(got) != want {
			t.Fatalf("file after Close\ngot:  %q\nwant: %q", got, want)
		}

		if _, err := os.Stat(fileName + "_temp"); err == nil {
			t.Fatal("Close left the temporary file behind")
		}
	})
}

// chooseRemove picks what to remove by. Only id (which is unique) or unindexed fields
// (which remove the first match in the file) are used, so the model knows which row is removed.
func chooseRemove(in *fuzzInput, rows []*fuzzRow, indexed map[string]bool) (string, string) {
	field := in.choose(fuzzFields)
	if indexed[field] {
		field = "id"
	}

	if field == "id" {
		choice := in.next()
		if len(rows) == 0 || choice%4 == 0 {
			return field, "missing"
		}
		return field, rows[choice%len(rows)].fields[0]
	}
	return field, in.choose(fuzzValues)
}

func fieldIndex(field string) int {
	for i, v := range fuzzFields {
		if v == field {
			return i
		}
	}
	return -1
}

// liveRowsWith returns the rows that have not been deleted and have key in the field, in file order
func liveRowsWith(rows []*fuzzRow, field string, key string) []*fuzzRow {
	column := fieldIndex(field)
	matching := []*fuzzRow{}
	for _, row := range rows {
		if !row.deleted && row.fields[column] == key {
			matching = append(matching, row)
		}
	}
	return matching
}

// checkAllFinds runs FindFirst for every field and every value that could be stored in it
func checkAllFinds(t *testing.T, table *Table, rows []*fuzzRow, indexed map[string]bool) {
	t.Helper()

	for _, field := range fuzzFields {
		keys := fuzzValues
		if field == "id" {
			keys = []string{"missing"}
			for _, row := range rows {
				keys = append(keys, row.fields[0])
			}
		}

		for _, key := range keys {
			checkFindFirst(t, table, rows, indexed, field, key)
		}
	}
}

func checkFindFirst(t *testing.T, table *Table, rows []*fuzzRow, indexed map[string]bool, field string, key string) {
	t.Helper()

	got, err := table.FindFirst(field, key)
	matching := liveRowsWith(rows, field, key)

	if len(matching) == 0 {
		if indexed[field] {
			if err == nil {
				t.Fatalf("FindFirst(%q, %q) on an indexed field found %q but no row matches", field, key, got)
			}
		} else if got != nil || err != nil {
			t.Fatalf("FindFirst(%q, %q) on an unindexed field returned (%q, %v) but no row matches", field, key, got, err)
		}
		return
	}

	if err != nil {
		t.Fatalf("FindFirst(%q, %q): %v", field, key, err)
	}

	if !indexed[field] {
		// an unindexed search reads the file in order so must return the first matching row
		if !reflect.DeepEqual(got, matching[0].fields) {
			t.Fatalf("FindFirst(%q, %q) = %q, want the first matching row %q", field, key, got, matching[0].fields)
		}
		return
	}

	// with duplicate keys the index may return any of the matching rows
	for _, row := range matching {
		if reflect.DeepEqual(got, row.fields) {
			return
		}
	}
	t.Fatalf("FindFirst(%q, %q) = %q which is not a matching row that has not been removed", field, key, got)
}

// checkRemove removes by field and key and marks the row that should have been removed as deleted in the model.
// Whether the table removed the right row is checked afterwards by checkAllFinds and the file written by Close.
func checkRemove(t *testing.T, table *Table, rows []*fuzzRow, indexed map[string]bool, field string, key string) {
	t.Helper()

	err := table.Remove(field, key)
	matching := liveRowsWith(rows, field, key)

	if len(matching) == 0 {
		if indexed[field] && err == nil {
			t.Fatalf("Remove(%q, %q) on an indexed field returned no error but no row matches", field, key)
		}
		if !indexed[field] && err != nil {
			t.Fatalf("Remove(%q, %q): %v", field, key, err)
		}
		return
	}

	if err != nil {
		t.Fatalf("Remove(%q, %q): %v", field, key, err)
	}

	matching[0].deleted = true
}
