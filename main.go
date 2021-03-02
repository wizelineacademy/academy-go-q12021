package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Book contains information about each book
type Book struct {
	ID       int    `json:"id"`
	Isbn     string `json:"isbn"`
	Authors  string `json:"authors"`
	Year     int    `json:"year"`
	ImageURL string `json:"imageURL"`
}

func loadData(books []Book, path string) []Book {

	f, err := os.Open("data/books.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// Parse id to int
		id, errConv := strconv.Atoi(record[0])

		if errConv != nil {
			log.Fatal(errConv)
		}

		// Parse year to int
		year, errConv := strconv.ParseFloat(record[4], 64)

		if errConv != nil {
			log.Fatal(errConv)
		}

		books = append(books, Book{ID: id, Isbn: record[1], Authors: record[3], Year: int(year), ImageURL: record[5]})

	}

	return books
}

// findBookByID returns the index that corresponds to an ID (if exists)
func findBookByID(books []Book, ID int) int {

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
var books []Book

func main() {

	//Init Router
	r := mux.NewRouter()

	books = loadData(books, "data/books.csv")

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
