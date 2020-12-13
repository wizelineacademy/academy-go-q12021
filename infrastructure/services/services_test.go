package services

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/domain/model"
)

func TestReadStudentsService(t *testing.T) {
	config.ReadConfig("config")
	c := NewClient()
	filePath := config.C.CsvPath.Test
	students, err := c.ReadStudentsService(filePath)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(students)
	t.Log(students)
	if len(students) < 1 {
		t.Error("want len >1 ")
	}
}


// Test fail read students
func TestFailReadStudentsService(t *testing.T) {
	c := NewClient()
	_, err := c.ReadStudentsService("wrongpath/dataFile.csv")
	if !errors.Is(err, os.ErrNotExist) {
		t.Error("File dont found.", err)
	}
}

// test store url service
func TestStoreURLService(t *testing.T) {
	err := config.ReadConfig("config")
	if err != nil {
		fmt.Println("read config fail")
	}
	c := NewClient()
	ApiUrl := config.C.Api.Url

	students, err := c.StoreURLService(ApiUrl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(students)
}

// test save succesfully csv file
func TestSaveToCsv(t *testing.T) {
	err := config.ReadConfig("config")
	if err != nil {
		fmt.Println("read config fail")
	}
	c := NewClient()
	s := model.Student{ID: 1, Name: "Ruben"}
	students := []model.Student{s}
	filePath := config.C.CsvPath.Test

	ok, err := c.SaveToCsv(students, filePath)
	if err != nil || ok == false {
		t.Error(err)
	}
}

// test cant save csv file
func TestNotSaveToCsv(t *testing.T) {
	c := NewClient()
	s := model.Student{}
	students := []model.Student{s}
	filePath := ""
	want := "could not create csv file"
	ok, err := c.SaveToCsv(students, filePath)
	if strings.Contains(err.Error(), want) {
		return
	} else {

		t.Errorf("unexpected error: %v   %v", err, ok)
	}
}
