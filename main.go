package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/lbswl/academy-go-q12021/domain/model"
	"github.com/lbswl/academy-go-q12021/infrastructure/datastore"

	"github.com/gorilla/mux"
)

// findBookByID returns the index that corresponds to an ID (if exists)
func findBookByID(books []model.Book, ID int) int {

	for index, item := range books {
		if item.ID == ID {
			return index
		}
	}

	return -1
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	id, errConv := strconv.Atoi(params["id"])

	if errConv != nil {
		log.Fatal(errConv)
	}

	index := findBookByID(books, id)

	if index > -1 {
		json.NewEncoder(w).Encode(books[index])
		return
	}
	json.NewEncoder(w).Encode(books)

}

// Data
var books []model.Book

func main() {

	//Init Router
	r := mux.NewRouter()

	books = datastore.LoadData(books, "data/books.csv")

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
