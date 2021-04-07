package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/grethelBello/academy-go-q12021/data"
	"github.com/grethelBello/academy-go-q12021/model"
	"github.com/grethelBello/academy-go-q12021/model/errs"
)

var csvRestGetDataMock func(...*model.SourceConfig) (*model.Data, error)

type secndDelSrcMock string

func (csm secndDelSrcMock) GetData(csvConfig ...*model.SourceConfig) (*model.Data, error) {
	return csvRestGetDataMock(csvConfig...)
}

func (csm secndDelSrcMock) SetData(data *model.Data) error {
	return nil
}

var doRestMock func(*http.Request) (*http.Response, error)

type httpCltMock struct{}

func (hcm httpCltMock) Do(request *http.Request) (*http.Response, error) {
	return doRestMock(request)
}

func TestRestServiceFailedInit(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		return &model.Data{}, errs.StorageError{TechnicalError: fmt.Errorf("testing")}
	}

	pokemonSource := RestPokemonDataService{CsvSource: secndDelSrcMock("")}
	initError := pokemonSource.Init()
	if initError == nil {
		t.Errorf("PokemonDataService expect return error, got '%v'", initError)
	} else if len(pokemonSource.Data) > 0 {
		t.Errorf("PokemonDataService must be empty when return error, got '%v'", pokemonSource)
	}
}

func TestRestInitWithHeaders(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"ID", "Name"},
			{"1", " bulbasaur"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{})
	if len(pokemonSource.Data) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource.Data)
	} else if len(pokemonSource.keys) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource.keys)
	}
	pokemon := pokemonSource.Data[1]
	if len(pokemon.Name) > 9 {
		t.Errorf("Pokemon name expects to not have spaces, got '%v'", pokemon.Name)
	}
}

func TestRestInitWithEmptyLine(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2 ", "Pikachu"},
			{"", ""},
			{"1", "bulbasaur"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{})
	if len(pokemonSource.Data) != 2 {
		t.Errorf("PokemonDataService expects 2 element, got '%v'", pokemonSource)
	} else if len(pokemonSource.keys) != 2 {
		t.Errorf("PokemonDataService expects 2 element, got '%v'", pokemonSource.keys)
	}
}

func TestRestInitEmptyField(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2 ", "Pikachu"},
			{"1", " "},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{})
	if len(pokemonSource.Data) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource)
	} else if len(pokemonSource.keys) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource.keys)
	}
}

func TestRestGetFromApiSuccess(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{
		Body:       getTestPokemon(),
		StatusCode: http.StatusOK,
	})
	response := pokemonSource.Get(1)
	if response.GetError() != nil {
		t.Errorf("PokemonSource should return an empty error when there is not pokemons, got '%v'", response.GetError())
	}

	pokemon := response.(model.StaticResponse).Result[0]
	if pokemon.Id != 1 && pokemon.Name != "bulbasaur" {
		t.Errorf("PokemonSource should return a correct pokemon successfully, got '%v'", pokemon)
	}
}

func TestRestGetFromCsvSuccess(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"1", "pikachu"},
			{"2", "charmander"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{})
	response := pokemonSource.Get(1)
	if response.GetError() != nil {
		t.Errorf("PokemonSource should return a pokemon successfully, got '%v'", response.GetError())
	}

	pokemon := response.(model.StaticResponse).Result[0]
	if pokemon.Id != 1 && pokemon.Name != "pikachu" {
		t.Errorf("PokemonSource should return a correct pokemon successfully, got '%v'", pokemon)
	}
}

func TestRestGetNotFoundError(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"1", "pikachu"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{
		StatusCode: http.StatusNotFound,
		Body:       "",
	})
	response := pokemonSource.Get(2)
	if response.GetError() == nil || response.GetError().Error() != "The pokemon with 2 ID was not found" {
		t.Errorf("PokemonSource should return an error when the ID does not exist, got '%v'", response.GetError())
	}
}

func TestRestListEmptyError(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{})
	response := pokemonSource.List(10, 1)
	if response.GetError() == nil || response.GetError().Error() != "There are not any pokemons" {
		t.Errorf("PokemonSource should return an empty error when the ID does not exist, got '%v'", response.GetError())
	}
}

func TestRestListCountOutOfLimit(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{})
	response := pokemonSource.List(3, 1)
	if response.GetError() != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.GetError())
	} else if len(response.(model.StaticResponse).Result) != 2 || response.(model.StaticResponse).Page > 1 {
		t.Errorf("PokemonSource should return the maximum of pokemons in one page when the count exceds the total, got '%v'", response.(model.StaticResponse).Result)
	}
}

func TestRestListPageOutOfLimit(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{})
	response := pokemonSource.List(1, 3)
	if response.GetError() != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.GetError())
	} else if len(response.(model.StaticResponse).Result) != 1 || response.(model.StaticResponse).Page > 2 {
		t.Errorf("PokemonSource should return the maximum of pokemons in one page when the count exceds the total, got '%v'", response.(model.StaticResponse).Result)
	}
}

func TestRestListSuccess(t *testing.T) {
	csvRestGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initRestPokemonSource(t, mockResponseData{})
	response := pokemonSource.List(1, 1)
	if response.GetError() != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.GetError())
	}

	pokemons := response.(model.StaticResponse).Result
	if len(pokemons) != 1 && pokemons[0].Id != 1 {
		t.Errorf("PokemonSource should list pokemons in ascendence order successfully, got '%v'", pokemons)
	}
}

func initRestPokemonSource(t *testing.T, response mockResponseData) *RestPokemonDataService {
	doRestMock = func(req *http.Request) (*http.Response, error) {
		if response.ErrorResponse != "" {
			return nil, fmt.Errorf(response.ErrorResponse)
		}

		bodyMock := ioutil.NopCloser(bytes.NewReader([]byte(response.Body)))
		return &http.Response{
			StatusCode: response.StatusCode,
			Body:       bodyMock,
		}, nil
	}

	clientMock := httpCltMock{}
	pokemonSource := &RestPokemonDataService{
		CsvSource: secndDelSrcMock(""),
		HttpSource: data.HttpSource{
			Data: model.HttpData{
				Url:    "https://test.com",
				Method: http.MethodGet,
			},
			Client: clientMock,
		},
	}
	initError := pokemonSource.Init()
	if initError != nil {
		t.Errorf("PokemonDataService should not return error, got '%v'", initError)
	}
	return pokemonSource
}
