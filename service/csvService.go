package service

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/oscarSantoyo/academy-go-q12021/model"

	"github.com/fatih/structs"
	"github.com/golobby/container"
	"github.com/labstack/gommon/log"
)

// CsvService is used as an interface for this service.
type CsvService interface {
	FilterByID(string) ([]model.Doc, error)
	DownloadCsvData() error
}

// CsvServiceImpl is used as the implementation of this service.
type CsvServiceImpl struct{}

// DownloadCsvData downloads information from external API.
func (c CsvServiceImpl) DownloadCsvData() error {
	getAndSaveData()
	return nil
}

// FilterByID reads records from a CSV and returns the filtered ones
func (c CsvServiceImpl) FilterByID(id string) ([]model.Doc, error) {
	file, err := os.Open(getConfigService().GetConfig().CSV.FileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	reader := csv.NewReader(file)
	records, errReader := reader.ReadAll()
	if errReader != nil {
		return nil, errReader
	}

	var recordsStruct []model.Doc

	for _, record := range records {

		if record[0] != id {
			continue
		}

		recordsStruct = append(recordsStruct, toDoc(record))
	}

	if len(records) == 0 {
		log.Info("The file is empty")
		return nil, errors.New("The file is empty")
	}
	defer file.Close()

	return recordsStruct, nil
}

func toDoc(record []string) model.Doc {
	return model.Doc{
		Key:       record[0],
		Title:     record[1],
		Type:      record[2],
		Published: record[3],
	}
}

func getAndSaveData() {
	response, err := http.Get(getConfigService().GetConfig().External.ApiUrl)

	if err != nil {
		log.Info(err.Error())
		panic(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.SearchResponse

	json.Unmarshal(responseData, &responseObject)

	file, err := os.Create(getConfigService().GetConfig().CSV.FileName)

	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	saveCsv(responseObject, *file)

}

func saveCsv(responseObject model.SearchResponse, file os.File) {
	writer := csv.NewWriter(&file)
	for _, doc := range responseObject.Doc {
		str := interfaceToString(structs.Values(doc))
		err := writer.Write(str)
		if err != nil {
			log.Fatal("cannot write CSV", err)
		}
	}
	defer writer.Flush()
}

func interfaceToString(record []interface{}) []string {
	a := make([]string, len(record))

	for i, row := range record {
		a[i] = row.(string)
	}
	return a
}

func getConfigService() ConfigService {
	var config ConfigService
	container.Make(&config)
	return config
}
