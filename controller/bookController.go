package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/joseantoniovz/academy-go-q12021/service"

	log "github.com/sirupsen/logrus"
)

type BookController interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	GetBookById(w http.ResponseWriter, r *http.Request)
	ConsumeAPI(w http.ResponseWriter, r *http.Request)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var books, err = service.GetAll()
	if err != nil {
		log.Error("Something failed ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting books")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	var book, err = service.GetById(mux.Vars(r)["id"])
	if err != nil {
		log.Error("Something failed ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting book with id: " + mux.Vars(r)["id"])
	} else if book.Title != "" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	} else {
		log.Warn("Book with id: " + mux.Vars(r)["id"] + " not found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Book with id: " + mux.Vars(r)["id"] + " not found")
	}

}

func ConsumeAPI(w http.ResponseWriter, r *http.Request) {
	var book, err = service.ConsumeAPI(mux.Vars(r)["id"])
	if err != nil {
		log.Error("Something failed ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

func ConcurrencyBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeNumber := vars["type"]
	items, er := strconv.Atoi(vars["items"])
	workers, er := strconv.Atoi(vars["items_per_workers"])
	if er != nil {
		log.Error("Something failed reading params", er)
	}

	var books, err = service.ConcurrencyBooks(typeNumber, items, workers)
	if err != nil {
		log.Error("Something failed ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting books")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}
}
