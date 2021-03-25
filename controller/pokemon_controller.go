package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/wizelineacademy/academy-go/model"
	"github.com/wizelineacademy/academy-go/service"
)

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
			http.Error(w, responseGet.Error.Error(), http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(responseGet)
		}
		return
	}

	// List all pokemons
	count, _ := queryParams["count"]
	page, _ := queryParams["page"]
	responseList := pc.DataService.List(count, page)
	if responseList.Error != nil {
		json.NewEncoder(w).Encode(model.Response{
			Result: make([]model.Pokemon, 0),
			Total:  0,
			Page:   1,
			Count:  count,
		})
	} else {
		json.NewEncoder(w).Encode(responseList)
	}
}

func getPokemonsQueryParamas(r *http.Request) map[string]int {
	queryParams := make(map[string]int, 3)

	queryID, ok := r.URL.Query()["id"]
	if ok && len(queryID[0]) > 0 {
		id, err := strconv.Atoi(queryID[0])
		if err == nil {
			queryParams["id"] = id
		}
	}

	count := 10
	queryCount, ok := r.URL.Query()["count"]
	if ok && len(queryCount[0]) > 0 {
		convCount, err := strconv.Atoi(queryCount[0])
		if err == nil {
			count = convCount
		}
	}
	queryParams["count"] = count

	page := 1
	queryPage, ok := r.URL.Query()["page"]
	if ok && len(queryPage[0]) > 0 {
		convPage, err := strconv.Atoi(queryPage[0])
		if err == nil {
			page = convPage
		}
	}
	queryParams["page"] = page

	return queryParams
}

func NewPokemonController(csvPath, apiEndpoint string) (PokemonController, error) {
	dataService := service.NewPokemonDataService(csvPath, apiEndpoint)
	iniErrror := dataService.Init()
	if iniErrror != nil {
		return PokemonController{}, iniErrror
	}
	return PokemonController{
		DataService: dataService,
	}, nil
}
