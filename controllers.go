package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const uri string = "https://pokeapi.co/api/v2/pokemon?limit=5&offset=300"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Poke-API")
}

func getPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons = nil
	readCsv()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	pokemons = nil
	readCsv()
	vars := mux.Vars(r)
	pokemonID, err := strconv.Atoi(vars["id"])
	var isPokemon bool

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, pokemon := range pokemons {
		if pokemon.ID == pokemonID {
			isPokemon = true
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pokemon)
		}
	}

	if !isPokemon {
		fmt.Fprintf(w, "There is no pokemon with ID: %d", pokemonID)
	}
}

func getPokemonFromAPI(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Something happened", err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject PokemonExternal
	json.Unmarshal(bodyBytes, &responseObject)

	fmt.Printf("API Response as struct %+v\n", responseObject)
	addLineCsv(responseObject.Results)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to my Poke-API")
}
