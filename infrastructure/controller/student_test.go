// unitest controller
package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-bootcamp-2020/infrastructure/services"
	"golang-bootcamp-2020/interface/usecase"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Test StoreStudentURLHandler
func TestStoreStudentURLHandler(t *testing.T) {
	s := services.NewClient()
	u := usecase.NewUsecase(s)
	c := NewController(u)
	router := mux.NewRouter()
	router.HandleFunc("/api/storedata", c.StoreStudentURLHandler).Methods("GET")
	req, _ := http.NewRequest("GET", "/api/storedata", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code, "Ok response expected")
}

func TestReadStudentsHandler(t *testing.T) {
	s := services.NewClient()
	u := usecase.NewUsecase(s)
	c := NewController(u)
	router := mux.NewRouter()
	router.HandleFunc("/api/readcsv", c.ReadStudentsHandler).Methods("GET")
	req, _ := http.NewRequest("GET", "/api/readcsv", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code, "Ok response expected")
}

// test for ResponseWithError and ResponseWithJSON
func TestResponseErrorJson(t *testing.T) {
	w := httptest.NewRecorder()
	respondWithError(w, http.StatusBadGateway, "bad gateway")
	assert.Equal(t, 502, w.Code)
	assert.Equal(t, "{\"msj\":\"bad gateway\",\"ok\":\"false\"}", w.Body.String())
}
