package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Controller interface with the structure of the implemented functions
type Controller interface {
	GetPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemon(w http.ResponseWriter, r *http.Request)
	GetPokemonsFromAPI(w http.ResponseWriter, r *http.Request)
	GetPokemonFromAPI(w http.ResponseWriter, r *http.Request)
	GetPokemonsConcurrency(w http.ResponseWriter, r *http.Request)
}

//New fuction to init the routers
func New(c Controller, r *mux.Router) {
	first := r.PathPrefix("/api/v1").Subrouter()
	first.HandleFunc("/pokemons", c.GetPokemons).Methods(http.MethodGet).Name("get_Pokemons")
	first.HandleFunc("/pokemon/{id}", c.GetPokemon).Methods(http.MethodGet).Name("get_Pokemon")

	api := r.PathPrefix("/api/v2").Subrouter()
	api.HandleFunc("/pokemons", c.GetPokemonsFromAPI).Methods(http.MethodGet).Name("get_Pokemons")
	api.HandleFunc("/pokemon/{pokemonNumber}", c.GetPokemonFromAPI).Methods(http.MethodGet).Name("get_pokemon")

	final := r.PathPrefix("/api/final").Subrouter()
	final.HandleFunc("/", c.GetPokemonsConcurrency).Methods(http.MethodGet).Name("get_dosomenthing")
}
