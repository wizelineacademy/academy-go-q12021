package main

import (
	"fmt"
	"net/http"

	"github.com/wizelineacademy/academy-go/config"
	"github.com/wizelineacademy/academy-go/constant"
	"github.com/wizelineacademy/academy-go/controller"
)

var serverPort = config.GetEnvVar(constant.PokemonSourceVarName)

func main() {
	pokemonCsvPath := config.GetEnvVar(constant.PokemonSourceVarName)
	pokemonApi := config.GetEnvVar(constant.PokemonServiceVarName)
	fmt.Println(pokemonCsvPath)
	fmt.Println(pokemonApi)
	pokemonController, error := controller.NewPokemonController(pokemonCsvPath, pokemonApi)
	if error != nil {
		panic(error)
	}

	http.HandleFunc("/pokemon", pokemonController.GetPokemons)
	http.ListenAndServe(":"+serverPort, nil)
}
