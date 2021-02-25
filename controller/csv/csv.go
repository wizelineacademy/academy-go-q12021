package csv

import (
	"bootcamp/usecase/csv"
	"bootcamp/service/network"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokeList, err := csv.GetPokemon()

	if err != nil {
		network.UnsuccessfulResponse(w, "Could not load information")
		return
	}

	params := mux.Vars(r)

	if params["id"] != "" {
		id, _ := strconv.Atoi(params["id"])
		pokemonId := id - 1

		if pokemonId <= len(pokeList) - 1 {
			network.SuccessfulResponse(w, pokeList[pokemonId])
		} else {
			network.UnsuccessfulResponse(w, "Invalid index")
		}
	} else {
		network.SuccessfulListResponse(w, pokeList)
	}
}