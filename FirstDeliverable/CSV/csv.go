package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Item struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	ItemType string `json:"item_type"`
}

func ReadCSV() ([][]string, error) {
	// Open the file
	csvfile, err := os.Open("../FirstDeliverable/CSV/result.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()

	// Parse the file
	reader := csv.NewReader(csvfile)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return records, err
}

func RetrieveFromCSV(itemID string) string {
	var items []Item
	records, err := ReadCSV()
	// Iterate through the records
	for _, record := range records {
		// Read each record from csv
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error getting the records: ", err)
		}
		if itemID != "" {
			if itemID == record[0] {
				itemRec := Item{ItemID: record[0], ItemName: record[1], ItemType: record[2]}
				items = append(items, itemRec)
			}
		} else {
			itemRec := Item{ItemID: record[0], ItemName: record[1], ItemType: record[2]}
			items = append(items, itemRec)
		}
	}
	jsonData, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(jsonData)
}
