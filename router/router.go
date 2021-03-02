package router

import (
	"log"
	"net/http"

	"academy/controller/jokes"
	"academy/controller/update"

	"github.com/gorilla/mux"
)

// InitServer will execute in the defined port
func InitServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/jokes", jokes.GetJokes).Methods("GET")
	router.HandleFunc("/api/v1/jokes/{id}", jokes.GetOneJoke).Methods("GET")
	router.HandleFunc("/api/v1/new-jokes", update.GetData).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
