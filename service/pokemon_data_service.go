package service

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"../config"
	"../data"
	"../model"
	"../model/errs"
)

const dataType = "pokemon"

var isHeaderLine = regexp.MustCompile(`\d+`)
var pokemonSource = config.PokemonSource

type PokemonDataService map[int]model.Pokemon

type PokemonsIdSorter []int

func (pds PokemonDataService) Init() {
	datasource := data.CsvSource(pokemonSource)
	results, err := datasource.Init()
	if err != nil {
		return
	}

	for _, line := range results {
		id := strings.Trim(line[0], " ")
		name := strings.Trim(line[1], " ")

		if isHeaderLine.MatchString(id) || len(name) == 0 {
			continue
		}

		convId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		pokemon := model.Pokemon{Id: convId, Name: name}
		pds[convId] = pokemon
	}

	if len(pds) == 0 {
		pds = nil
	}
}

func (pds PokemonDataService) Get(id int) (model.Pokemon, error) {
	if pds == nil {
		emptyError := errs.EmptyDataError(dataType)
		return model.Pokemon{}, emptyError
	}

	pokemon, ok := pds[id]
	if ok {
		return pokemon, nil
	}

	notFoundError := errs.NotFoundError{Id: id, Datatype: dataType}
	return model.Pokemon{}, notFoundError
}

func (pds PokemonDataService) List(count, page int) ([]model.Pokemon, error) {
	if pds == nil {
		emptyError := errs.EmptyDataError(dataType)
		return []model.Pokemon{}, emptyError
	}

	result := pds.getPage(count, page)
	return result, nil
}

func (pds PokemonDataService) getPokemonKeys() []int {
	keys := make([]int, 1)
	for key, _ := range pds {
		keys = append(keys, key)
	}

	return keys
}

func (pds PokemonDataService) getPage(count, page int) []model.Pokemon {
	pokemonIds := PokemonsIdSorter(pds.getPokemonKeys())
	sort.Sort(pokemonIds)

	total := len(pokemonIds)
	if count > total {
		count = total
	}

	startFrom := count*page - count
	endAt := count * page
	index := 0
	resultList := make([]model.Pokemon, count)

	for pokemonKey := range pokemonIds {
		if index >= startFrom && index < endAt {
			pokemon, ok := pds[pokemonKey]
			if ok {
				resultList = append(resultList, pokemon)
			}
		}
		index++
		if index == endAt {
			break
		}
	}

	return resultList
}

func (pis PokemonsIdSorter) Len() int { return len(pis) }

func (pis PokemonsIdSorter) Less(i, j int) bool { return pis[i] < pis[j] }

func (pis PokemonsIdSorter) Swap(i, j int) { pis[i], pis[j] = pis[j], pis[i] }
