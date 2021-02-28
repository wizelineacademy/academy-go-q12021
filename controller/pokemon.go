package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokeapi/model"
	"pokeapi/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

const uri string = "https://pokeapi.co/api/v2/pokemon?limit=5&offset=300"

type PokemonController struct {
	useCase usecase.IUsecase
}

type IPokemonController interface {
	Index(w http.ResponseWriter, r *http.Request)
	GetPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemon(w http.ResponseWriter, r *http.Request)
	GetPokemonsFromExternalAPI(w http.ResponseWriter, r *http.Request)
}

func New(pc usecase.IUsecase) *PokemonController {
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

func (pc *PokemonController) GetPokemonsFromExternalAPI(
	w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	w.Header().Set("Content-Type", "application/json")

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Fprintf(w, "Something happened: %v", err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	var response model.PokemonExternal
	json.Unmarshal(bodyBytes, &response)

	newPokemons := response.Results
	e := pc.useCase.GetPokemonsFromExternalAPI(&newPokemons)

	if e != nil {
		w.WriteHeader(e.Code)
		fmt.Fprintf(w, "There was some errors, please try again.")
	} else {
		fmt.Fprintf(w, "API Response as struct %+v\n", response)
	}

}
