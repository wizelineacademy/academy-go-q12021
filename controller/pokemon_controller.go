package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/grethelBello/academy-go-q12021/model"
	"github.com/grethelBello/academy-go-q12021/service"
)

var intConvertion = func(value []string) (int, error) {
	if len(value) == 0 || value[0] == "" {
		return 0, errors.New("Empty value cannot be converted to int")
	}

	num, err := strconv.Atoi(value[0])
	if err == nil {
		return num, nil
	}

	return 0, err
}

type queryParams struct {
	Name    string
	Default int
}

var queryParamsList = []queryParams{
	{
		Name: "id",
	},
	{
		Name:    "count",
		Default: 10,
	},
	{
		Name:    "page",
		Default: 1,
	},
}

type PokemonController struct {
	DataService service.DataService
}

func (pc PokemonController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	queryParams := getPokemonsQueryParamas(r)
	fmt.Printf("%v %v: %v\n", r.Method, r.URL.Path, queryParams)

	// Return just one pokemon by ID
	id, ok := queryParams["id"]
	if ok {
		responseGet := pc.DataService.Get(id)
		if responseGet.Error != nil {
			printResponse(r.Method, r.URL.Path, http.StatusNotFound, responseGet)
			http.Error(w, responseGet.Error.Error(), http.StatusNotFound)
		} else {
			printResponse(r.Method, r.URL.Path, http.StatusOK, responseGet)
			json.NewEncoder(w).Encode(responseGet)
		}
		return
	}

	// List all pokemons
	count, _ := queryParams["count"]
	page, _ := queryParams["page"]
	responseList := pc.DataService.List(count, page)
	if responseList.Error != nil {
		response := model.Response{
			Result: make([]model.Pokemon, 0),
			Total:  0,
			Page:   1,
			Count:  count,
		}
		printResponse(r.Method, r.URL.Path, http.StatusOK, response)
		json.NewEncoder(w).Encode(response)
	} else {
		printResponse(r.Method, r.URL.Path, http.StatusOK, responseList)
		json.NewEncoder(w).Encode(responseList)
	}
}

func getPokemonsQueryParamas(r *http.Request) map[string]int {
	queryParams := make(map[string]int, 3)
	query := r.URL.Query()

	for _, nameParam := range queryParamsList {
		valueParam, ok := query[nameParam.Name]
		if convValue, convError := intConvertion(valueParam); ok && convError == nil {
			queryParams[nameParam.Name] = convValue
		} else if nameParam.Default != 0 {
			queryParams[nameParam.Name] = nameParam.Default
		}
	}

	return queryParams
}

func printResponse(method, path string, statusCode int, response model.Response) {
	fmt.Printf("%v %v(%v): %v\n", method, path, statusCode, response)
}

func NewPokemonController() (PokemonController, error) {
	dataService, serviceError := service.NewPokemonDataService()
	if serviceError != nil {
		return PokemonController{}, serviceError
	}
	iniErrror := dataService.Init()
	if iniErrror != nil {
		return PokemonController{}, iniErrror
	}
	return PokemonController{
		DataService: dataService,
	}, nil
}
