package router

import (
	"first/controller"

	"github.com/gorilla/mux"
)

func NewRouting() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/pokemons", controller.GetAllPokemons).Methods("GET")
	r.HandleFunc("/api/v1/pokemons/{pokemonId}", controller.GetPokemonByID).Methods("GET")
	return r
}
