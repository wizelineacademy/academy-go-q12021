package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

type CatFact struct {
	Id string `json:"id"`
	Fact string `json:"fact"`
}

type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

var catFacts []CatFact

func getCatFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catFacts)
}

func getCatFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, catFact := range catFacts {
		if catFact.Id == params["id"] {
			json.NewEncoder(w).Encode(catFact)
			return
		}
	}
	json.NewEncoder(w).Encode(Error{Code: 1, Message: "We could not find a fact with the specified id"})
}

func main() {
	// Initialize with data
	csvFile, _ := os.Open("catfacts.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		catFacts = append(catFacts, CatFact{
			Id: line[0],
			Fact:  line[1],
		})
	}

	// Creates router
	router := mux.NewRouter()
	router.HandleFunc("/cat-facts", getCatFacts).Methods("GET")
	router.HandleFunc("/cat-facts/{id}", getCatFact).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", router))
}