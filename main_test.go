package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"./routes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var r *mux.Router

func init() {
	r = Router()
}

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/bitcoin", routes.GetBicoins).Methods("GET")
	router.HandleFunc("/bitcoin/{id}", routes.GetBicoin).Methods("GET")
	return router
}

// func TestIndexGetEndpoint(t *testing.T) {
// 	request, _ := http.NewRequest("GET", "/", nil)
// 	response := httptest.NewRecorder()
// 	r.ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code, "Ok")
// 	assert.Equal(t, "Service Status: Ok", response.Body.String(), "Ok")
// }

// func TestBitcoinGetEndpoint(t *testing.T) {
// 	request, _ := http.NewRequest("GET", "/bitcoin", nil)
// 	response := httptest.NewRecorder()
// 	r.ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code, "Ok")
// 	assert.Contains(t, response.Body.String(), "\"success\":true", "Ok")
// }

// func TestBitcoinIDGetEndpoint(t *testing.T) {
// 	request, _ := http.NewRequest("GET", "/bitcoin/1", nil)
// 	response := httptest.NewRecorder()
// 	r.ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code, "Ok")
// 	assert.Contains(t, response.Body.String(), "\"success\":true", "Ok")
// }

// func TestBitcoinIDGetEndpointFail(t *testing.T) {
// 	request, _ := http.NewRequest("GET", "/bitcoin/kad", nil)
// 	response := httptest.NewRecorder()
// 	r.ServeHTTP(response, request)
// 	assert.Equal(t, 400, response.Code, "Ok")
// 	assert.Contains(t, response.Body.String(), "\"success\":false", "Ok")
// }

func TestMigrate(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		route       string
		responseInt int
		response    string
	}{
		{
			name:        "Test Index",
			method:      "GET",
			route:       "/",
			responseInt: 200,
			response:    "Service Status: Ok",
		},
		{
			name:        "Test Bitcoin GET",
			method:      "GET",
			route:       "/bitcoin",
			responseInt: 200,
			response:    "\"success\":true",
		},
		{
			name:        "Test Bitcoin ID",
			method:      "GET",
			route:       "/bitcoin/1",
			responseInt: 200,
			response:    "\"success\":true",
		},
		{
			name:        "Test Bitcoin False",
			method:      "GET",
			route:       "/bitcoin/1g",
			responseInt: 400,
			response:    "\"success\":false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(tt.method, tt.route, nil)
			response := httptest.NewRecorder()
			r.ServeHTTP(response, request)
			assert.Equal(t, tt.responseInt, response.Code, "Ok")
			assert.Contains(t, response.Body.String(), tt.response, "Ok")
		})
	}
}
