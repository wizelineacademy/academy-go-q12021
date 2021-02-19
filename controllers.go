package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Poke-API")
}

func getPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonID, err := strconv.Atoi(vars["id"])
	var isPokemon bool

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, pokemon := range pokemons {
		if pokemon.ID == pokemonID {
			isPokemon = true
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pokemon)
		}
	}

	if !isPokemon {
		fmt.Fprintf(w, "There is no pokemon with ID: %d", pokemonID)
	}
}
