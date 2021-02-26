package pokemon

import (
	"bootcamp/utils"
	"bootcamp/usecase/pokemon"
	"bootcamp/service/network"
	"net/http"
	"github.com/gorilla/mux"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if id := params["id"]; id != "" {
		objectId, err := utils.GetObjectIdFromParams(params["id"])

		if err != nil {
			network.UnsuccessfulResponse(w, err.Error())
			return
		}

		pokemon, err := pokemon.GetPokemonById(objectId)

		if err != nil {
			network.UnsuccessfulResponse(w, "Could not get information from Pokedex for requested id")
			return
		} else {
			network.SuccessfulResponse(w, pokemon)
		}
	} else {
		pokemonList, err := pokemon.GetPokemon()

		if err != nil {
			network.UnsuccessfulResponse(w, "Could not get information from Pokedex")
			return
		} else {
			network.SuccessfulListResponse(w, pokemonList)
		}
	}
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon.AddPokemon(r.Body)
	if err != nil {
		network.UnsuccessfulResponse(w, "Could not add Pokemon to Pokedex")
		return
	} else {
		network.SuccessfulResponse(w, pokemon)
	}
}