package controller

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"encoding/csv"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/domain/model"
)

// StudentController interface
type StudentController interface {
	GetStudents() []model.Student
}

// GetStudents get students from csv
func GetStudents() []model.Student {
	csvFile, err := os.Open(config.CsvPath)
	check(err)

	var students []model.Student

	csvReader := csv.NewReader(csvFile)
	csvReader.Comment = '#'
	for {
		dataRow, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		check(err)
		id, err := strconv.Atoi(dataRow[0])
		check(err)
		zip, err := strconv.Atoi(dataRow[6])
		check(err)
		student := model.Student{
			ID:       id,
			Name:     dataRow[1],
			LastName: dataRow[2],
			Gender:   dataRow[3],
			City:     dataRow[4],
			State:    dataRow[5],
			Zip:      zip,
			Email:    dataRow[7],
			Age:      dataRow[8],
		}
		students = append(students, student)
	}
	return students
}

// check log if error exist
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
