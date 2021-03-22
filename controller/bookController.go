package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joseantoniovz/academy-go-q12021/model"
	"github.com/joseantoniovz/academy-go-q12021/util"
)

func loadData() []model.Book {
	var books, err = util.GetAll()

	if err != nil {
		//fmt.Println(err)
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		fmt.Println(err)
	}
	return books

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books = loadData()
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books = loadData()
	id := mux.Vars(r)["id"]
	for _, book := range books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(model.Error{Message: "Error find the book", Code: 1})
}
