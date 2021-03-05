package data

import (
	"encoding/csv"
	"os"

	"../model/errs"
)

type CsvSource string

func (source CsvSource) Init() ([][]string, error) {
	file, err := os.Open(string(source))
	if err != nil {
		storageError := errs.StorageError{TechnicalError: err}
		return [][]string{}, storageError
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		storageError := errs.StorageError{TechnicalError: err}
		return [][]string{}, storageError
	}

	return lines, nil
}
