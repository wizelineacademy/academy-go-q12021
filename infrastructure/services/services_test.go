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
	config.ReadConfig()
	c := NewClient()
	filePath:=config.C.CsvPath.Test
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

func TestFailReadStudentsService(t *testing.T) {
	c := NewClient()
	_, err := c.ReadStudentsService("wrongpath/dataFile.csv")
	if !errors.Is(err, os.ErrNotExist) {
		t.Error("File dont found.", err)
	}
}

func TestEmptyReadStudentsService(t *testing.T) {
	config.ReadConfig()
	c := NewClient()
	pathFile:= config.C.CsvPath.Empty
	_, err := c.ReadStudentsService(pathFile)
	fmt.Printf("%T %v", err.Error(), err.Error())
	t.Log(err.Error())
	if err != nil {
		t.Error(err)
	}

}

func TestStoreURLService(t *testing.T) {
	config.ReadConfig()
	c := NewClient()
	ApiUrl := config.C.Api.Url

	students, err := c.StoreURLService(ApiUrl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(students)
}

func TestSaveToCsv(t *testing.T) {
	config.ReadConfig()
	c := NewClient()
	s := model.Student{ID: 1, Name: "Ruben"}
	students := []model.Student{s}
	filePath := config.C.CsvPath.Test

	ok, err := c.SaveToCsv(students, filePath)
	if err != nil || ok == false {
		t.Error(err)
	}
}
func TestNotSaveToCsv(t *testing.T) {
	c := NewClient()
	s := model.Student{}
	students := []model.Student{s}
	filePath := ""
	want:= "could not create csv file"
	ok, err := c.SaveToCsv(students, filePath)
	 if strings.Contains(err.Error(), want){
	 	return
	 }else{

		 t.Errorf("unexpected error: %v   %v", err, ok)
	 }

}