package csv

import (
	"bootcamp/usecase/csv"
	"bootcamp/service/network"
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

	network.ResponseList(w, pokemonList, err)
}