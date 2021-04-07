package controller

import (
	"encoding/json"
	"errors"
	"log"
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
		if typeFilter.IsValid() {
			return typeFilter, nil
		}

		return model.TypeFilter(""), errors.New("Type value not valid: '" + typeFilter.GetVal() + "'")
	},
}

type queryParams struct {
	name       string
	defaultVal string
	convertion func([]string) (interface{}, error)
}

var queryParamsFilterList = []queryParams{
	{
		name:       "id",
		convertion: convertionFunctions[intKey],
	},
	{
		name:       "items",
		convertion: convertionFunctions[intKey],
	},
	{
		name:       "items_per_worker",
		convertion: convertionFunctions[intKey],
	},
	{
		name:       "type",
		convertion: convertionFunctions[typeKey],
	},
}

var queryParamsList = []queryParams{
	{
		name:       "id",
		convertion: convertionFunctions[intKey],
	},
	{
		name:       "count",
		defaultVal: "10",
		convertion: convertionFunctions[intKey],
	},
	{
		name:       "page",
		defaultVal: "1",
		convertion: convertionFunctions[intKey],
	},
}

type PokemonController struct {
	firstDelivery  service.DataService
	secondDelivery service.DataService
	thirdDelivery  service.DataService
}

// GetCsvPokemons returns pokemons contained into a CSV file.
// It returns a list or a specific pokemon if the id query param is present
func (pc PokemonController) GetCsvPokemons(w http.ResponseWriter, r *http.Request) {
	queryParams := getPokemonsQueryParamas(r, queryParamsList)
	log.Printf("%v %v: %v\n", r.Method, r.URL.Path, queryParams)

	// Return just one pokemon by ID
	id, ok := queryParams["id"]
	if ok {
		responseGet := pc.firstDelivery.Get(id.(int))
		if responseGet.GetError() != nil {
			http.Error(w, responseGet.GetError().Error(), http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(responseGet)
		}
		return
	}

	// List all pokemons
	count, _ := queryParams["count"]
	page, _ := queryParams["page"]
	responseList := pc.firstDelivery.List(count.(int), page.(int))
	if responseList.GetError() != nil {
		http.Error(w, responseList.GetError().Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(responseList)
	}
}

// GetDynamicPokemons returns pokemons contained into a CSV file.
// It returns a list or a specific pokemon if the id query param is present
// If the pokemon does not exist in the CSV file, it is looked for in a REST API.
func (pc PokemonController) GetDynamicPokemons(w http.ResponseWriter, r *http.Request) {
	queryParams := getPokemonsQueryParamas(r, queryParamsList)
	log.Printf("%v %v: %v\n", r.Method, r.URL.Path, queryParams)

	// Return just one pokemon by ID
	id, ok := queryParams["id"]
	if ok {
		responseGet := pc.secondDelivery.Get(id.(int))
		if responseGet.GetError() != nil {
			printResponse(r.Method, r.URL.Path, http.StatusNotFound, responseGet)
			http.Error(w, responseGet.GetError().Error(), http.StatusNotFound)
		} else {
			printResponse(r.Method, r.URL.Path, http.StatusOK, responseGet)
			json.NewEncoder(w).Encode(responseGet)
		}
		return
	}

	// List all pokemons
	count, _ := queryParams["count"]
	page, _ := queryParams["page"]
	responseList := pc.secondDelivery.List(count.(int), page.(int))
	if responseList.GetError() != nil {
		response := model.StaticResponse{
			Result: make([]model.Pokemon, 0),
			Total:  0,
			Page:   1,
			Count:  count.(int),
		}
		printResponse(r.Method, r.URL.Path, http.StatusOK, response)
		json.NewEncoder(w).Encode(response)
	} else {
		printResponse(r.Method, r.URL.Path, http.StatusOK, responseList)
		json.NewEncoder(w).Encode(responseList)
	}
}

// GetCurrentPokemons returns pokemons contained into a CSV file.
// It returns a filtered list or a specific pokemon if the id query param is present
func (pc PokemonController) GetCurrentPokemons(w http.ResponseWriter, r *http.Request) {
	queryParams := getPokemonsQueryParamas(r, queryParamsFilterList)
	log.Printf("%v %v: %v\n", r.Method, r.URL.Path, queryParams)

	// Return just one pokemon by ID
	id, ok := queryParams[queryParamsList[0].name]
	if ok {
		responseGet := pc.thirdDelivery.Get(id.(int))
		if responseGet.GetError() != nil {
			printResponse(r.Method, r.URL.Path, http.StatusNotFound, responseGet)
			http.Error(w, responseGet.GetError().Error(), http.StatusNotFound)
		} else {
			printResponse(r.Method, r.URL.Path, http.StatusOK, responseGet)
			json.NewEncoder(w).Encode(responseGet)
		}
		return
	}

	// Filter pokemons
	responseItems, okItems := queryParams[queryParamsFilterList[1].name]
	itemsPerWorker, okItemsWorker := queryParams[queryParamsFilterList[2].name]
	typeFilter, okFilter := queryParams[queryParamsFilterList[3].name]
	log.Printf("%v %v %v", okFilter, okItems, okItemsWorker)
	if !okFilter || !okItems || !okItemsWorker {
		response := model.ConcurrentResponse{Error: errors.New("The query params type, items and items_per_worker are required, otherwise you can request just for an id")}
		printResponse(r.Method, r.URL.Path, http.StatusBadRequest, response)
		http.Error(w, response.GetError().Error(), http.StatusBadRequest)
		return
	}

	responseList := pc.thirdDelivery.Filter(typeFilter.(model.TypeFilter), responseItems.(int), itemsPerWorker.(int))
	if responseList.GetError() != nil {
		response := model.ConcurrentResponse{Result: []model.Pokemon{}}
		printResponse(r.Method, r.URL.Path, http.StatusOK, response)
		json.NewEncoder(w).Encode(response)
	} else {
		printResponse(r.Method, r.URL.Path, http.StatusOK, responseList)
		json.NewEncoder(w).Encode(responseList)
	}
}

func (pc PokemonController) syncServices() {
	pc.firstDelivery.Sync()
	pc.secondDelivery.Sync()
	pc.thirdDelivery.Sync()
}

func getPokemonsQueryParamas(r *http.Request, requiredParams []queryParams) map[string]interface{} {
	queryParams := make(map[string]interface{}, 3)
	query := r.URL.Query()

	for _, param := range requiredParams {
		valueParam, ok := query[param.name]
		if !ok && param.defaultVal != "" {
			valueParam = []string{param.defaultVal}
		}

		if convValue, convError := param.convertion(valueParam); convError == nil {
			queryParams[param.name] = convValue
		}
	}

	return queryParams
}

func printResponse(method, path string, statusCode int, response model.Response) {
	log.Printf("%v %v(%v): %v\n", method, path, statusCode, response)
}

func NewPokemonController() (PokemonController, error) {
	serviceInstances := []func() (service.DataService, error){
		service.NewStaticPokemonDataService,
		service.NewRestPokemonDataService,
		service.NewPokemonDataService,
	}

	services := [3]service.DataService{}
	for index, sc := range serviceInstances {
		service, instanceError := sc()
		if instanceError != nil {
			return PokemonController{}, instanceError
		}
		initError := service.Init()
		if initError != nil {
			return PokemonController{}, initError
		}
		services[index] = service
	}

	return PokemonController{
		firstDelivery:  services[0],
		secondDelivery: services[1],
		thirdDelivery:  services[2],
	}, nil
}
