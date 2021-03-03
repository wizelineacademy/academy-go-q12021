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

var books []model.Book

func main() {

	//Init Router
	r := mux.NewRouter()

	books = datastore.LoadData(books, "data/books.csv")

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", GetBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook returns a book given its ID
func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	id, errConv := strconv.Atoi(params["id"])

	if errConv != nil {
		log.Fatal(errConv)
	}

	index := model.FindBookByID(books, id)

	if index > -1 {
		json.NewEncoder(w).Encode(books[index])
		return
	}
	json.NewEncoder(w).Encode(books)

}
