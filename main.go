package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Movie struct
type Movie struct {
	ID       int    `json:"ID"`
	Title    string `json:"Title"`
	Director string `json:"Director"`
}

// Movies list
var Movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	csvFile, err := os.Open("movies.csv")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	for i, line := range csvLines {
		movie := Movie{
			ID:       i,
			Title:    line[1],
			Director: line[2],
		}
		Movies = append(Movies, movie)
	}
	json.NewEncoder(w).Encode(Movies)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/movies", getMovies)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
