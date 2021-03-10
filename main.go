package main

import (
	"github.com/gorilla/mux"
	"github.com/wl-project/academy-go-q12021/usecases"
	"log"
	"net/http"
)

func main() {
	// Creates router
	router := mux.NewRouter()
	router.HandleFunc("/cat-facts", usecases.GetCatFacts).Methods("GET")
	router.HandleFunc("/cat-facts/{id}", usecases.GetCatFact).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", router))
}