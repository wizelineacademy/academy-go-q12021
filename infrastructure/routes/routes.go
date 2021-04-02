package routes

import (
	"github.com/gorilla/mux"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/handler"
)

func SetupRoutes(r *mux.Router, h *handler.NewsHandlers) {

	r.HandleFunc("/news", h.ListAll).Methods("GET")
	r.HandleFunc("/news/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/news/current", h.FetchAll).Methods("POST")
	r.HandleFunc("/concurrent/news", h.ListAllConcurrenlty).Methods("GET")
}
