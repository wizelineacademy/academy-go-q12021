package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("First Deliverable")

	csvFile, err := os.Open("dataFile.csv")
	check(err)

	count := 0
	csvReader := csv.NewReader(csvFile)
	for {
		dataRow, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			fmt.Printf("Done ... %v rows were read.", count)
			break
		}
		check(err)
		fmt.Printf("%v\n", dataRow)
		count++
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
