package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokeapi/usecases"
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

func NewPokemonController(pc usecases.IUsecase) *PokemonController {
	return &PokemonController{pc}
}

func (pc *PokemonController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Poke-API")
}

func (pc *PokemonController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons := pc.useCase.GetPokemons()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func (pc *PokemonController) GetPokemon(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pokemonId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	pokemon := pc.useCase.GetPokemon(pokemonId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func (pc *PokemonController) GetPokemonsFromExternalAPI(w http.ResponseWriter, r *http.Request) {
	pc.useCase.GetPokemonsFromExternalAPI()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Pokemons saved correctly")
}
