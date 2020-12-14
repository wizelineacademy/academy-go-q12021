// uni test for services
package services

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/ruvaz/golang-bootcamp-2020/domain/model"
)

// TestReadStudentsService: Successful student CSV reading test
func TestReadStudentsService(t *testing.T) {
	c := NewClient()
	filePath := "../../tmp/dataFile.csv"
	students, err := c.ReadStudentsService(filePath)
	if err != nil {
		t.Error(err)
	}
	if len(students) < 1 {
		t.Error("want len >1 ")
	}
}

// TestFailReadStudentsService: Test failed when trying to read csv to get students
func TestFailReadStudentsService(t *testing.T) {
	c := NewClient()
	_, err := c.ReadStudentsService("wrongpath/dataFile.csv")
	if !errors.Is(err, os.ErrNotExist) {
		t.Error("File dont found.", err)
	}
}

// TestStoreURLService: Service test to get students from an api
func TestStoreURLService(t *testing.T) {
	c := NewClient()
	ApiUrl := "https://login-app-crud.firebaseio.com/.json"
	students, err := c.StoreURLService(ApiUrl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(students)
}

// TestSaveToCsv: Test on successfully saving csv file
func TestSaveToCsv(t *testing.T) {
	c := NewClient()
	s := model.Student{ID: 1, Name: "Ruben"}
	students := []model.Student{s}
	filePath := "../../tmp/dataFile.csv"
	ok, err := c.SaveToCsv(students, filePath)
	if err != nil || ok == false {
		t.Error(err)
	}
}

// TestNotSaveToFolder: test unsuccessfully saving csv file
func TestNotSaveToFolder(t *testing.T) {
	c := NewClient()
	s := model.Student{}
	students := []model.Student{s}
	filePath := ""
	want := "could not create tmp folder"
	ok, err := c.SaveToCsv(students, filePath)
	if !strings.Contains(err.Error(), want) {
		t.Errorf("unexpected error: %v   %v", err, ok)
	}
}
