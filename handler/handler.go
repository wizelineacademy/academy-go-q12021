package handler

import (
	"github.com/javiertlopez/golang-bootcamp-2020/usecase"

	"github.com/gorilla/mux"
)

// Handler struct holds the usecase
type Handler struct {
	events usecase.Events
}

// New returns a Handler
func New(events usecase.Events) Handler {
	return Handler{
		events,
	}
}

// Router returns a *mux.Router
func (h *Handler) Router() *mux.Router {
	// Create a Gorilla Mux router
	router := mux.NewRouter()

	// addEventHandler
	h.addEventHandler(router)

	return router
}
