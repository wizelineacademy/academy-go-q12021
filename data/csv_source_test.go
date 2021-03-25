package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/wizelineacademy/academy-go/model"
)

const correctFile = "./testdata/pokemons_test.csv"
const incorrectFile = "./testdata/pokemons.csv"
const emptyFile = "./testdata/empty_test.csv"

func TestCsvDataSourceGetDataSuccess(t *testing.T) {
	datasource := CsvSource(correctFile)
	data, err := datasource.GetData()
	if err != nil {
		t.Errorf("Error initiating CSV datasource: %v", err)
	}

	csvData := data.CsvData
	if len(csvData) == 6 {
		t.Errorf("Data contains empty lines")
	}
}
func TestCsvDataSourceGetDataError(t *testing.T) {
	csvSource := CsvSource(incorrectFile)
	_, err := csvSource.GetData()
	if err == nil {
		t.Errorf("Initiating CSV datasource should fail")
	}

	txtSource := CsvSource("./testdata/not_csv_file.csv")
	_, err2 := txtSource.GetData()
	if err2 == nil {
		fmt.Println(err2)
		t.Errorf("Initiating CSV datasource should fail")
	}
}

func TestCsvDataSourceSetDataError(t *testing.T) {
	csvSource := CsvSource(incorrectFile)
	testData := [][]string{
		{"5", "test"},
	}
	err := csvSource.SetData(model.NewCsvData(testData))
	if err == nil {
		t.Errorf("Initiating CSV datasource should fail")
	} else if !strings.Contains(err.Error(), "Error accessing the storage:") {
		t.Errorf("Expected a storage error, got '%v'", err.Error())
	}
}

func TestCsvDataSourceSetDataSuccess(t *testing.T) {
	createTestFile()

	csvSource := CsvSource(emptyFile)
	testData := [][]string{
		{"5", "test"},
	}
	err := csvSource.SetData(model.NewCsvData(testData))
	if err != nil {
		t.Errorf("CSV datasource should not fail writing data, got '%v'", err)
	}

	csvData, getDataError := csvSource.GetData()
	if getDataError != nil {
		t.Errorf("CSV datasource should not fail reading data, got '%v'", err)
	}

	if len(csvData.CsvData) != 2 {
		t.Errorf("CSV datasource should add one line, got '%v'", csvData.CsvData)
	}

	deleteTestFile()
}

func createTestFile() {
	header := []byte("ID,Name\n")
	err := ioutil.WriteFile(emptyFile, header, 0666)
	if err != nil {
		panic(err)
	}
}

func deleteTestFile() {
	err := os.Remove(emptyFile)
	if err != nil {
		panic(err)
	}
}
