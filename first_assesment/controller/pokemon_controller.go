package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wizelineacademy/academy-go-q12021/core"
	"github.com/wizelineacademy/academy-go-q12021/service"

	"github.com/gorilla/mux"
)

// GetAllPokemons get all pokemons
func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pokemonService, err := service.NewPokemonService()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: "Error creating service",
		})
		return
	}
	pokemons, err := pokemonService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemons)
}

// GetPokemonByID get pokemon based on ID
func GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pokemonService, err := service.NewPokemonService()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: "Error creating service",
		})
		return
	}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["pokemonId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	pokemon, err := pokemonService.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemon)
}
