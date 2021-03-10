package usecases

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wl-project/academy-go-q12021/entities"
	"github.com/wl-project/academy-go-q12021/repository"
	"net/http"
)

var catFacts = repository.LoadData()

func GetCatFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catFacts)
}

func GetCatFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, catFact := range catFacts {
		if catFact.Id == params["id"] {
			json.NewEncoder(w).Encode(catFact)
			return
		}
	}
	json.NewEncoder(w).Encode(entities.HTTPError{Code: 1, Message: "We could not find a fact with the specified id"})
}
