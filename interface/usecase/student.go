package usecase

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/domain/model"
)

// StudentController interface
type StudentController interface {
	GetStudentsFromCsv() ([]model.Student , error)
}

type Usecase struct {
	service StudentController
}

func New (s StudentController) *Usecase{
	return &Usecase{s}
}

// usecase
// GetStudents get students from csv
func (u *Usecase) GetStudentsFromCsv() ([]model.Student, error) {
	csvFile, err := os.Open(config.C.CsvPath.Path)
	check(err)

	u.service.GetStudentsFromCsv()

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
	return students, err
}

// check log if error exist
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
