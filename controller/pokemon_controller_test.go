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
	switch id {
	case 1:
		return responses["errorConResponse"]
	case 2:
		return responses["successGetConResponse"]
	case 3:
		return responses["successGetStaReponse"]
	default:
		return responses["errorStaResponse"]
	}
}
func (pdsm pokemonDataServiceMock) List(count, page int) model.Response {
	if page == 1 {
		return responses["successListStaResponse"]
	}

	return responses["errorStaResponse"]
}
func (pdsm pokemonDataServiceMock) Filter(typeFilter model.TypeFilter, items, itemsPerWorker int) model.Response {
	if !typeFilter.IsOdd() {
		return responses["errorConResponse"]
	}

	return responses["successFilterConResponse"]
}

var responses = map[string]model.Response{
	"successGetStaReponse": model.StaticResponse{
		Result: []model.Pokemon{
			{
				Id:   3,
				Name: "venusaur",
			},
		},
		Total: 2,
		Page:  1,
		Count: 1,
	},
	"errorStaResponse": model.StaticResponse{Error: fmt.Errorf("Testing")},
	"successListStaResponse": model.StaticResponse{
		Result: []model.Pokemon{
			{
				Id:   1,
				Name: "bulbasaur",
			},
			{
				Id:   2,
				Name: "ivysaur",
			},
		},
		Total: 2,
		Page:  2,
		Count: 10,
	},
	"emptyListStaResponse": model.StaticResponse{
		Result: make([]model.Pokemon, 0),
		Total:  0,
		Page:   1,
		Count:  10,
	},
	"successGetConResponse": model.ConcurrentResponse{
		Result: []model.Pokemon{
			{
				Id:   2,
				Name: "ivysaur",
			},
		},
		Total: 2,
		Items: 1,
	},
	"successFilterConResponse": model.ConcurrentResponse{
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
	},
	"errorConResponse": model.ConcurrentResponse{Error: fmt.Errorf("Testing")},
	"emptyFilterConResponse": model.ConcurrentResponse{
		Result: []model.Pokemon{},
		Total:  0,
		Items:  0,
	},
}

var controller = PokemonController{
	firstDelivery:  pokemonDataServiceMock(""),
	secondDelivery: pokemonDataServiceMock(""),
	thirdDelivery:  pokemonDataServiceMock(""),
}

type testCases struct {
	Name                 string
	RequestPath          string
	ExpectedResponseCode int
	ExpectedResponseBody string
	Handler              http.HandlerFunc
	Delivery             int
}

