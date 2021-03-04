package main

import (
	"api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/read", handlers.GetAll).Methods(http.MethodGet)
	api.HandleFunc("/read/{pokemonId}", handlers.GetById).Methods(http.MethodGet)
	api.HandleFunc("", handlers.Post).Methods(http.MethodPost)
	api.HandleFunc("", handlers.Put).Methods(http.MethodPut)
	api.HandleFunc("", handlers.Delete).Methods(http.MethodDelete)

	api.HandleFunc("/user/{userID}/comment/{commentID}", handlers.Params).Methods(http.MethodGet)

	log.Println("Server started listening on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
