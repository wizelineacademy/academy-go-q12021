package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var pokemons AllPokemons

func main() {
	openCsv()
	readCsv()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")
	router.HandleFunc("/pokemons/{id}", getPokemon).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

}
