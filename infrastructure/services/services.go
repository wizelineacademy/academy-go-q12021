package services

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/domain/model"
)

func (c *Client) GetStudentsService() ([]model.Student, error) {
	fmt.Println("Reading csv...")
	// open csv
	csvFile, err := os.Open(config.C.CsvPath.Path)
	checkError("Can't open csv", err)

	var students []model.Student
	csvReader := csv.NewReader(csvFile)
	csvReader.Comment = '#'
	for {
		// read csv
		dataRow, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		checkError("Can't read csv file", err)

		// fill struct with data
		student, err := model.Student{}.ToStruct(dataRow)
		checkError("Can't convert data to Student{}", err)
		// add struct student to array  student
		students = append(students, student)
	}
	return students, err
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

/**
ReadURL and return students Array from URL in structure
*/
func (c *Client) GetUrlService() ([]model.Student, error) {
	//consule url
	//unmarshall
	//writetoCSV
	var url = "https://login-app-crud.firebaseio.com/.json"

	var students []model.Student

	resp, err := c.client.R().SetHeader("Accept", "application/json").Get(url)
	checkError("fallo get", err)

	errjson := json.Unmarshal(resp.Body(), &students)
	if errjson != nil {
		log.Fatal("Error in unmarshall", errjson)
	}
	return students, err
}

// write info to CSV

func (c *Client) SaveToCsv(students []model.Student) (bool, error) {
	file, err := os.Create(config.C.CsvPath.Path)
	checkError("Cannot open file", err)
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"#id", "name", "lastname", "gender", "city", "state", "zip", "email", "age"}
	err = w.Write(header)
	for _, v := range students {
		err = w.Write(v.ToSlice())
		if err := w.Error(); err != nil {
			log.Panic("Error writing csv:", err)
		}
	}
	return false, err
}
