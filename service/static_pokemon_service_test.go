package service

import (
	"fmt"
	"testing"

	"github.com/grethelBello/academy-go-q12021/model"
	"github.com/grethelBello/academy-go-q12021/model/errs"
)

type dataSourceMock string

func (dsm dataSourceMock) GetData(...*model.SourceConfig) (*model.Data, error) {
	return GetStaticDataMock()
}
func (dsm dataSourceMock) SetData(*model.Data) error {
	return nil
}

var GetStaticDataMock func() (*model.Data, error)

func TestStaticFailedInit(t *testing.T) {
	GetStaticDataMock = func() (*model.Data, error) {
		return &model.Data{}, errs.StorageError{TechnicalError: fmt.Errorf("testing")}
	}

	pokemonSource := StaticPokemonDataService{
		csvSource: dataSourceMock(""),
	}
	initError := pokemonSource.Init()
	if initError == nil {
		t.Errorf("PokemonDataService expect return error, got '%v'", initError)
	} else if len(pokemonSource.data) > 0 {
		t.Errorf("PokemonDataService must be empty when return error, got '%v'", pokemonSource)
	}
}

func TestInitWithHeaders(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{
		{"ID", "Name"},
		{"1", " bulbasaur"},
	})
	if len(pokemonSource.data) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource)
	}
	pokemon := pokemonSource.data[1]
	if len(pokemon.Name) > 9 {
		t.Errorf("Pokemon name expects to not have spaces, got '%v'", pokemon.Name)
	}
}

func TestInitWithEmptyLine(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{
		{"2 ", "Pikachu"},
		{"", ""},
		{"1", "bulbasaur"},
	})
	if len(pokemonSource.data) != 2 {
		t.Errorf("PokemonDataService expects 2 element, got '%v'", pokemonSource)
	}
}

func TestInitEmptyField(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{
		{"2 ", "Pikachu"},
		{"1", " "},
	})
	if len(pokemonSource.data) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource)
	}
}

func TestGetEmptyError(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{})
	response := pokemonSource.Get(2)
	if response.GetError() == nil || response.GetError().Error() != "There are not any pokemons" {
		t.Errorf("PokemonSource should return an empty error when there is not pokemons, got '%v'", response.GetError())
	}
}

func TestGetSuccess(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{
		{"1", "pikachu"},
		{"2", "charmander"},
	})
	response := pokemonSource.Get(1)
	if response.GetError() != nil {
		t.Errorf("PokemonSource should return a pokemon successfully, got '%v'", response.GetError())
	}

	pokemon := response.(model.StaticResponse).Result[0]
	if pokemon.Id != 1 && pokemon.Name != "pikachu" {
		t.Errorf("PokemonSource should return a correct pokemon successfully, got '%v'", pokemon)
	}
}

func TestStaticGetNotFoundError(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{
		{"1", "pikachu"},
	})
	response := pokemonSource.Get(2)
	if response.GetError() == nil || response.GetError().Error() != "The pokemon with 2 ID was not found" {
		t.Errorf("PokemonSource should return an error when the ID does not exist, got '%v'", response.GetError())
	}
}

func TestListEmptyError(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{})
	response := pokemonSource.List(10, 1)
	if response.GetError() == nil || response.GetError().Error() != "There are not any pokemons" {
		t.Errorf("PokemonSource should return an empty error when the ID does not exist, got '%v'", response.GetError())
	}
}

func TestListCountOutOfLimit(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{
		{"2", "pikachu"},
		{"1", "charmander"},
	})
	response := pokemonSource.List(3, 1)
	responseError := response.GetError()
	if responseError != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", responseError)
	} else if len(response.(model.StaticResponse).Result) != 2 || response.(model.StaticResponse).Page > 1 {
		t.Errorf("PokemonSource should return the maximum of pokemons in one page when the count exceds the total, got '%v'", response.(model.StaticResponse).Result)
	}
}

func TestListPageOutOfLimit(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{
		{"2", "pikachu"},
		{"1", "charmander"},
	})
	response := pokemonSource.List(1, 3)
	responseError := response.GetError()
	if responseError != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", responseError)
	} else if len(response.(model.StaticResponse).Result) != 1 || response.(model.StaticResponse).Page > 2 {
		t.Errorf("PokemonSource should return the maximum of pokemons in one page when the count exceds the total, got '%v'", response.(model.StaticResponse).Result)
	}
}

func TestListSuccess(t *testing.T) {
	pokemonSource := initStaticPokemonSource(t, [][]string{
		{"2", "pikachu"},
		{"1", "charmander"},
	})
	response := pokemonSource.List(1, 1)
	responseError := response.GetError()
	if responseError != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", responseError)
	}

	pokemons := response.(model.StaticResponse).Result
	if len(pokemons) != 1 || pokemons[0].Id != 1 {
		t.Errorf("PokemonSource should list pokemons in ascendence order successfully, got '%v'", pokemons)
	}
}

func initStaticPokemonSource(t *testing.T, csvData [][]string) StaticPokemonDataService {
	GetStaticDataMock = func() (*model.Data, error) {
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := StaticPokemonDataService{
		csvSource: dataSourceMock(""),
	}
	initError := pokemonSource.Init()
	if initError != nil {
		t.Errorf("PokemonDataService should not return error, got '%v'", initError)
	}
	return pokemonSource
}
