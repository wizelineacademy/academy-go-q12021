package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/etyberick/golang-bootcamp-2020/entity"
	"github.com/etyberick/golang-bootcamp-2020/service/repository"
	"github.com/etyberick/golang-bootcamp-2020/usecase/interactor"
)

const quoteRoute = "/v0/quote"

// Normal update test
func TestUpdate(t *testing.T) {
	// Mock a new request
	req, err := http.NewRequest("POST", "/v0/quote", nil)
	if err != nil {
		t.Errorf("failed mocking request, %v", err.Error())
	}
	// Initialize interactor and pass it to the handler
	c := &entity.Config{
		CSVFilepath: "test.csv",
		Port:        "8080",
	}
	qi := interactor.NewQuoteInteractor(*c)
	q := NewQuote(qi)
	// Execute request and check result
	rec := httptest.NewRecorder()
	q.Update(rec, req)
	if rec.Result().StatusCode != http.StatusOK {
		t.FailNow()
	}
}

// Test when fetching from an empty database
func TestGetAllEmptyDatabase(t *testing.T) {
	// Mock a new request
	req, err := http.NewRequest("GET", "/v0/quote", nil)
	if err != nil {
		t.Errorf("failed mocking request, %v", err.Error())
	}
	// Initialize interactor and pass it to the handler
	c := &entity.Config{
		CSVFilepath: "empty.csv",
		Port:        "8080",
	}
	qi := interactor.NewQuoteInteractor(*c)
	q := NewQuote(qi)
	// Execute request and check result
	rec := httptest.NewRecorder()
	q.GetAll(rec, req)
	b, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	if strings.Compare(strings.TrimSpace(string(b)), repository.EmptyDatabase) != 0 {
		t.FailNow()
	}
}

// Test when fetching from a database recently populated
func TestGetAll(t *testing.T) {
	t.Run("TestUpdate", TestUpdate)
	// Mock a new request
	req, err := http.NewRequest("GET", "/v0/quote", nil)
	if err != nil {
		t.Errorf("failed mocking request, %v", err.Error())
	}
	// Initialize interactor and pass it to the handler
	c := &entity.Config{
		CSVFilepath: "test.csv",
		Port:        "8080",
	}
	qi := interactor.NewQuoteInteractor(*c)
	q := NewQuote(qi)
	// Execute request and check result
	rec := httptest.NewRecorder()
	q.GetAll(rec, req)
	if err != nil {
		t.Errorf(err.Error())
	}
	if rec.Result().StatusCode != http.StatusOK {
		t.FailNow()
	}
}
