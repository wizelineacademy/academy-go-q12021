package modules

import (
	"models"
	"utils"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func GetPokemonCsv(w http.ResponseWriter, r *http.Request) {
	pokeList := utils.ReadCSV()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println("Cannot get id from params")
	}

	pokemonId := id - 1

	if pokemonId <= len(pokeList) - 1 {
		json.NewEncoder(w).Encode(pokeList[pokemonId])
	} else {
		fmt.Fprintf(w, "There is no information for given id")
	}
}

func GetPokemonListCsv(w http.ResponseWriter, r *http.Request) {
	pokeList := utils.ReadCSV()
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pokeList)
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data models.Pokemon
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	fmt.Println(data)

	pokeList := utils.ReadCSV()
	pokeList = append(pokeList, data)
	fmt.Println(pokeList)
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pokeList)
}