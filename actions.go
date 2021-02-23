package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	pokemonId := params["id"]
	fmt.Fprintf(w, "Getting pokemon: %s", pokemonId)
}

func GetPokemonList(w http.ResponseWriter, r *http.Request) {
	pokeList := ReadCSV()
	fmt.Println(pokeList)
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pokeList)
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Pokemon
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	fmt.Println(data)

	pokeList := ReadCSV()
	pokeList = append(pokeList, data)
	fmt.Println(pokeList)
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pokeList)
}