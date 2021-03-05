package main

import (
	"fmt"
	"net/http"

	"github.com/ToteEmmanuel/academy-go-q12021/registry"
	"github.com/ToteEmmanuel/academy-go-q12021/router"
	csv_file_reader "github.com/ToteEmmanuel/academy-go-q12021/tools/reader"
	"github.com/urfave/negroni"
)

func main() {
	csvPokeStorage, err := csv_file_reader.NewCsvStorage("./poke.csv")
	if err != nil {
		fmt.Println("No Source for Pokemons!")
		panic(err)
	}
	registry := registry.NewRegistry(csvPokeStorage)
	negi := negroni.Classic()
	router := router.NewRouter(registry.NewAppController())
	negi.UseHandler(router)
	if err := http.ListenAndServe("localhost:3000", negi); err != nil {
		fmt.Printf("Sudden death %v\n", err)
	}
}
