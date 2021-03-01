package main

import (
	"first/controller"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/pokemons", controller.GetAllPokemons).Methods("GET")
	r.HandleFunc("/api/v1/pokemons/{pokemonId}", controller.GetPokemonById).Methods("GET")

	server := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
