package controller

import (
	"log"
	"net/http"
	"time"

	"first/controller"

	"github.com/gorilla/mux"
)

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

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
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
	pokemon, err := pokemonService.GetById(id)
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
