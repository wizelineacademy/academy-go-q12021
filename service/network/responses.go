package network

import (
	"bootcamp/domain/model"
	"net/http"
	"encoding/json"
	"fmt"
)

func Response(w http.ResponseWriter, pokemon model.Pokemon, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemon)
}

func ResponseList(w http.ResponseWriter, pokemonList model.PokemonList, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemonList)
}