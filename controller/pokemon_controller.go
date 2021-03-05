package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../model"
	"../service"
)

var dataService = service.PokemonDataService(make(map[int]model.Pokemon))

func getPokemonByIdQueryParamas(r *http.Request) map[string]int {
	queryParams := make(map[string]int, 3)

	queryId, ok := r.URL.Query()["id"]
	if ok && len(queryId[0]) > 0 {
		id, err := strconv.Atoi(queryId[0])
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

var getPokemonById = func(w http.ResponseWriter, r *http.Request) {
	queryParams := getPokemonByIdQueryParamas(r)

	id, ok := queryParams["id"]
	if ok {
		pokemon, err := dataService.Get(id)
		if err != nil {
			response := model.ResponsePokemon{Error: fmt.Sprintf("%v", err)}
			json.NewEncoder(w).Encode(response)
			return
		}

		result := make([]model.Pokemon, 1)
		result = append(result, pokemon)
		response := model.ResponsePokemon{Result: result, Total: len(result), Page: 1}
		json.NewEncoder(w).Encode(response)
		return
	}

	count, _ := queryParams["count"]
	page, _ := queryParams["page"]
	pokemons, err := dataService.List(count, page)
	if err != nil {
		response := model.ResponsePokemon{Error: fmt.Sprintf("%v", err)}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.ResponsePokemon{Result: pokemons, Total: len(dataService), Page: page}
	json.NewEncoder(w).Encode(response)
}

func GetPokemonRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	dataService.Init()
	pokemonRoutes := map[string]func(http.ResponseWriter, *http.Request){
		"/pokemons": getPokemonById,
	}

	return pokemonRoutes
}
