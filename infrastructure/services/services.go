package services

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/domain/model"
)

func (c *Client) GetStudentsService() ([]model.Student, error) {
	var students []model.Student

	// open csv
	csvFile, err := os.Open(config.C.CsvPath.Path)
	if err != nil {
		return students, fmt.Errorf("Unable to open csv file")
	}
	defer csvFile.Close()

	// setup csv
	csvReader := csv.NewReader(csvFile)
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = 9
	for {
		// read csv row
		dataRow, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return students, fmt.Errorf("csv reader failure")
		}

		// fill struct with data
		student, err := model.Student{}.ToStruct(dataRow)
		if err != nil {
			return students, fmt.Errorf("cannot convert data to Student {}")
		}
		// add struct student to []Student
		students = append(students, student)
	}
	if students != nil {
		return students, err
	} else {
		return students, fmt.Errorf("csv is empty")
	}
}

/**
ReadURL and return students Array from URL in structure
*/
func (c *Client) GetURLService() ([]model.Student, error) {
	const ApiUrl = "https://login-app-crud.firebaseio.com/.json"
	var students []model.Student

	resp, err := c.client.R().SetHeader(
		"Accept",
		"application/json",
	).Get(ApiUrl)
	if err != nil {
		return students, fmt.Errorf("Could not get the URL information")
	}

	// convert json to []Students
	err = json.Unmarshal(resp.Body(), &students)
	if err != nil {
		return students, fmt.Errorf("Error converting json to [] students")
	}
	return students, err
}

// SaveToCsv  take and []Student and save it in a csv file
func (c *Client) SaveToCsv(students []model.Student) (bool, error) {
	// create csv file
	file, err := os.Create(config.C.CsvPath.Path)
	if err != nil {
		return false, fmt.Errorf("Could not create csv file")
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()
	// add headers to the csv file
	header := []string{"#id", "name", "lastname", "gender", "city", "state", "zip", "email", "age"}
	err = w.Write(header)
	if err != nil {
		return false, err
	}

	// save each struct as a row in csv
	for _, s := range students {
		err = w.Write(s.ToSlice())
		if err := w.Error(); err != nil {
			return false, err
		}
	}
	return true, nil
}
