package service

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/grethelBello/academy-go-q12021/config"
	"github.com/grethelBello/academy-go-q12021/constant"
	"github.com/grethelBello/academy-go-q12021/data"
	"github.com/grethelBello/academy-go-q12021/model"
	"github.com/grethelBello/academy-go-q12021/model/errs"
)

const staticDataType = "pokemon"

var staticIsDigit = regexp.MustCompile(`\d+`)

// StaticPokemonDataService is a service layer to work with the data (list, filter, etc.)
type StaticPokemonDataService struct {
	data      map[int]model.Pokemon
	csvSource data.Source
}

// Init initiliazes the data layer
func (pds *StaticPokemonDataService) Init() error {
	data, err := pds.csvSource.GetData()
	if err != nil {
		return err
	}

	if pds.data == nil || len(pds.data) == 0 {
		pds.data = make(map[int]model.Pokemon)
	}

	for _, line := range data.CsvData {
		id := strings.Trim(line[0], " ")
		name := strings.Trim(line[1], " ")

		if !staticIsDigit.MatchString(id) || len(name) == 0 {
			fmt.Println("Header is present")
			continue
		}

		convID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		pokemon := model.Pokemon{Id: convID, Name: name}
		pds.data[convID] = pokemon
	}

	return nil
}

// Get returns a pokemon by ID
func (pds *StaticPokemonDataService) Get(id int) model.Response {
	if pds.data == nil || len(pds.data) == 0 {
		emptyError := errs.EmptyDataError(staticDataType)
		return model.StaticResponse{Error: emptyError}
	}

	pokemon, ok := pds.data[id]
	if ok {
		pokemons := []model.Pokemon{pokemon}
		// TODO return total of registers
		response := model.StaticResponse{Result: pokemons, Total: len(pds.data), Count: 1, Page: 1}
		return response
	}

	notFoundError := errs.NotFoundError{Id: id, Datatype: staticDataType}
	return model.StaticResponse{Error: notFoundError}
}

// List returns all the pokemons by page
func (pds *StaticPokemonDataService) List(count, page int) model.Response {
	if pds.data == nil || len(pds.data) == 0 {
		emptyError := errs.EmptyDataError(staticDataType)
		return model.StaticResponse{Error: emptyError}
	}

	pokemons, page := pds.getPage(count, page)
	return model.StaticResponse{Result: pokemons, Total: len(pds.data), Page: page, Count: count}
}

// Filter disabled for this delivery
func (pds *StaticPokemonDataService) Filter(typeFilter model.TypeFilter, items, itemsPerWorker int) model.Response {
	return model.StaticResponse{}
}

// Sync is used to read the data from CSV to have consistent data
func (pds *StaticPokemonDataService) Sync() error {
	return pds.Init()
}

func (pds *StaticPokemonDataService) getPokemonKeys() []int {
	keys := make([]int, len(pds.data))
	index := 0
	for key := range pds.data {
		keys[index] = key
		index++
	}

	return keys
}

func (pds *StaticPokemonDataService) getPage(count, page int) ([]model.Pokemon, int) {
	pokemonIds := staticPokemonsIDSorter(pds.getPokemonKeys())
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
			pokemon, ok := pds.data[pokemonKey]
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

type staticPokemonsIDSorter []int

func (pis staticPokemonsIDSorter) Len() int { return len(pis) }

func (pis staticPokemonsIDSorter) Less(i, j int) bool { return pis[i] < pis[j] }

func (pis staticPokemonsIDSorter) Swap(i, j int) { pis[i], pis[j] = pis[j], pis[i] }

func NewStaticPokemonDataService() (DataService, error) {
	csvPath, csvError := config.GetEnvVar(constant.PokemonSourceVarName)
	if csvError != nil {
		return &StaticPokemonDataService{}, csvError
	}
	csvSource := data.CsvSource(csvPath)
	service := StaticPokemonDataService{
		data:      make(map[int]model.Pokemon),
		csvSource: csvSource,
	}
	service.Init()

	return &service, nil
}
