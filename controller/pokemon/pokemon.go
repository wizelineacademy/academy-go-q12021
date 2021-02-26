package pokemon

import (
	"bootcamp/usecase/pokemon"
	"bootcamp/service/network"
	"net/http"
	"github.com/gorilla/mux"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if id := params["id"]; id != "" {
		pokemon, err := pokemon.GetPokemonById(params)
		network.Response(w, pokemon, err)
	} else {
		pokemonList, err := pokemon.GetPokemon()
		network.ResponseList(w, pokemonList, err)
	}
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon.AddPokemon(r.Body)
	network.Response(w, pokemon, err)
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon.UpdatePokemon(mux.Vars(r), r.Body)
	network.Response(w, pokemon, err)
}

func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon.DeletePokemon(mux.Vars(r))
	network.Response(w, pokemon, err)
}