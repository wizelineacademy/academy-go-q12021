package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestNewQuote(t *testing.T) {
	type args struct {
		qi interactor.QuoteInteractor
	}
	tests := []struct {
		name string
		args args
		want Quote
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuote(tt.args.qi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quote_Update(t *testing.T) {
	type fields struct {
		quoteInteractor interactor.QuoteInteractor
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quote{
				quoteInteractor: tt.fields.quoteInteractor,
			}
			q.Update(tt.args.w, tt.args.r)
		})
	}
}

func Test_quote_GetAll(t *testing.T) {
	type fields struct {
		quoteInteractor interactor.QuoteInteractor
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quote{
				quoteInteractor: tt.fields.quoteInteractor,
			}
			q.GetAll(tt.args.w, tt.args.r)
		})
	}
}
