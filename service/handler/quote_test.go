package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/etyberick/golang-bootcamp-2020/entity"
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

func TestGet(t *testing.T) {
	// We need to set some variables before starting the test tables
	req, err := http.NewRequest("GET", "/v0/quote", nil)
	if err != nil {
		t.Errorf("failed mocking request, %v", err.Error())
	}
	// Initialize test tables
	type args struct {
		qi interactor.QuoteInteractor
	}
	tests := []struct {
		args             args
		name             string
		pretest          func(t *testing.T)
		request          *http.Request
		responseRecorder *httptest.ResponseRecorder
	}{
		{
			args: args{
				qi: interactor.NewQuoteInteractor(entity.Config{
					CSVFilepath: "empty.csv",
					Port:        "8080",
				}),
			},
			name:             "empty database get",
			pretest:          nil,
			request:          req,
			responseRecorder: httptest.NewRecorder(),
		},
		{
			args: args{
				qi: interactor.NewQuoteInteractor(entity.Config{
					CSVFilepath: "test.csv",
					Port:        "8080",
				}),
			},
			name:             "populated database GET",
			pretest:          TestUpdate,
			request:          req,
			responseRecorder: httptest.NewRecorder(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Perform pretest conditions
			if tt.pretest != nil {
				t.Run("Pre-test condition", tt.pretest)
			}
			NewQuote(tt.args.qi).GetAll(tt.responseRecorder, tt.request)
			b, err := ioutil.ReadAll(tt.responseRecorder.Body)
			if err != nil {
				t.Errorf(err.Error())
			}
			if len(b) == 0 {
				t.Errorf("GET = %q, want length different from 0", string(b))
			}
		})
	}
}
