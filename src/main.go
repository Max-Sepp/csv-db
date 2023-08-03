package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test_data.csv")

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {
		fmt.Println(r.InputOffset())
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}

	fmt.Println("--------- Testing accessing specific record ----------")

	ret, err := f.Seek(122, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ret)

	ret, err = f.Seek(86, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ret)

	record, err := r.Read()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(record)

	err = f.Close()

	if err != nil {
		log.Fatal(err)
	}

}
