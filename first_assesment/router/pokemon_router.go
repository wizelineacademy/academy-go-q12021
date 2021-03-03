package router

import (
	"github.com/wizelineacademy/academy-go-q12021/controller"

	"github.com/gorilla/mux"
)

// NewRouting Create microservice routing urls
func NewRouting() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/pokemons", controller.GetAllPokemons).Methods("GET")
	r.HandleFunc("/api/v1/pokemons/{pokemonId}", controller.GetPokemonByID).Methods("GET")
	return r
}
