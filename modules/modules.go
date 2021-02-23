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

func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	return w
}

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
		w = setHeaders(w)
		json.NewEncoder(w).Encode(pokeList[pokemonId])
	} else {
		fmt.Fprintf(w, "There is no information for given id")
	}
}

func GetPokemonListCsv(w http.ResponseWriter, r *http.Request) {
	pokeList := utils.ReadCSV()
	w = setHeaders(w)
	json.NewEncoder(w).Encode(pokeList)
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	var data models.Pokemon
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	pokeList := utils.ReadCSV()
	pokeList = append(pokeList, data)

	w = setHeaders(w)
	json.NewEncoder(w).Encode(pokeList)
}