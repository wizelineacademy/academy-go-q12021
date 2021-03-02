package pokemon

import (
	"net/http"
	"bootcamp/usecase/pokemon"
	"bootcamp/service/network"
	"github.com/gorilla/mux"
)

/*
GetPokemon returns a JSON Pokemon array or a Pokemon information
If URL not contains /{id} returns a Pokemon array
If URL contains /{id} return the Pokemon for the given index
*/
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

/*
AddPokemon returns a JSON Pokemon struct with the new added Pokemon information
*/
func AddPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon.AddPokemon(r.Body)
	network.Response(w, pokemon, err)
}

/*
UpdatePokemon returns a JSON Pokemon struct with the updated Pokemon information
*/
func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon.UpdatePokemon(mux.Vars(r), r.Body)
	network.Response(w, pokemon, err)
}

/*
DeletePokemon returns a JSON Pokemon struct with the deleted Pokemon information
*/
func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon.DeletePokemon(mux.Vars(r))
	network.Response(w, pokemon, err)
}