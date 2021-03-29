package config

import (
	"errors"
	"os"

	"github.com/wizelineacademy/academy-go/constant"
)

var defaultValues = map[string]string{
	constant.EnvironmentVarName:    constant.DefaultEnvironment,
	constant.PokemonServiceVarName: constant.DefaultPokemonService,
	constant.ServerPortVarName:     constant.DefaultServerPort,
}

// GetEnvVar returns the value for any environment variable or panic if a default is not defined
func GetEnvVar(envVarName string) (string, error) {
	envVarValue, ok := os.LookupEnv(envVarName)
	if ok {
		return envVarValue, nil
	}

	defaultValue, ok := defaultValues[envVarName]
	if ok {
		return defaultValue, nil
	}

	return "", errors.New(envVarName + " env var is required")
}
