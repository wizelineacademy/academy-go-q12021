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

const intKey = "int"
const typeKey = "typeFilter"

var convertionFunctions = map[string]func([]string) (interface{}, error){
	intKey: func(value []string) (interface{}, error) {
		if len(value) == 0 || value[0] == "" {
			return 0, errors.New("Empty value cannot be converted to int")
		}

		num, err := strconv.Atoi(value[0])
		if err == nil {
			return num, nil
		}

		return 0, err
	},
	typeKey: func(value []string) (interface{}, error) {
		if len(value) == 0 || value[0] == "" {
			return model.TypeFilter(""), errors.New("Empty value cannot be converted to model.TypeFilter")
		}

		typeFilter := model.TypeFilter(value[0])
		if typeFilter.isValid() {
			return typeFilter, nil
		}

		return model.TypeFilter(""), errors.New("Type value not valid: '%v'", typeFilter)
	},
}

type queryParams struct {
	name       string
	defaultVal string
	convertion func([]string) (interface{}, error)
}

var queryParamsList = []queryParams{
	{
		name:       "id",
		convertion: convertionFunctions[intKey],
	},
	{
		name:       "items",
		defaultVal: "10",
		convertion: convertionFunctions[intKey],
	},
	{
		name:       "items_per_worker",
		defaultVal: "10",
		convertion: convertionFunctions[intKey],
	},
	{
		name:       "type",
		defaultVal: model.Odd(),
		convertion: convertionFunctions[typeKey],
	},
}

type PokemonController struct {
	DataService service.DataService
}

func (pc PokemonController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	queryParams := getPokemonsQueryParamas(r)
	fmt.Printf("%v %v: %v\n", r.Method, r.URL.Path, queryParams)

	// Return just one pokemon by ID
	id, ok := queryParams[queryParamsList[0].name]
	if ok {
		responseGet := pc.DataService.Get(id.(int))
		if responseGet.Error != nil {
			printResponse(r.Method, r.URL.Path, http.StatusNotFound, responseGet)
			http.Error(w, responseGet.Error.Error(), http.StatusNotFound)
		} else {
			printResponse(r.Method, r.URL.Path, http.StatusOK, responseGet)
			json.NewEncoder(w).Encode(responseGet)
		}
		return
	}

	// Filter pokemons
	typeFilter, _ := queryParams[queryParamsList[3].name]
	responseItems, _ := queryParams[queryParamsList[1].name]
	itemsPerWorker, _ := queryParams[queryParamsList[2].name]
	responseList := pc.DataService.List(typeFilter.(model.TypeFilter), responseItems.(int), itemsPerWorker.(int))
	if responseList.Error != nil {
		response := model.Response{
			Result: []model.Pokemon{},
			Total:  0,
			Items:  0,
		}
		printResponse(r.Method, r.URL.Path, http.StatusOK, response)
		json.NewEncoder(w).Encode(response)
	} else {
		printResponse(r.Method, r.URL.Path, http.StatusOK, responseList)
		json.NewEncoder(w).Encode(responseList)
	}
}

func getPokemonsQueryParamas(r *http.Request) map[string]interface{} {
	queryParams := make(map[string]interface{}, 3)
	query := r.URL.Query()

	for _, param := range queryParamsList {
		valueParam, ok := query[param.name]
		if !ok && param.defaultVal != "" {
			valueParam[0] = param.defaultVal
		}
		if convValue, convError := param.convertion(valueParam); ok && convError == nil {
			queryParams[param.name] = convValue
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
