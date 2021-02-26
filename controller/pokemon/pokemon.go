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
			network.UnsuccessfulResponse(w, err.Error())
			return
		}
		network.SuccessfulResponse(w, pokemon)
	} else {
		pokemonList, err := pokemon.GetPokemon()

		if err != nil {
			network.UnsuccessfulResponse(w, err.Error())
			return
		}
		network.SuccessfulListResponse(w, pokemonList)
	}
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon.AddPokemon(r.Body)
	if err != nil {
		network.UnsuccessfulResponse(w,  err.Error())
		return
	}
	network.SuccessfulResponse(w, pokemon)
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	objectId, err := utils.GetObjectIdFromParams(params["id"])

	if err != nil {
		network.UnsuccessfulResponse(w, err.Error())
		return
	}

	pokemon, err := pokemon.UpdatePokemon(objectId, r.Body)

	if err != nil {
		network.UnsuccessfulResponse(w, err.Error())
		return
	}
	network.SuccessfulResponse(w, pokemon)
}