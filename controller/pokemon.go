package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	usecases "pokeapi/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type PokemonController struct {
	useCase usecases.IUsecase
}

type IPokemonController interface {
	Index(w http.ResponseWriter, r *http.Request)
	GetPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemon(w http.ResponseWriter, r *http.Request)
	GetPokemonsFromExternalAPI(w http.ResponseWriter, r *http.Request)
}

func New(pc usecases.IUsecase) *PokemonController {
	return &PokemonController{pc}
}

func (pc *PokemonController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Poke-API")
}

func (pc *PokemonController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := pc.useCase.GetPokemons()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, "Uhh oh... %v", err.Message)
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
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func (pc *PokemonController) GetPokemonsFromExternalAPI(w http.ResponseWriter, r *http.Request) {
	err := pc.useCase.GetPokemonsFromExternalAPI()

	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, "Uhh oh... %v", err.Message)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Pokemons saved correctly, go and checkout your DB! (csv)")
}
