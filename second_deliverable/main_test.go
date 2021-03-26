package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"
)

var Items []Item = []Item{
	{Id: "1", Title: "Netflix"},
	{Id: "2", Title: "Dispney+"},
	{Id: "3", Title: "HBO Max"},
	{Id: "4", Title: "Paramount+"},
	{Id: "5", Title: "Universal+"},
}

func GetDataFromCSVFile(filePath string) ([][] string)  {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("\n- An error ocurred while reading the file. \n", err)
	} else {
		fmt.Println("\nSuccessfully Opened CSV file") 
	}
	csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
		log.Println("\n- An error ocurred while reading the file. \n", err)
		return nil
    }
	return csvLines
}

func ParseCSVDataToItemsList(csvLines [][]string) (listOfItems []Item ) { 
	// Convert csv lines to a generic item structure and append them to the array of items
    for _, line := range csvLines {
		newItem := Item{
            Id: line[0],
            Title: line[1],
        }
        listOfItems = append(listOfItems, newItem)
        log.Println(newItem.Id + " " + newItem.Title + " ")
    }
	return 
}

var csvFileName string = "test.csv"

func TestWriteDataToCSVFile(t *testing.T) {
	type args struct {
		fileName string
		items    []Item
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Write struct Item to csv",
			args: args{
				fileName: csvFileName, 
				items: Items,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// write the data
			WriteDataToCSVFile(tt.args.fileName, tt.args.items)
			// check if data was written properly
			csvLines := GetDataFromCSVFile(tt.args.fileName)
			listOfItems :=  ParseCSVDataToItemsList(csvLines)   
			log.Print(listOfItems)
			if len(listOfItems) != len(Items) { // is the same length
				t.Errorf("Expected %v items but instead got %v!", len(Items), len(listOfItems))
			}
		})
	}
}