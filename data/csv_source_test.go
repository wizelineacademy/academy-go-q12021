package data

import (
	"fmt"
	"testing"
)

func TestCsvDataSourceSuccess(t *testing.T) {
	datasource := CsvSource("./testdata/pokemons_test.csv")
	data, err := datasource.GetData()
	if err != nil {
		t.Errorf("Error initiating CSV datasource: %v", err)
	}

	csvData := data.CsvData
	if len(csvData) == 6 {
		t.Errorf("Data contains empty lines")
	}
}
func TestCsvDataSourceError(t *testing.T) {
	csvSource := CsvSource("./testdata/pokemons.csv")
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
