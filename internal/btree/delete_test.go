package btree

import (
	"log"
	"testing"
)

func TestDelete1(t *testing.T) {
	B := setupDeleteTest()

	_, err := B.Delete("1")

	if err != nil {
		t.Fatalf("deletion error: %s", err)
	}
}

func TestDelete2(t *testing.T) {
	B := setupDeleteTest()

	_, err := B.Delete("2")

	if err != nil {
		t.Fatalf("deletion error: %s", err)
	}
}
func TestDelete3(t *testing.T) {
	// should require combining of nodes
	B := setupDeleteTest()

	_, err := B.Delete("3")

	if err != nil {
		t.Fatalf("deletion error: %s", err)
	}
}
func TestDelete4(t *testing.T) {
	B := setupDeleteTest()

	_, err := B.Delete("4")

	if err != nil {
		t.Fatalf("deletion error: %s", err)
	}
}
func TestDelete5(t *testing.T) {
	B := setupDeleteTest()

	_, err := B.Delete("5")

	if err != nil {
		t.Fatalf("deletion error: %s", err)
	}
}
func TestDelete6(t *testing.T) {
	B := setupDeleteTest()

	_, err := B.Delete("6")

	if err != nil {
		t.Fatalf("deletion error: %s", err)
	}
}

func TestDelete12(t *testing.T) {
	B := setupDeleteTest()

	_, err := B.Delete("12")

	if err != nil {
		t.Fatalf("deletion error: %s", err)
	}
}

func TestDelete12then11(t *testing.T) {
	B := setupDeleteTest()

	B.Delete("12")
	B.Delete("11")
}

func TestDelete12then11then15(t *testing.T) {
	B := setupDeleteTest()

	B.Delete("12")
	B.Delete("11")
	_, err := B.Delete("15")

	if err != nil {
		log.Fatal(err)
	}

}

func TestDelete12then11then1(t *testing.T) {
	B := setupDeleteTest()

	B.Delete("12")
	B.Delete("11")
	_, err := B.Delete("1")

	if err != nil {
		log.Fatal(err)
	}

}

func TestDelete4then8(t *testing.T) {
	B := setupDeleteTest()

	B.Delete("4")
	_, err := B.Delete("8")

	if err != nil {
		log.Fatal(err)
	}

}

func TestDeleteLastElement(t *testing.T) {
	B := New(5)

	B.Insert("1", 0)

	_, err := B.Delete("1")

	if err != nil {
		log.Fatal(err)
	}
}

func setupDeleteTest() *Btree {
	B := New(5)

	data := []string{"14", "7", "12", "3", "4", "10", "19", "11", "1", "2", "20", "6", "17", "13", "9", "8", "18", "5", "15", "16"}

	for _, item := range data {
		B.Insert(item, 0)
	}

	return B
}
