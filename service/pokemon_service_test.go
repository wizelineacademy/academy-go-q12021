package service

import (
	"fmt"
	"testing"

	"github.com/wizelineacademy/academy-go/data"
	"github.com/wizelineacademy/academy-go/model"
	"github.com/wizelineacademy/academy-go/model/errs"
)

type dataSourceMock string

func (dsm dataSourceMock) GetData() (data.Data, error) {
	return GetDataMock()
}

var GetDataMock func() (data.Data, error)

func TestFailedInit(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		return data.Data{}, errs.StorageError{TechnicalError: fmt.Errorf("testing")}
	}

	pokemonSource := PokemonDataService(make(map[int]model.Pokemon))
	initError := pokemonSource.Init(dataSourceMock(""))
	if initError == nil {
		t.Errorf("PokemonDataService expect return error, got '%v'", initError)
	} else if len(pokemonSource) > 0 {
		t.Errorf("PokemonDataService must be empty when return error, got '%v'", pokemonSource)
	}
}

func TestInitWithHeaders(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{
			{"ID", "Name"},
			{"1", " bulbasaur"},
		}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	if len(pokemonSource) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource)
	}
	pokemon := pokemonSource[1]
	if len(pokemon.Name) > 9 {
		t.Errorf("Pokemon name expects to not have spaces, got '%v'", pokemon.Name)
	}
}

func TestInitWithEmptyLine(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{
			{"2 ", "Pikachu"},
			{"", ""},
			{"1", "bulbasaur"},
		}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	if len(pokemonSource) != 2 {
		t.Errorf("PokemonDataService expects 2 element, got '%v'", pokemonSource)
	}
}

func TestInitEmptyField(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{
			{"2 ", "Pikachu"},
			{"1", " "},
		}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	if len(pokemonSource) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource)
	}
}

func TestGetEmptyError(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	response := pokemonSource.Get(2)
	if response.Error == nil || response.Error.Error() != "There are not any pokemons" {
		t.Errorf("PokemonSource should return an empty error when there is not pokemons, got '%v'", response.Error)
	}
}

func TestGetSuccess(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{
			{"1", "pikachu"},
			{"2", "charmander"},
		}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	response := pokemonSource.Get(1)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemon successfully, got '%v'", response.Error)
	}

	pokemon := response.Result[0]
	if pokemon.Id != 1 && pokemon.Name != "pikachu" {
		t.Errorf("PokemonSource should return a correct pokemon successfully, got '%v'", pokemon)
	}
}

func TestGetNotFoundError(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{
			{"1", "pikachu"},
		}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	response := pokemonSource.Get(2)
	if response.Error == nil || response.Error.Error() != "The pokemon with 2 ID was not found" {
		t.Errorf("PokemonSource should return an error when the ID does not exist, got '%v'", response.Error)
	}
}

func TestListEmptyError(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	response := pokemonSource.List(10, 1)
	if response.Error == nil || response.Error.Error() != "There are not any pokemons" {
		t.Errorf("PokemonSource should return an empty error when the ID does not exist, got '%v'", response.Error)
	}
}

func TestListCountOutOfLimit(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	response := pokemonSource.List(3, 1)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.Error)
	} else if len(response.Result) != 2 || response.Page > 1 {
		t.Errorf("PokemonSource should return the maximum of pokemons in one page when the count exceds the total, got '%v'", response.Result)
	}
}

func TestListPageOutOfLimit(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t)
	response := pokemonSource.List(1, 3)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.Error)
	} else if len(response.Result) != 1 || response.Page > 2 {
		t.Errorf("PokemonSource should return the maximum of pokemons in one page when the count exceds the total, got '%v'", response.Result)
	}
}

func TestListSuccess(t *testing.T) {
	GetDataMock = func() (data.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := data.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := PokemonDataService(make(map[int]model.Pokemon))
	initError := pokemonSource.Init(dataSourceMock(""))
	if initError != nil {
		t.Errorf("PokemonDataService should not return error, got '%v'", initError)
	}

	response := pokemonSource.List(1, 1)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.Error)
	}

	pokemon := response.Result[0]
	if len(response.Result) != 1 && pokemon.Id != 1 {
		t.Errorf("PokemonSource should list pokemons in ascendence order successfully, got '%v'", response.Result)
	}
}

func initPokemonSource(t *testing.T) PokemonDataService {
	pokemonSource := PokemonDataService(make(map[int]model.Pokemon))
	initError := pokemonSource.Init(dataSourceMock(""))
	if initError != nil {
		t.Errorf("PokemonDataService should not return error, got '%v'", initError)
	}
	return pokemonSource
}
