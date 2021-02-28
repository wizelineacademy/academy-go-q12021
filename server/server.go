package server

import (
	"log"
	"net/http"

	"academy/handlers/fetch"
	"academy/handlers/jokes"

	"github.com/gorilla/mux"
)

// InitServer will execute in the defined port
func InitServer() {

	jokes.Load()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", jokes.IndexRoute)
	router.HandleFunc("/api/v1/jokes", jokes.GetJokes).Methods("GET")
	router.HandleFunc("/api/v1/jokes/{id}", jokes.GetOneJoke).Methods("GET")
	router.HandleFunc("/api/v1/fetch", fetch.GetData).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
