package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wizelineacademy/academy-go/data"
	"github.com/wizelineacademy/academy-go/model"
	"github.com/wizelineacademy/academy-go/service"
)

type dataSourceMock string

func (dsm dataSourceMock) GetData() (data.Data, error) {
	csvData := getDataMock()
	data := data.NewCsvData(csvData)
	return data, nil
}

var getDataMock func() [][]string
var dataServiceMock service.DataService
var emptyData = [][]string{
	{"2", ""},
	{"1", ""},
}
var validData = [][]string{
	{"2", "pikachu"},
	{"1", "charmander"},
}

type testCases struct {
	Name                 string
	Seed                 [][]string
	RequestPath          string
	ExpectedResponseCode int
	ExpectedResponseBody string
}

var pokemonControllerCases = []testCases{
	{
		Name:                 "Get a pokemon by ID successfully",
		Seed:                 validData,
		RequestPath:          "/pokemons?id=1",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(model.Response{Result: []model.Pokemon{
			{
				Id:   1,
				Name: "charmander",
			},
		},
			Total: 1,
			Page:  1,
			Count: 1,
		}),
	},
	{
		Name:                 "Get a pokemon by ID when source is empty",
		Seed:                 emptyData,
		RequestPath:          "/pokemons?id=1",
		ExpectedResponseCode: http.StatusNotFound,
		ExpectedResponseBody: fmt.Sprint("There are not any pokemons\n"),
	},
	{
		Name:                 "Get a pokemon by ID when ID does not exist",
		Seed:                 validData,
		RequestPath:          "/pokemons?id=3",
		ExpectedResponseCode: http.StatusNotFound,
		ExpectedResponseBody: fmt.Sprint("The pokemon with 3 ID was not found\n"),
	},
	{
		Name:                 "List all pokemons sucessfully",
		Seed:                 validData,
		RequestPath:          "/pokemons",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(model.Response{Result: []model.Pokemon{
			{
				Id:   1,
				Name: "charmander",
			},
			{
				Id:   2,
				Name: "pikachu",
			},
		},
			Total: 2,
			Page:  1,
			Count: 10,
		}),
	},
	{
		Name:                 "List all pokemons when there is not data",
		Seed:                 emptyData,
		RequestPath:          "/pokemons",
		ExpectedResponseCode: http.StatusNotFound,
		ExpectedResponseBody: fmt.Sprint("There are not any pokemons\n"),
	},
}

func TestPokemonController(t *testing.T) {
	for _, testCase := range pokemonControllerCases {
		fmt.Println(testCase.Name)
		getDataMock = func() [][]string {
			return testCase.Seed
		}
		initDataService(t)

		handler := getHandler()
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

func initDataService(t *testing.T) {
	dataServiceMock = service.PokemonDataService(make(map[int]model.Pokemon))
	initDataServiceErr := dataServiceMock.Init(dataSourceMock(""))
	if initDataServiceErr != nil {
		t.Errorf("PokemonDataService init failed: '%v'", initDataServiceErr)
	}
}

func createRequest(t *testing.T, path string) *http.Request {
	request, createRequestError := http.NewRequest("GET", path, nil)
	if createRequestError != nil {
		t.Errorf("Create request failed: '%v'", createRequestError)
	}
	return request
}

func getHandler() http.Handler {
	pokemonRoutes := GetPokemonRoutes(dataServiceMock)
	getPokemons := pokemonRoutes["/pokemons"]
	handler := http.HandlerFunc(getPokemons)
	return handler
}
