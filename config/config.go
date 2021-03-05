package config

import (
	"errors"
	"os"
)

var ServerPort = getServerPort()
var PokemonSource = getPokemonSource()
var Environment = getEnv()

func getServerPort() string {
	serverPort := os.Getenv("SERVER_PORT")
	if len(serverPort) > 0 {
		return serverPort
	}

	return "8080"
}

func getPokemonSource() string {
	pokemonSource := os.Getenv("POKEMON_SOURCE")
	if len(pokemonSource) > 0 {
		return pokemonSource
	}

	panic(errors.New("POKEMON_SOURCE env var is required"))
}

func getEnv() string {
	environment := os.Getenv("ENVIRONMENT")
	if len(environment) > 0 {
		return environment
	}

	return "dev"
}
