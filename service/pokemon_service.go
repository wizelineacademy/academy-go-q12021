package service

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/wizelineacademy/academy-go/data"
	"github.com/wizelineacademy/academy-go/model"
	"github.com/wizelineacademy/academy-go/model/errs"
)

const dataType = "pokemon"

var isDigit = regexp.MustCompile(`\d+`)

// PokemonDataService is a service layer to work with the data (list, filter, etc.)
type PokemonDataService map[int]model.Pokemon

// Init initiliazes the data layer
func (pds PokemonDataService) Init(datasource data.Source) error {
	data, err := datasource.GetData()
	if err != nil {
		return err
	}

	for _, line := range data.CsvData {
		id := strings.Trim(line[0], " ")
		name := strings.Trim(line[1], " ")

		if !isDigit.MatchString(id) || len(name) == 0 {
			fmt.Println("Header is present")
			continue
		}

		convID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		pokemon := model.Pokemon{Id: convID, Name: name}
		pds[convID] = pokemon
	}

	return nil
}

// Get returns a pokemon by ID
func (pds PokemonDataService) Get(id int) model.Response {
	if pds == nil || len(pds) == 0 {
		emptyError := errs.EmptyDataError(dataType)
		return model.Response{Error: emptyError}
	}

	pokemon, ok := pds[id]
	if ok {
		pokemons := []model.Pokemon{pokemon}
		response := model.Response{Result: pokemons, Total: 1, Count: 1, Page: 1}
		return response
	}

	notFoundError := errs.NotFoundError{Id: id, Datatype: dataType}
	return model.Response{Error: notFoundError}
}

// List returns all the pokemons by page
func (pds PokemonDataService) List(count, page int) model.Response {
	if pds == nil || len(pds) == 0 {
		emptyError := errs.EmptyDataError(dataType)
		return model.Response{Error: emptyError}
	}

	pokemons, page := pds.getPage(count, page)
	return model.Response{Result: pokemons, Total: len(pds), Page: page, Count: count}
}

func (pds PokemonDataService) getPokemonKeys() []int {
	keys := make([]int, len(pds))
	index := 0
	for key := range pds {
		keys[index] = key
		index++
	}

	return keys
}

func (pds PokemonDataService) getPage(count, page int) ([]model.Pokemon, int) {
	pokemonIds := pokemonsIDSorter(pds.getPokemonKeys())
	sort.Sort(pokemonIds)

	total := len(pokemonIds)
	if count > total {
		count = total
	}

	if count*page > total {
		page = total / count
	}

	startFrom := count*page - count
	endAt := count * page
	index := 0
	resultList := make([]model.Pokemon, count)

	for pokeKeyIndex, pokemonKey := range pokemonIds {
		if pokeKeyIndex >= startFrom && pokeKeyIndex < endAt {
			pokemon, ok := pds[pokemonKey]
			if ok {
				resultList[index] = pokemon
				index++
			}
		}

		if pokeKeyIndex == endAt {
			break
		}
	}

	return resultList, page
}

type pokemonsIDSorter []int

func (pis pokemonsIDSorter) Len() int { return len(pis) }

func (pis pokemonsIDSorter) Less(i, j int) bool { return pis[i] < pis[j] }

func (pis pokemonsIDSorter) Swap(i, j int) { pis[i], pis[j] = pis[j], pis[i] }
