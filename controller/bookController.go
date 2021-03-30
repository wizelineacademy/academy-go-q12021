package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joseantoniovz/academy-go-q12021/service"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	var books, err = service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	var book, err = service.GetById(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else if book.Title != "" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}

}

func ConsumeAPI(w http.ResponseWriter, r *http.Request) {
	var book, err = service.ConsumeAPI()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}
