// nolint
package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/wizelineacademy/academy-go/constant"
)

type envVars struct {
	VarName      string
	DefaultValue string
	TestValue    string
}

var testCases = []envVars{
	{
		VarName:      constant.EnvironmentVarName,
		DefaultValue: constant.DefaultEnvironment,
		TestValue:    "prod",
	},
	{
		VarName:      constant.PokemonServiceVarName,
		DefaultValue: constant.DefaultPokemonService,
		TestValue:    "https://my-url.com",
	},
	{
		VarName:   constant.PokemonSourceVarName,
		TestValue: "../test.csv",
	},
	{
		VarName:      constant.ServerPortVarName,
		DefaultValue: constant.DefaultServerPort,
		TestValue:    "2021",
	},
}

func TestDefaultConfig(t *testing.T) {
	for _, testCase := range testCases {
		if testCase.DefaultValue == "" {
			defer func() {
				if error := recover(); error == nil {
					t.Errorf("It should panic when %v is not defined", testCase.VarName)
				} else {
					fmt.Printf("The config has throwed an exception for %v variable\n", testCase.VarName)
				}
			}()
		}

		envVar := GetEnvVar(testCase.VarName)
		if envVar != testCase.DefaultValue {
			t.Errorf("%v default should be '%v', got '%v'", testCase.VarName, testCase.DefaultValue, envVar)
		}
	}
}

func TestAssignedConfig(t *testing.T) {
	for _, testCase := range testCases {
		os.Setenv(testCase.VarName, testCase.TestValue)

		if testCase.DefaultValue == "" {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%v should not trigger panic", testCase.VarName)
				}
			}()
		}

		envVar := GetEnvVar(testCase.VarName)
		if envVar != testCase.TestValue {
			t.Errorf("%v did not change its value, expected '%v', got '%v'", testCase.VarName, testCase.TestValue, envVar)
		}
	}
}
