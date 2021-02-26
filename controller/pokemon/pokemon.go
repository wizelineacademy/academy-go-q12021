package pokemon

import (
	"bootcamp/usecase/pokemon"
	"bootcamp/service/network"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if id := params["id"]; id != "" {
		if !bson.IsObjectIdHex(id) {
			network.UnsuccessfulResponse(w, "Invalid id provided")
			return
		}

		objectId := bson.ObjectIdHex(id)

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