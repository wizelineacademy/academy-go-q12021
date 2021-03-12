// nolint
package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/wizelineacademy/academy-go/constant"
)

const serverPortValue = "2021"
const environmentValue = "prod"
const pokemonSourceValue = "../test.csv"

// Test default values and required env vars
func TestDefaultConfig(t *testing.T) {
	environment := Environment
	if environment != constant.DefaultEnvironment {
		t.Errorf("%v default should be %v, got %v", constant.EnvironmentVarName, constant.DefaultEnvironment, environment)
	}

	serverPort := ServerPort
	if serverPort != constant.DefaultServerPort {
		t.Errorf("%v default should be %v, got %v", constant.ServerPortVarName, constant.DefaultServerPort, serverPort)
	}

	defer func() {
		if error := recover(); error == nil {
			t.Errorf("It should panic when %v is not defined", constant.PokemonSourceVarName)
		} else {
			fmt.Println("The config has throwed an exception")
		}
	}()

	fmt.Println(PokemonSource)
}

// Test assignation for env vars
func TestAssignedConfig(t *testing.T) {
	os.Setenv(constant.ServerPortVarName, serverPortValue)
	os.Setenv(constant.EnvironmentVarName, environmentValue)
	os.Setenv(constant.PokemonSourceVarName, pokemonSourceValue)

	environment := Environment
	if environment != environmentValue {
		t.Errorf("%v did not change its value, expected '%v', got '%v'", constant.EnvironmentVarName, environmentValue, environment)
	}

	serverPort := ServerPort
	if serverPort != serverPortValue {
		t.Errorf("%v did not change its value, expected '%v', got '%v'", constant.ServerPortVarName, serverPortValue, serverPort)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("%v should not trigger panic", constant.PokemonSourceVarName)
		}
	}()

	pokemonSource := PokemonSource
	if pokemonSource != pokemonSourceValue {
		t.Errorf("%v did not change its value, expected '%v', got '%v'", constant.PokemonSourceVarName, pokemonSourceValue, pokemonSource)
	}
}
