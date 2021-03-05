// nolint
package test

import (
	"os"
	"testing"

	"../config"
)

// Test default values and required env vars
func TestDefaultConfig(t *testing.T) {
	environment := config.Environment
	if environment != "dev" {
		t.Errorf("ENVIRONMENT default should be dev")
	}

	serverPort := config.ServerPort
	if serverPort != "8080" {
		t.Errorf("SERVER_PORT default should be 8080")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("POKEMON_SOURCE should panic when it is not defined")
		}
	}()

	config.PokemonSource
}

// Test assignation for env vars
func TestAssignedConfig(t *testing.T) {
	os.Setenv("SERVER_PORT", "2021")
	os.Setenv("ENVIRONMENT", "prod")
	os.Setenv("POKEMON_SOURCE", "../test.csv")

	environment := config.Environment
	if environment != "prod" {
		t.Errorf("ENVIRONMENT did not change its value")
	}

	serverPort := config.ServerPort
	if serverPort != "2021" {
		t.Errorf("SEVER_PORT did not change its value")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("POKEMON_SOURCE should not trigger panic")
		}
	}()

	pokemonSource := config.PokemonSource
	if pokemonSource != "../test/csv" {
		t.Errorf("POKEMON_SOURCE did not change its value")
	}
}
