package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"./routes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/bitcoin", routes.GetBicoins).Methods("GET")
	router.HandleFunc("/bitcoin/{id}", routes.GetBicoin).Methods("GET")
	return router
}

func TestIndexGetEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok")
	assert.Equal(t, "Service Status: Ok", response.Body.String(), "Ok")
}

func TestBitcoinGetEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/bitcoin", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok")
	assert.Contains(t, response.Body.String(), "\"success\":true", "Ok")
}

func TestBitcoinIDGetEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/bitcoin/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok")
	assert.Contains(t, response.Body.String(), "\"success\":true", "Ok")
}

func TestBitcoinIDGetEndpointFail(t *testing.T) {
	request, _ := http.NewRequest("GET", "/bitcoin/kad", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code, "Ok")
	assert.Contains(t, response.Body.String(), "\"success\":false", "Ok")
}
