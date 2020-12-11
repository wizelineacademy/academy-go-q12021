/**
Student Services
*/
package services

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"golang-bootcamp-2020/domain/model"
)

// ReadStudentsService read students from csv file return []Students
func (c *Client) ReadStudentsService(filePath string) ([]model.Student, error) {
	var students []model.Student

	// open csv
	csvFile, err := os.OpenFile(filePath,os.O_RDWR, 0664)
	if err != nil {
		return students, err
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
			return students, err
		}

		// fill struct with data
		student, err := model.Student{}.ToStruct(dataRow)
		if err != nil {
			return students, err
		}
		// add struct student to []Student
		students = append(students, student)
	}

	return students,nil
}

// StoreURLService and return students Array from URL in structure
func (c *Client) StoreURLService(apiURL string) ([]model.Student, error) {
	var students []model.Student

	resp, err := c.client.R().SetHeader(
		"Accept",
		"application/json",
	).Get(apiURL)
	if err != nil {
		return students, fmt.Errorf("could not get the URL information")
	}

	// convert json to []Students
	err = json.Unmarshal(resp.Body(), &students)
	if err := json.Unmarshal(resp.Body(), &students); err != nil {
		return students, fmt.Errorf("error converting json to [] students")
	}
	return students, err
}

// SaveToCsv take and []Student and save it in a csv file
func (c *Client) SaveToCsv(students []model.Student, filePath string) (bool, error) {
	// create csv file
	//file, err := os.Create(config.C.CsvPath.Path)
	file, err := os.Create(filePath)
	if err != nil {
		return false, fmt.Errorf("could not create csv file")
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()
	// add headers to the csv file
	header := []string{"#id", "name", "lastname", "gender", "city", "state", "zip", "email", "age"}
	err = w.Write(header)
	if err != nil {
		return false, errors.New("fail create csv headers")
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
