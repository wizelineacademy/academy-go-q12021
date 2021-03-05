package test

import (
	"path/filepath"
	"testing"

	"../data"
)

func TestCsvDataSourceSuccess(t *testing.T) {
	testFilePath, err := filepath.Abs("./pokemons.csv")
	if err == nil {
		datasource := data.CsvSource(testFilePath)
		_, err := datasource.Init()
		if err != nil {
			t.Errorf("Error initiating CSV datasource: %v", err)
		}
	}

	t.Errorf("Error loading the absolute file: %v", err)
}
func TestCsvDataSourceError(t *testing.T) {
	testFilePath, err := filepath.Abs("../pokemons.csv")
	if err == nil {
		datasource := data.CsvSource(testFilePath)
		_, err := datasource.Init()
		if err != nil {
			t.Errorf("Initiating CSV datasource should fail")
		}
	}
}
