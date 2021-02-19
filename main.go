package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Pokemon struct {
	ID     int    `json:ID`
	Name   string `json:Name`
	Weight string `json:Weight`
	Height string `json:Height`
}

type allPokemons []Pokemon

var pokemons allPokemons

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Poke-API")
}

func getPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
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

func openCsv() {
	f, err := os.Open("./csv/pokemon.csv")
	if err != nil {
		fmt.Printf("There was an error opening the file: %v\n", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("There was an error reading something in line: %v\n", err)
		}
		tempPokemon := Pokemon{
			Name:   line[1],
			Weight: line[2],
			Height: line[3],
		}

		if line[0] != "" {
			id, err := strconv.Atoi(line[0])
			if err != nil {
				fmt.Printf("There was an error trying to process the ID: %v\n", err)
			}
			tempPokemon.ID = id
		}

		pokemons = append(pokemons, tempPokemon)
	}
}

func main() {
	openCsv()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")
	router.HandleFunc("/pokemons/{id}", getPokemon).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

}
