package handler

import (
	"net/http"

	"github.com/etyberick/golang-bootcamp-2020/usecase/interactor"
)

type quote struct {
	quoteInteractor interactor.QuoteInteractor
}

// Quote handles the quotes related requests
type Quote interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

// NewQuote handler
func NewQuote(qi interactor.QuoteInteractor) Quote {
	return &quote{quoteInteractor: qi}
}

// Update triggers updating the quotes repository
func (q *quote) Update(w http.ResponseWriter, r *http.Request) {
	// Request the interactor to perform the update
	quote, err := q.quoteInteractor.Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
	w.Write(quote)
	return
}

// GetQuotes will return all the available quotes on each call
func (q *quote) GetAll(w http.ResponseWriter, r *http.Request) {
	// Request the interactor for all the stored quotes
	quotes, err := q.quoteInteractor.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(quotes)
}
