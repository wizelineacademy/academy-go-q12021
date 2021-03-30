package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grethelBello/academy-go-q12021/model"
	"github.com/grethelBello/academy-go-q12021/model/errs"
)

// CsvSource is a module to get information from a CSV file. To init, indicate the path to the file
type CsvSource string

// GetData is an implementation of Source interface which returns a Data struct with the data from the CSV file
func (source CsvSource) GetData(csvConfig ...*model.SourceConfig) (*model.Data, error) {
	path := string(source)
	if len(csvConfig) > 0 {
		path = *&csvConfig[0].CsvConfig
	}

	file, err := os.Open(path)
	if err != nil {
		storageError := errs.StorageError{TechnicalError: err}
		return &model.Data{}, storageError
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		storageError := errs.StorageError{TechnicalError: err}
		return &model.Data{}, storageError
	}

	data := model.NewCsvData(lines)
	return data, nil
}

func (source CsvSource) SetData(generalData *model.Data) error {
	// Open the file to append at the end
	file, err := os.OpenFile(string(source), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return errs.StorageError{TechnicalError: err}
	}
	defer file.Close()

	csvData := generalData.CsvData
	for _, row := range csvData {
		line := strings.Join(row, ",")
		if _, err := file.WriteString(fmt.Sprintf("%v\n", line)); err != nil {
			log.Printf("Error writing '%v' in '%v' file: %v", line, string(source), err)
		}
	}

	return nil
}
