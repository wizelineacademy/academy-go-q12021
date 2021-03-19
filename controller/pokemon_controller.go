package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wizelineacademy/academy-go/service"
)

var dataService service.DataService

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

var getPokemons = func(w http.ResponseWriter, r *http.Request) {
	queryParams := getPokemonsQueryParamas(r)

	// Return just one pokemon by ID
	id, ok := queryParams["id"]
	if ok {
		responseGet := dataService.Get(id)
		if responseGet.Error != nil {
			http.Error(w, responseGet.Error.Error(), http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(responseGet)
		}
		return
	}

	count, _ := queryParams["count"]
	page, _ := queryParams["page"]
	responseList := dataService.List(count, page)
	if responseList.Error != nil {
		http.Error(w, responseList.Error.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(responseList)
	}
}

// GetPokemonRoutes returns a map with the handlers for pokemon endpoints
func GetPokemonRoutes(pokemonService service.DataService) map[string]func(http.ResponseWriter, *http.Request) {
	dataService = pokemonService
	pokemonRoutes := map[string]func(http.ResponseWriter, *http.Request){
		"/pokemons": getPokemons,
	}

	return pokemonRoutes
}
