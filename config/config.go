package config

import (
	"errors"
	"os"

	"github.com/wizelineacademy/academy-go/constant"
)

// Environment contains the value of the current environment
var Environment = getEnv()

// PokemonSource contains the value of data source for pokemon endpoint
var PokemonSource = getPokemonSource()

// ServerPort contains the value of server port
var ServerPort = getServerPort()

func getServerPort() string {
	serverPort, ok := os.LookupEnv(constant.ServerPortVarName)
	if ok {
		return serverPort
	}

	return constant.DefaultServerPort
}

func getPokemonSource() string {
	pokemonSource, ok := os.LookupEnv(constant.PokemonSourceVarName)
	if ok {
		return pokemonSource
	}

	panic(errors.New(constant.PokemonSourceVarName + " env var is required"))
}

func getEnv() string {
	environment, ok := os.LookupEnv(constant.EnvironmentVarName)
	if ok {
		return environment
	}

	return constant.DefaultEnvironment
}
