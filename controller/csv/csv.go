package csv

import (
	"net/http"
	"bootcamp/service/network"
	"bootcamp/usecase/csv"
)

/*
GetPokemon returns a JSON with the Pokemon information
If URL not contains /{id} nor query params return a Pokemon array
If URL contains /{id} return the Pokemon for the given index
If URL contains a query params look for a Pokemon that matches with that search filter
*/
func GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonList, err := csv.GetPokemon(r)

	if len(pokemonList) == 1 {
		network.Response(w, pokemonList[0], err)
		return
	}

	network.ResponseList(w, pokemonList, err)
}