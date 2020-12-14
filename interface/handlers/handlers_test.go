package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/wizelineacademy/golang-bootcamp-2020/domain/repositories"
)

// Dummy logger
var l = log.New(ioutil.Discard, "", 0)

func TestInsertCharacterByID(t *testing.T) {
	// CSV
	csvFile, err := os.OpenFile("./rickandmortytestdata.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		l.Fatalf("message: fatal error config file not found: %s", err)
	}
	defer csvFile.Close()
	// HTTP client for requests
	client := &http.Client{}
	// HelloWorld Handler
	rm := &RickMortyHandler{l, l, client, repositories.NewCharRepo(csvFile)}

	// Create a new ResponseRecorder
	rr := httptest.NewRecorder()
	// Create a dummy http.Request
	req, err := http.NewRequest(http.MethodPost, "/rickmorty/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc("/rickmorty/{id}", rm.InsertCharacterByID)
	router.ServeHTTP(rr, req)

	// Get the handler response
	rs := rr.Result()

	// Check if we get 200
	if rs.StatusCode != http.StatusOK {
		t.Errorf("wrong handler status code: got %v want %v", rs.StatusCode, http.StatusOK)
	}

	expected := "Rick Sanchez"
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	if !strings.Contains(string(body), expected) {
		t.Errorf("body response doesn't contain expected string: got %v want %v", rr.Body.String(), expected)
	}
}
