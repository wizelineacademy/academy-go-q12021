// nolint
package config

import (
	"os"
	"testing"

	"github.com/grethelBello/academy-go-q12021/constant"
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
		envVar, envError := GetEnvVar(testCase.VarName)
		if testCase.DefaultValue == "" && envError == nil {
			t.Errorf("%v should return an error when it is not defined, got '%v' as variable value and '%v' as error", testCase.VarName, envVar, envError)
		} else if testCase.DefaultValue != "" && envError != nil {
			t.Errorf("%v default should return a default value, got '%v' instance of '%v'", testCase.VarName, envVar, testCase.DefaultValue)
		} else if envVar != testCase.DefaultValue {
			t.Errorf("%v default should be '%v', got '%v'", testCase.VarName, testCase.DefaultValue, envVar)
		}
	}
}

func TestAssignedConfig(t *testing.T) {
	for _, testCase := range testCases {
		os.Setenv(testCase.VarName, testCase.TestValue)

		envVar, envError := GetEnvVar(testCase.VarName)
		if envError != nil {
			t.Errorf("%v should not return an error, got '%v'", testCase.VarName, envError)
		} else if envVar != testCase.TestValue {
			t.Errorf("%v did not change its value, expected '%v', got '%v'", testCase.VarName, testCase.TestValue, envVar)
		}
	}
}
