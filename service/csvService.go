package service

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fatih/structs"
)

type CsvService interface {
	FilterById(string) ([]Doc, error)
	DownloadCsvData() error
}

type CsvServiceImpl struct{}

type SearchResponse struct {
	Doc      []Doc `json:"docs"`
	NumFound int   `json:"numFound"`
	Start    int   `json:"start"`
}

type Doc struct {
	Key       string `json:"key"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Published string `json:"first_published_year"`
}

var CSVFile = "/home/omar/Desktop/saved.csv"

// func (c )LoadCsvData
func (c CsvServiceImpl) DownloadCsvData() error {
	getAndSaveData()
	return nil
}

func (c CsvServiceImpl) FilterById(id string) ([]Doc, error) {
	file, err := os.Open(CSVFile)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, errReader := reader.ReadAll()
	if errReader != nil {
		return nil, errReader
	}

	var recordsStruct []Doc

	for _, record := range records {
		fmt.Println(record)

		doc := Doc{
			Key:       record[0],
			Title:     record[1],
			Type:      record[2],
			Published: record[3],
		}
		if doc.Key == id {
			recordsStruct = append(recordsStruct, doc)
		}
	}

	if len(records) == 0 {
		fmt.Println("The file is empty")
		return nil, errors.New("The file is empty")
	}
	defer file.Close()

	return recordsStruct, nil
}

// func fileExists(file string) bool {
// 	if _, err := os.Stat(file); os.IsNotExist(err) {
// 		fmt.Println("File "+file+" does not exists")
// 		return false
// 	}
// 	return true
// }

// func readFileByRow(file string) ([]string, error) {

// }
func getAndSaveData() {

	response, err := http.Get("http://openlibrary.org/search.json?q=the+lord+of+the+rings&page=1")

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject SearchResponse

	json.Unmarshal(responseData, &responseObject)

	file, err := os.Create(CSVFile)
	defer file.Close()

	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	SaveCsv(responseObject, *file)

}

func SaveCsv(responseObject SearchResponse, file os.File) {
	writer := csv.NewWriter(&file)
	defer writer.Flush()
	for _, doc := range responseObject.Doc {
		err := writer.Write(interfaceToString(structs.Values(doc)))
		if err != nil {
			log.Fatal("cannot write CSV", err)
		}
	}
}

func interfaceToString(record []interface{}) []string {
	var a []string

	for _, row := range record {
		a = append(a, row.(string))
	}
	return a
}
