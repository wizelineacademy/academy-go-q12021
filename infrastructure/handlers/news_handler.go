package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jesus-mata/academy-go-q12021/infrastructure/dto"
	"github.com/jesus-mata/academy-go-q12021/usecase/interactors"
)

type NewsHandlers struct {
	newsInteractor interactors.NewsArticlesInteractor
}

func NewNewsHandlers(newsInteractor interactors.NewsArticlesInteractor) *NewsHandlers {
	return &NewsHandlers{newsInteractor}
}

func (h *NewsHandlers) Setup(r *mux.Router) {

	r.HandleFunc("/news", h.ListAll).Methods("GET")
	r.HandleFunc("/news/{id}", h.GetByID).Methods("GET")
}

func (h *NewsHandlers) ListAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=uft-8")

	newsArticles, err := h.newsInteractor.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "Unable to retrieve all articles."))
		return
	}

	json.NewEncoder(w).Encode(newsArticles)

}

func (h *NewsHandlers) GetByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=uft-8")

	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(fmt.Sprintf("'%v' is not a valid ID. It must be a numeric value.", vars["id"]), "Article ID is not valid."))
		return
	}

	newsArticles, err := h.newsInteractor.GetByID(id)
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "Unable to Retrieve Article"))
		return
	}

	json.NewEncoder(w).Encode(newsArticles)

}
