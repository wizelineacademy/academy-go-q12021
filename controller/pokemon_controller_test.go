package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grethelBello/academy-go-q12021/model"
)

type pokemonDataServiceMock string

func (pdsm pokemonDataServiceMock) Init() error {
	return nil
}
func (pdsm pokemonDataServiceMock) Get(id int) model.Response {
	if id == 1 {
		return errorResponse
	}

	return getPokemonSuccess
}
func (pdsm pokemonDataServiceMock) List(typeFilter model.TypeFilter, items, itemsPerWorker int) model.Response {
	if typeFilter.isOdd() {
		return errorResponse
	}

	return listPokemonSuccess
}

var getPokemonSuccess = model.Response{
	Result: []model.Pokemon{
		{
			Id:   2,
			Name: "ivysaur",
		},
	},
	Total: 2,
	Items: 1,
}
var listPokemonSuccess = model.Response{
	Result: []model.Pokemon{
		{
			Id:   4,
			Name: "bulbasaur",
		},
		{
			Id:   2,
			Name: "ivysaur",
		},
	},
	Total: 2,
	Items: 2,
}
var errorResponse = model.Response{Error: fmt.Errorf("Testing")}
var emptyResponseList = model.Response{
	Result: []model.Pokemon{},
	Total:  0,
	Items:  0,
}

var controller = PokemonController{
	DataService: pokemonDataServiceMock(""),
}
var handler = http.HandlerFunc(controller.GetPokemons)

type testCases struct {
	Name                 string
	RequestPath          string
	ExpectedResponseCode int
	ExpectedResponseBody string
}

var pokemonControllerCases = []testCases{
	{
		Name:                 "Get a pokemon by ID successfully",
		RequestPath:          "/pokemons?id=2",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(getPokemonSuccess),
	},
	{
		Name:                 "Get a pokemon by ID when ID does not exist",
		RequestPath:          "/pokemons?id=1",
		ExpectedResponseCode: http.StatusNotFound,
		ExpectedResponseBody: fmt.Sprint("Testing\n"),
	},
	{
		Name:                 "List all pokemons sucessfully",
		RequestPath:          "/pokemons?page=2",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(listPokemonSuccess),
	},
	{
		Name:                 "List all pokemons when there is not data or an error was triggered",
		RequestPath:          "/pokemons",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(emptyResponseList),
	},
}

func TestPokemonController(t *testing.T) {
	for _, testCase := range pokemonControllerCases {
		fmt.Println(testCase.Name)

		responseRecorder := httptest.NewRecorder()
		request := createRequest(t, testCase.RequestPath)
		handler.ServeHTTP(responseRecorder, request)

		if status := responseRecorder.Code; status != testCase.ExpectedResponseCode {
			t.Errorf("Expected '%v' status code, got '%v'", testCase.ExpectedResponseCode, status)
		}

		var body string
		if testCase.ExpectedResponseCode == http.StatusOK {
			var response model.Response
			json.Unmarshal([]byte(responseRecorder.Body.Bytes()), &response)
			body = fmt.Sprint(response)
		} else {
			body = responseRecorder.Body.String()
		}

		if body != testCase.ExpectedResponseBody {
			t.Errorf("Expected '%q', got '%q'", testCase.ExpectedResponseBody, body)
		}
	}
}

func createRequest(t *testing.T, path string) *http.Request {
	request, createRequestError := http.NewRequest("GET", path, nil)
	if createRequestError != nil {
		t.Errorf("Create request failed: '%v'", createRequestError)
	}
	return request
}
