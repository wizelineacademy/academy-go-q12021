package network

import (
	"bootcamp/domain/model"
	"net/http"
	"encoding/json"
	"fmt"
	"errors"
)

func SuccessfulListResponse (w http.ResponseWriter, pokemonList model.PokemonList) {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemonList)
	return
}

func SuccessfulResponse (w http.ResponseWriter, pokemon model.Pokemon) {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemon)
	return
}

func UnsuccessfulResponse (w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusBadRequest)
	fmt.Println(errors.New(err))
	fmt.Fprintf(w, err)
	return
}