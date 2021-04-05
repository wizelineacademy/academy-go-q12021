package test

import (
	"encoding/csv"
	"fmt"
	"log"
	"main/controller"
	"main/model"
	"os"
	"testing"
)

var Items []model.TechStackItem = []model.TechStackItem{
	{Id: "1", Title: "Netflix", Years: "5"},
	{Id: "2", Title: "Dispney+", Years: "1"},
	{Id: "3", Title: "HBO Max", Years: "4"},
	{Id: "4", Title: "Paramount+", Years: "2"},
	{Id: "5", Title: "Universal+", Years: "4"},
}

func GetDataFromCSVFile(filePath string) [][]string {
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

func ParseCSVDataToItemsList(csvLines [][]string) (listOfItems []model.TechStackItem) {
	// Convert csv lines to a generic item structure and append them to the array of items
	for _, line := range csvLines {
		newItem := model.TechStackItem{
			Id:    line[0],
			Title: line[1],
			Years: line[2],
		}
		listOfItems = append(listOfItems, newItem)
	}
	return
}

var csvFileName string = "test.csv"

func TestWriteDataToCSVFile(t *testing.T) {
	type args struct {
		fileName string
		items    []model.TechStackItem
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Write struct model.TechStackItem to csv",
			args: args{
				fileName: csvFileName,
				items:    Items,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// write the data
			controller.WriteDataToCSVFile(tt.args.fileName, tt.args.items)
			// check if data was written properly
			csvLines := GetDataFromCSVFile(tt.args.fileName)
			listOfItems := ParseCSVDataToItemsList(csvLines)
			log.Print(listOfItems)
			if len(listOfItems) != len(Items) { // is the same length
				t.Errorf("Expected %v items but instead got %v!", len(Items), len(listOfItems))
			}
			if listOfItems[0].Id != Items[0].Id || listOfItems[0].Title != Items[0].Title { // has the same content
				t.Errorf("Expected hardcoded item %v to match stored item on the csv %v!", Items[0], listOfItems[0])
			}

		})
	}
}
