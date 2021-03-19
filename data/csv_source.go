package data

import (
	"encoding/csv"
	"os"

	"github.com/wizelineacademy/academy-go/model/errs"
)

// CsvSource is a module to get information from a CSV file. To init, indicate the path to the file
type CsvSource string

// GetData is an implementation of Source interface which returns a Data struct with the data from the CSV file
func (source CsvSource) GetData() (Data, error) {
	file, err := os.Open(string(source))
	if err != nil {
		storageError := errs.StorageError{TechnicalError: err}
		return Data{}, storageError
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		storageError := errs.StorageError{TechnicalError: err}
		return Data{}, storageError
	}

	data := NewCsvData(lines)
	return data, nil
}
