package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/grethelBello/academy-go-q12021/config"
	"github.com/grethelBello/academy-go-q12021/constant"
	"github.com/grethelBello/academy-go-q12021/data"
	"github.com/grethelBello/academy-go-q12021/model"
	"github.com/grethelBello/academy-go-q12021/model/errs"
)

const restDataType = "pokemon"

var restIsDigit = regexp.MustCompile(`\d+`)

// RestPokemonDataService is a service layer to work with the data (list, filter, etc.)
type RestPokemonDataService struct {
	Data       map[int]model.Pokemon
	keys       pokemonsIDSorter
	CsvSource  data.Source
	HttpSource data.Source
}

// Init initiliazes the data layer
func (pds *RestPokemonDataService) Init() error {
	data, err := pds.CsvSource.GetData()
	if err != nil {
		fmt.Printf("Error initiating pokemon service: %v\n", err)
		return err
	}

	if pds.Data == nil || len(pds.Data) == 0 {
		pds.Data = make(map[int]model.Pokemon)
	}

	for _, line := range data.CsvData {
		id := strings.Trim(line[0], " ")
		name := strings.Trim(line[1], " ")

		if !restIsDigit.MatchString(id) || len(name) == 0 {
			fmt.Println("Header is present")
			continue
		}

		convID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		pokemon := model.Pokemon{Id: convID, Name: name}
		pds.Data[convID] = pokemon
	}

	pds.setPokemonKeys()
	fmt.Printf("Pokemon Service initiated: %v\n", *pds)
	return nil
}

// Get returns a pokemon by ID
func (pds *RestPokemonDataService) Get(id int) model.Response {
	// Look for Pokemon in CSV Data
	pokemon, ok := pds.Data[id]
	if ok {
		pokemons := []model.Pokemon{pokemon}
		response := model.StaticResponse{Result: pokemons, Total: 1, Count: 1, Page: 1}
		return response
	}
	fmt.Printf("Pokemon %v not found in CSV source\n", id)
	notFoundError := errs.NotFoundError{Id: id, Datatype: restDataType}

	// Look for Pokemon in API
	if httpsource, ok := pds.HttpSource.(data.HttpSource); ok {
		pokemon, apiError := pds.getPokemonFromAPI(id, &httpsource)
		if apiError == nil {
			response := model.StaticResponse{Result: []model.Pokemon{pokemon}, Total: len(pds.Data), Count: 1, Page: 1}
			return response
		}
		fmt.Printf("Pokemon %v not found in API source\n", id)
		notFoundError.TechnicalError = apiError
	} else {
		fmt.Println("Error converting HttpSource")
		notFoundError.TechnicalError = errors.New("Error converting HttpSource")
	}

	return model.StaticResponse{Error: notFoundError}
}

// List returns all the pokemons by page
func (pds *RestPokemonDataService) List(count, page int) model.Response {
	if pds.Data == nil || len(pds.Data) == 0 {
		emptyError := errs.EmptyDataError(restDataType)
		return model.StaticResponse{Error: emptyError}
	}

	pokemons, page := pds.getPage(count, page)
	return model.StaticResponse{Result: pokemons, Total: len(pds.Data), Page: page, Count: count}
}

// Filter disabled for this delivery
func (pds *RestPokemonDataService) Filter(typeFilter model.TypeFilter, items, itemsPerWorker int) model.Response {
	return model.StaticResponse{}
}

// Sync is used to read the data from CSV to have consistent data
func (pds *RestPokemonDataService) Sync() error {
	return pds.Init()
}

func (pds *RestPokemonDataService) setPokemonKeys() {
	keys := make([]int, len(pds.Data))
	index := 0
	for key := range pds.Data {
		keys[index] = key
		index++
	}

	pds.keys = pokemonsIDSorter(keys)
	sort.Sort(pds.keys)
}

func (pds *RestPokemonDataService) getPage(count, page int) ([]model.Pokemon, int) {
	total := len(pds.keys)
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

	for pokeKeyIndex, pokemonKey := range pds.keys {
		if pokeKeyIndex >= startFrom && pokeKeyIndex < endAt {
			pokemon, ok := pds.Data[pokemonKey]
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

func (pds *RestPokemonDataService) getPokemonFromAPI(id int, httpSource *data.HttpSource) (model.Pokemon, error) {
	pokemonService, getPokeServError := config.GetEnvVar(constant.PokemonServiceVarName)
	if getPokeServError != nil {
		return model.Pokemon{}, getPokeServError
	}
	httpData := model.HttpData{
		Url:    fmt.Sprintf("%v/%v", pokemonService, id),
		Method: http.MethodGet,
	}

	httpSource.NewData(httpData)
	apiResponse, error := httpSource.GetData()
	if error != nil {
		return model.Pokemon{}, error
	}

	var pokemon model.Pokemon
	if unmarshallError := json.Unmarshal([]byte(apiResponse.HttpData), &pokemon); unmarshallError != nil {
		return model.Pokemon{}, unmarshallError
	}

	appendPokemon := model.Data{
		CsvData: [][]string{
			{
				fmt.Sprint(pokemon.Id),
				pokemon.Name,
			},
		}}
	defer pds.CsvSource.SetData(&appendPokemon)
	pds.Data[pokemon.Id] = pokemon
	pds.setPokemonKeys()
	return pokemon, nil
}

func NewRestPokemonDataService() (DataService, error) {
	csvPath, csvError := config.GetEnvVar(constant.PokemonSourceVarName)
	if csvError != nil {
		return &RestPokemonDataService{}, csvError
	}
	csvSource := data.CsvSource(csvPath)
	httpSource := data.HttpSource{
		Client: &http.Client{Timeout: time.Minute},
	}
	service := &RestPokemonDataService{
		CsvSource:  csvSource,
		HttpSource: httpSource,
	}

	return service, nil
}
