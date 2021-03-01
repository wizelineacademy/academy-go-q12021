package csv

import (
	"bootcamp/usecase/csv"
	"bootcamp/service/network"
	"bootcamp/utils"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"errors"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonList, err := csv.GetPokemon()

	if err == nil {
		params := mux.Vars(r)

		if params["id"] != "" {
			id, _ := strconv.Atoi(params["id"])
			pokemonId := id - 1
	
			if pokemonId <= len(pokemonList) - 1 {
				network.Response(w, pokemonList[pokemonId], err)
				return
			}

			err = errors.New("Invalid index")
		}
	}

	queryParams := r.URL.Query()

	if len(queryParams) > 0 {
		pokemon:= utils.GetPokemonByKey(queryParams, pokemonList)
		network.Response(w, pokemon, err)
		return
	}

	network.ResponseList(w, pokemonList, err)
}