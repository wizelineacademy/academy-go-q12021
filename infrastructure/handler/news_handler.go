package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	err := h.newsInteractor.FetchAll()
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "Unable to retrieve all articles."))
		return
	}
	resp := dto.Respose{Message: "Current News fetched and saved on csv file."}
	json.NewEncoder(w).Encode(resp)

}

func (h *NewsHandlers) ListAllConcurrenlty(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=uft-8")

	v := req.URL.Query()
	limitResults := v.Get("limit")
	itemsPerWorker := v.Get("itemsPerWorker")
	category := v.Get("category")
	if limitResults == "" {
		limitResults = "10"
	}

	if itemsPerWorker == "" {
		itemsPerWorker = "3"
	}

	nLimit, err := strconv.Atoi(limitResults)
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "'limit' query param must be a number greater than zero."))
		return
	}

	nItemsPerWorker, err := strconv.Atoi(itemsPerWorker)
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "'itemsPerWorker' query param must be a number greater than zero."))
		return
	}

	if nLimit <= 0 {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "'limit' query param must be greater than zero."))
		return
	}

	if nItemsPerWorker <= 0 {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "'itemsPerWorker' query param must be greater than zero."))
		return
	}

	newsArticles, err := h.newsInteractor.FindAllByCategoryConcurrenlty(category, nLimit, nItemsPerWorker)
	if err != nil {
		json.NewEncoder(w).Encode(dto.NewErrorResponse(err.Error(), "Unable to retrieve all articles."))
		return
	}

	json.NewEncoder(w).Encode(newsArticles)

}
