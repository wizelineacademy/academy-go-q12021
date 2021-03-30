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

type csvSourceMock string

func (csm csvSourceMock) GetData(csvConfig ...*model.SourceConfig) (*model.Data, error) {
	return csvGetDataMock(csvConfig...)
}

func (csm csvSourceMock) SetData(data *model.Data) error {
	return nil
}

type mockResponseData struct {
	Body          string
	StatusCode    int
	ErrorResponse string
}
type httpClientMock struct{}

func (hcm httpClientMock) Do(request *http.Request) (*http.Response, error) {
	return doMock(request)
}

var doMock func(*http.Request) (*http.Response, error)

var csvGetDataMock func(...*model.SourceConfig) (*model.Data, error)

func TestFailedInit(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		return &model.Data{}, errs.StorageError{TechnicalError: fmt.Errorf("testing")}
	}

	pokemonSource := PokemonDataService{CsvSource: csvSourceMock("")}
	initError := pokemonSource.Init()
	if initError == nil {
		t.Errorf("PokemonDataService expect return error, got '%v'", initError)
	} else if len(pokemonSource.Data) > 0 {
		t.Errorf("PokemonDataService must be empty when return error, got '%v'", pokemonSource)
	}
}

func TestInitWithHeaders(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"ID", "Name"},
			{"1", " bulbasaur"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{})
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

func TestInitWithEmptyLine(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2 ", "Pikachu"},
			{"", ""},
			{"1", "bulbasaur"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{})
	if len(pokemonSource.Data) != 2 {
		t.Errorf("PokemonDataService expects 2 element, got '%v'", pokemonSource)
	} else if len(pokemonSource.keys) != 2 {
		t.Errorf("PokemonDataService expects 2 element, got '%v'", pokemonSource.keys)
	}
}

func TestInitEmptyField(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2 ", "Pikachu"},
			{"1", " "},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{})
	if len(pokemonSource.Data) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource)
	} else if len(pokemonSource.keys) != 1 {
		t.Errorf("PokemonDataService expects 1 element, got '%v'", pokemonSource.keys)
	}
}

func TestGetFromApiSuccess(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{
		Body:       getTestPokemon(),
		StatusCode: http.StatusOK,
	})
	response := pokemonSource.Get(1)
	if response.Error != nil {
		t.Errorf("PokemonSource should return an empty error when there is not pokemons, got '%v'", response.Error)
	}

	pokemon := response.Result[0]
	if pokemon.Id != 1 && pokemon.Name != "bulbasaur" {
		t.Errorf("PokemonSource should return a correct pokemon successfully, got '%v'", pokemon)
	}
}

func TestGetFromCsvSuccess(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"1", "pikachu"},
			{"2", "charmander"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{})
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
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"1", "pikachu"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{
		StatusCode: http.StatusNotFound,
		Body:       "",
	})
	response := pokemonSource.Get(2)
	if response.Error == nil || response.Error.Error() != "The pokemon with 2 ID was not found" {
		t.Errorf("PokemonSource should return an error when the ID does not exist, got '%v'", response.Error)
	}
}

func TestListEmptyError(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{})
	response := pokemonSource.List(10, 1)
	if response.Error == nil || response.Error.Error() != "There are not any pokemons" {
		t.Errorf("PokemonSource should return an empty error when the ID does not exist, got '%v'", response.Error)
	}
}

func TestListCountOutOfLimit(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{})
	response := pokemonSource.List(3, 1)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.Error)
	} else if len(response.Result) != 2 || response.Page > 1 {
		t.Errorf("PokemonSource should return the maximum of pokemons in one page when the count exceds the total, got '%v'", response.Result)
	}
}

func TestListPageOutOfLimit(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{})
	response := pokemonSource.List(1, 3)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.Error)
	} else if len(response.Result) != 1 || response.Page > 2 {
		t.Errorf("PokemonSource should return the maximum of pokemons in one page when the count exceds the total, got '%v'", response.Result)
	}
}

func TestListSuccess(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		csvData := [][]string{
			{"2", "pikachu"},
			{"1", "charmander"},
		}
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{})
	response := pokemonSource.List(1, 1)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.Error)
	}

	pokemon := response.Result[0]
	if len(response.Result) != 1 && pokemon.Id != 1 {
		t.Errorf("PokemonSource should list pokemons in ascendence order successfully, got '%v'", response.Result)
	}
}

func initPokemonSource(t *testing.T, response mockResponseData) *PokemonDataService {
	doMock = func(req *http.Request) (*http.Response, error) {
		if response.ErrorResponse != "" {
			return nil, fmt.Errorf(response.ErrorResponse)
		}

		bodyMock := ioutil.NopCloser(bytes.NewReader([]byte(response.Body)))
		return &http.Response{
			StatusCode: response.StatusCode,
			Body:       bodyMock,
		}, nil
	}

	clientMock := httpClientMock{}
	pokemonSource := &PokemonDataService{
		CsvSource: csvSourceMock(""),
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

func getTestPokemon() string {
	fileContent, error := ioutil.ReadFile("./testdata/pokemon_api.json")
	if error != nil {
		panic(error)
	}
	return string(fileContent)
}
