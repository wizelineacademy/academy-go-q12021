package config

import (
	"errors"
	"os"

	"github.com/wizelineacademy/academy-go/constant"
)

// GetServerPort returns the value of server port or '8080' by default if 'SERVER_PORT' is not defined
func GetServerPort() string {
	serverPort, ok := os.LookupEnv(constant.ServerPortVarName)
	if ok {
		return serverPort
	}

	return constant.DefaultServerPort
}

// GetPokemonSource returns the value of data source for pokemon endpoint and panic when 'POKEMON_SOURCE' is not defined
func GetPokemonSource() string {
	pokemonSource, ok := os.LookupEnv(constant.PokemonSourceVarName)
	if ok {
		return pokemonSource
	}

	panic(errors.New(constant.PokemonSourceVarName + " env var is required"))
}

// GetEnv returns the value of the current environment or 'dev' by default if 'ENVIRONMENT' is not defined
func GetEnv() string {
	environment, ok := os.LookupEnv(constant.EnvironmentVarName)
	if ok {
		return environment
	}

	return constant.DefaultEnvironment
}