var pokemonControllerCases = []testCases{
	{
		Name:                 "First Delivery: Get a pokemon by ID successfully",
		RequestPath:          "/first-delivery/pokemons?id=3",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(responses["successGetStaReponse"]),
		Handler:              http.HandlerFunc(controller.GetCsvPokemons),
		Delivery:             1,
	},
	{
		Name:                 "First Delivery: Get a pokemon by ID when ID does not exist or source is empty",
		RequestPath:          "/first-delivery/pokemons?id=4",
		ExpectedResponseCode: http.StatusNotFound,
		ExpectedResponseBody: fmt.Sprint("Testing\n"),
		Handler:              http.HandlerFunc(controller.GetCsvPokemons),
		Delivery:             1,
	},
	{
		Name:                 "First Delivery: List all pokemons sucessfully with default parameters",
		RequestPath:          "/first-delivery/pokemons",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(responses["successListStaResponse"]),
		Handler:              http.HandlerFunc(controller.GetCsvPokemons),
		Delivery:             1,
	},
	{
		Name:                 "First Delivery: List all pokemons when there is not data or an error is triggered",
		RequestPath:          "/first-delivery/pokemons?page=2",
		ExpectedResponseCode: http.StatusNotFound,
		ExpectedResponseBody: fmt.Sprint("Testing\n"),
		Handler:              http.HandlerFunc(controller.GetCsvPokemons),
		Delivery:             1,
	},
	{
		Name:                 "Second Delivery: Get a pokemon by ID successfully",
		RequestPath:          "/second-delivery/pokemons?id=3",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(responses["successGetStaReponse"]),
		Handler:              http.HandlerFunc(controller.GetDynamicPokemons),
		Delivery:             2,
	},
	{
		Name:                 "Second Delivery: Get a pokemon by ID when ID does not exist",
		RequestPath:          "/second-delivery/pokemons?id=4",
		ExpectedResponseCode: http.StatusNotFound,
		ExpectedResponseBody: fmt.Sprint("Testing\n"),
		Handler:              http.HandlerFunc(controller.GetDynamicPokemons),
		Delivery:             2,
	},
	{
		Name:                 "Second Delivery: List all pokemons sucessfully",
		RequestPath:          "/second-delivery/pokemons",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(responses["successListStaResponse"]),
		Handler:              http.HandlerFunc(controller.GetDynamicPokemons),
		Delivery:             2,
	},
	{
		Name:                 "Second Delivery: List all pokemons when there is not data or an error was triggered",
		RequestPath:          "/second-delivery/pokemons?page=2",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(responses["emptyListStaResponse"]),
		Handler:              http.HandlerFunc(controller.GetDynamicPokemons),
		Delivery:             2,
	},
	{
		Name:                 "Final Delivery: Get a pokemon by ID successfully",
		RequestPath:          "/final-delivery/pokemons?id=2",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(responses["successGetConResponse"]),
		Handler:              http.HandlerFunc(controller.GetCurrentPokemons),
		Delivery:             3,
	},
	{
		Name:                 "Final Delivery: Get a pokemon by ID when ID does not exist",
		RequestPath:          "/final-delivery/pokemons?id=1",
		ExpectedResponseCode: http.StatusNotFound,
		ExpectedResponseBody: fmt.Sprint("Testing\n"),
		Handler:              http.HandlerFunc(controller.GetCurrentPokemons),
		Delivery:             3,
	},
	{
		Name:                 "Final Delivery: Filter pokemons sucessfully",
		RequestPath:          "/final-delivery/pokemons?type=odd&items=1&items_per_worker=4",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(responses["successFilterConResponse"]),
		Handler:              http.HandlerFunc(controller.GetCurrentPokemons),
		Delivery:             3,
	},
	{
		Name:                 "Final Delivery: Filter pokemons when there is not data or an error was triggered",
		RequestPath:          "/final-delivery/pokemons?type=even&items=4&items_per_worker=1",
		ExpectedResponseCode: http.StatusOK,
		ExpectedResponseBody: fmt.Sprint(responses["emptyFilterConResponse"]),
		Handler:              http.HandlerFunc(controller.GetCurrentPokemons),
		Delivery:             3,
	},
	{
		Name:                 "Final Delivery: Filter pokemons bad request",
		RequestPath:          "/final-delivery/pokemons",
		ExpectedResponseCode: http.StatusBadRequest,
		ExpectedResponseBody: fmt.Sprintf("The query params type, items and items_per_worker are required, otherwise you can request just for an id\n"),
		Handler:              http.HandlerFunc(controller.GetCurrentPokemons),
		Delivery:             3,
	},
}

func TestPokemonController(t *testing.T) {
	for _, testCase := range pokemonControllerCases {
		fmt.Println(testCase.Name)

		responseRecorder := httptest.NewRecorder()
		request := createRequest(t, testCase.RequestPath)
		testCase.Handler.ServeHTTP(responseRecorder, request)

		if status := responseRecorder.Code; status != testCase.ExpectedResponseCode {
			t.Errorf("Expected '%v' status code, got '%v'", testCase.ExpectedResponseCode, status)
		}

		parseFunc := parseStaResponse
		if testCase.Delivery == 3 {
			parseFunc = parseConResponse
		}
		var body string
		if testCase.ExpectedResponseCode == http.StatusOK {
			if parseResult, parseError := parseFunc(responseRecorder); parseError != nil {
				t.Errorf("Error parsing response '%v', got '%v'", responseRecorder.Body.String(), parseError)
			} else {
				body = parseResult
			}
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

func parseConResponse(recorder *httptest.ResponseRecorder) (string, error) {
	var response model.ConcurrentResponse
	unmarshallError := json.Unmarshal([]byte(recorder.Body.Bytes()), &response)
	if unmarshallError != nil {
		return "", unmarshallError
	}
	return fmt.Sprint(response), nil
}

func parseStaResponse(recorder *httptest.ResponseRecorder) (string, error) {
	var response model.StaticResponse
	unmarshallError := json.Unmarshal([]byte(recorder.Body.Bytes()), &response)
	if unmarshallError != nil {
		return "", unmarshallError
	}
	return fmt.Sprint(response), nil
}
