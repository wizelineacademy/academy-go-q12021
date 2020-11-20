package main

import (
	"log"
	"os"

	"github.com/etyberick/golang-bootcamp-2020/csv"
	"github.com/etyberick/golang-bootcamp-2020/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//Check if csv file exists
	filename := "quotes.csv"
	_, err := os.Stat(filename)

	//Create if it doesn't exists
	if os.IsNotExist(err) {
		_, err = os.Create(filename)
		if err != nil {
			log.Fatalf("error creating file - %v", err)
		}
	}

	//Serve
	quoteStorage := csv.NewCsvQuoteRepository(filename)
	http.NewQuoteHandler(r, quoteStorage)
}
