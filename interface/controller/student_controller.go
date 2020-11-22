package controller

import (
	"encoding/csv"
	"errors"
	"golang-bootcamp-2020/domain/model"
	"io"
	"log"
	"os"
	"strconv"
)

type StudentController interface {
	GetStudents() []model.Student
}

func GetStudents() []model.Student {
	csvFile, err := os.Open("infrastructure/datastore/dataFile.csv")
	check(err)

	var students []model.Student = nil

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
		students = append(
			students,
			student,
		)
	}
	return students
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
