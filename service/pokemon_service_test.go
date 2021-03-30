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

var csvData = [][]string{
	{"ID", "Name"},
	{"2", "ivysaur"},
	{"", ""},
	{"1", " bulbasaur"},
	{"4", "charmander"},
	{"7", "squirtle"},
	{"3", "venusaur"},
	{"5", "charmeleon"},
	{"6", "charizard"},
	{"100", " "},
}

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

func TestInitDirtyData(t *testing.T) {
	pokemonSource := initPokemonSource(t, mockResponseData{}, csvData)
	if len(pokemonSource.Data) != 7 {
		t.Errorf("PokemonDataService expects 7 element, got '%v'", pokemonSource.Data)
	}

	pokemon := pokemonSource.Data[1]
	if len(pokemon.Name) > 9 {
		t.Errorf("Pokemon name expects to not have spaces, got '%v'", pokemon.Name)
	}
}

func TestGetFromApiSuccess(t *testing.T) {
	pokemonSource := initPokemonSource(t, mockResponseData{
		Body:       getTestPokemon(),
		StatusCode: http.StatusOK,
	}, [][]string{})
	response := pokemonSource.Get(1)
	if response.Error != nil {
		t.Errorf("PokemonSource should return an empty error when there is not pokemons, got '%v'", response.Error)
	}

	pokemon := response.Result[0]
	if pokemon.Id != 1 || pokemon.Name != "bulbasaur" {
		t.Errorf("PokemonSource should return a correct pokemon successfully, got '%v'", pokemon)
	}
}

func TestGetFromCsvSuccess(t *testing.T) {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		data := model.NewCsvData(csvData)
		return data, nil
	}

	pokemonSource := initPokemonSource(t, mockResponseData{}, csvData)
	response := pokemonSource.Get(1)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemon successfully, got '%v'", response.Error)
	}

	pokemon := response.Result[0]
	if pokemon.Id != 1 && pokemon.Name != "bulbasaur" {
		t.Errorf("PokemonSource should return a correct pokemon successfully, got '%v'", pokemon)
	}
}

func TestGetNotFoundError(t *testing.T) {
	pokemonSource := initPokemonSource(t, mockResponseData{
		StatusCode: http.StatusNotFound,
		Body:       "",
	}, csvData)
	response := pokemonSource.Get(2000)
	if response.Error == nil || response.Error.Error() != "The pokemon with 2000 ID was not found" {
		t.Errorf("PokemonSource should return an error when the ID does not exist, got '%v'", response.Error)
	}
}

func TestListEmptyError(t *testing.T) {
	pokemonSource := initPokemonSource(t, mockResponseData{}, [][]string{})
	response := pokemonSource.List(model.TypeFilter("odd"), 4, 3)
	if response.Error == nil || response.Error.Error() != "There are not any pokemons" {
		t.Errorf("PokemonSource should return an empty error when the ID does not exist, got '%v'", response.Error)
	}
}

func TestListItemsGraterThanData(t *testing.T) {
	pokemonSource := initPokemonSource(t, mockResponseData{}, csvData)
	response := pokemonSource.List(model.TypeFilter("odd"), 5, 3)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.Error)
	} else if len(response.Result) != 4 {
		t.Errorf("PokemonSource should return all the odd Pokemon's IDs in data without repetition, got '%v'", response.Result)
	}
}

func TestListItemsPerWorkEqualsToAllData(t *testing.T) {
	pokemonSource := initPokemonSource(t, mockResponseData{}, csvData)
	response := pokemonSource.List(model.TypeFilter("even"), 5, 7)
	if response.Error != nil {
		t.Errorf("PokemonSource should return a pokemons list successfully, got '%v'", response.Error)
	} else if len(response.Result) != 4 {
		t.Errorf("PokemonSource should return all the even Pokemon's IDs in data without repetition, got '%v'", response.Result)
	}
}

func initPokemonSource(t *testing.T, response mockResponseData, csvDataRespone [][]string) *PokemonDataService {
	csvGetDataMock = func(csvConfig ...*model.SourceConfig) (*model.Data, error) {
		data := model.NewCsvData(csvDataRespone)
		return data, nil
	}
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
