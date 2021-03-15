package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"pokeapi/usecase"

	"github.com/gorilla/mux"
)

type PokemonController struct {
	useCase usecase.NewPokemonUsecase
}

type NewPokemonController interface {
	Index(w http.ResponseWriter, r *http.Request)
	GetPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemon(w http.ResponseWriter, r *http.Request)
	GetPokemonsFromExternalAPI(w http.ResponseWriter, r *http.Request)
	GetPokemonConcurrently(w http.ResponseWriter, r *http.Request)
}

func New(pc usecase.NewPokemonUsecase) *PokemonController {
	return &PokemonController{pc}
}

func (pc *PokemonController) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{ "message": "Welcome to my Poke-API" }`)
}

func (pc *PokemonController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := pc.useCase.GetPokemons()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, "Uhh oh... %v", err.Message)
		return
	}

	json.NewEncoder(w).Encode(pokemons)
}

func (pc *PokemonController) GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	pokemon, errs := pc.useCase.GetPokemon(pokemonId)

	if err != nil {
		w.WriteHeader(errs.Code)
		fmt.Fprintf(w, "Uhh oh... %v", errs.Message)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func (pc *PokemonController) GetPokemonsFromExternalAPI(
	w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response, err := pc.useCase.GetPokemonsFromExternalAPI()
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, "There was some errors, please try again.")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (pc *PokemonController) GetPokemonConcurrently(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeConcurrency := vars["type"]
	itemsS := r.FormValue("items")
	itemsPerWorkerS := r.FormValue("items_per_worker")

	items, _ := strconv.Atoi(r.FormValue("items"))
	itemsPerWorker, _ := strconv.Atoi(r.FormValue("items_per_worker"))

	pokemons, _ := pc.useCase.GetPokemonsConcurrently(items, itemsPerWorker)

	if typeConcurrency == "even" || typeConcurrency == "odd" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(typeConcurrency + " " + itemsS + " " + itemsPerWorkerS)
		json.NewEncoder(w).Encode(&pokemons)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "You only can use \"even\" or \"odd\"")
	}
}
