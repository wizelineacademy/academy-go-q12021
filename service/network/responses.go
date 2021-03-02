package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bootcamp/domain/model"
)

func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return w
}

func validateError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
}

/*
Response returns a JSON Pokemon
*/
func Response(w http.ResponseWriter, pokemon model.Pokemon, err error) {
	if err != nil {
		validateError(w, err)
	} else {
		w = setHeaders(w)
		json.NewEncoder(w).Encode(pokemon)	
	}
}

/*
Response returns a JSON PokemonList
*/
func ResponseList(w http.ResponseWriter, pokemonList model.PokemonList, err error) {
	if err != nil {
		validateError(w, err)
	} else {
		w = setHeaders(w)
		json.NewEncoder(w).Encode(pokemonList)
	}
}