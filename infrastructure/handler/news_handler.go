package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jesus-mata/academy-go-q12021/application/interactors"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/dto"
)

type NewsHandlers struct {
	newsInteractor interactors.NewsArticlesInteractor
}

func NewNewsHandlers(newsInteractor interactors.NewsArticlesInteractor) *NewsHandlers {
	return &NewsHandlers{newsInteractor}
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

	id := vars["id"]

	newsArticles, err := h.newsInteractor.GetByID(id)
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "Unable to Retrieve Article"))
		return
	}

	json.NewEncoder(w).Encode(newsArticles)

}

func (h *NewsHandlers) FetchAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=uft-8")

	newsArticles, err := h.newsInteractor.FetchAll()
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "Unable to retrieve all articles."))
		return
	}

	json.NewEncoder(w).Encode(newsArticles)

}
