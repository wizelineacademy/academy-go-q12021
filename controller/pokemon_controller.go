package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wizelineacademy/academy-go-q12021/business"
	"github.com/wizelineacademy/academy-go-q12021/core"
	"github.com/wizelineacademy/academy-go-q12021/repository"
	"github.com/wizelineacademy/academy-go-q12021/service"

	"github.com/gorilla/mux"
)

// GetAllPokemons get all pokemons
func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pokemonRepository, err := repository.NewPokemonRepository()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: "Error creating repository",
		})
		return
	}

	pokemonService := service.NewExternalPokemonAPI()

	pokemonBusiness, err := business.NewPokemonBusiness(pokemonRepository, pokemonService)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: "Error creating business",
		})
		return
	}
	pokemons, err := pokemonBusiness.GetAll()
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

	pokemonRepository, err := repository.NewPokemonRepository()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: "Error creating repository",
		})
		return
	}

	pokemonService := service.NewExternalPokemonAPI()

	pokemonBusiness, err := business.NewPokemonBusiness(pokemonRepository, pokemonService)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: "Error creating business",
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
	pokemon, err := pokemonBusiness.GetByID(id)
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

// StorePokemonByID get pokemon based on ID
func StorePokemonByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pokemonRepository, err := repository.NewPokemonRepository()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: "Error creating repository",
		})
		return
	}
	pokemonService := service.NewExternalPokemonAPI()

	pokemonBusiness, err := business.NewPokemonBusiness(pokemonRepository, pokemonService)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorResponse{
			Error: "Error creating business",
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
	pokemon, err := pokemonBusiness.StoreByID(id)
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
