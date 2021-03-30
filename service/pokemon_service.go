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

const dataType = "pokemon"

var isDigit = regexp.MustCompile(`\d+`)

// PokemonDataService is a service layer to work with the data (list, filter, etc.)
type PokemonDataService struct {
	Data       map[int]model.Pokemon
	keys       pokemonsIDSorter
	CsvSource  data.Source
	HttpSource data.Source
}

// Init initiliazes the data layer
func (pds *PokemonDataService) Init() error {
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
		pds.Data[convID] = pokemon
	}

	pds.setPokemonKeys()
	fmt.Printf("Pokemon Service initiated: %v\n", *pds)
	return nil
}

// Get returns a pokemon by ID
func (pds *PokemonDataService) Get(id int) model.Response {
	// Look for Pokemon in CSV Data
	pokemon, ok := pds.Data[id]
	if ok {
		pokemons := []model.Pokemon{pokemon}
		response := model.Response{Result: pokemons, Total: 1, Count: 1, Page: 1}
		return response
	}
	fmt.Printf("Pokemon %v not found in CSV source\n", id)
	notFoundError := errs.NotFoundError{Id: id, Datatype: dataType}

	// Look for Pokemon in API
	if httpsource, ok := pds.HttpSource.(data.HttpSource); ok {
		pokemon, apiError := pds.getPokemonFromAPI(id, &httpsource)
		if apiError == nil {
			response := model.Response{Result: []model.Pokemon{pokemon}, Total: len(pds.Data), Count: 1, Page: 1}
			return response
		}
		fmt.Printf("Pokemon %v not found in API source\n", id)
		notFoundError.TechnicalError = apiError
	} else {
		fmt.Println("Error converting HttpSource")
		notFoundError.TechnicalError = errors.New("Error converting HttpSource")
	}

	return model.Response{Error: notFoundError}
}

// List returns all the pokemons by page
func (pds *PokemonDataService) List(count, page int) model.Response {
	if pds.Data == nil || len(pds.Data) == 0 {
		emptyError := errs.EmptyDataError(dataType)
		return model.Response{Error: emptyError}
	}

	pokemons, page := pds.getPage(count, page)
	return model.Response{Result: pokemons, Total: len(pds.Data), Page: page, Count: count}
}

func (pds *PokemonDataService) setPokemonKeys() {
	keys := make([]int, len(pds.Data))
	index := 0
	for key := range pds.Data {
		keys[index] = key
		index++
	}

	pds.keys = pokemonsIDSorter(keys)
	sort.Sort(pds.keys)
}

func (pds *PokemonDataService) getPage(count, page int) ([]model.Pokemon, int) {
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

func (pds *PokemonDataService) getPokemonFromAPI(id int, httpSource *data.HttpSource) (model.Pokemon, error) {
	domainApi, envError := config.GetEnvVar(constant.PokemonServiceVarName)
	if envError != nil {
		return model.Pokemon{}, envError
	}
	httpData := model.HttpData{
		Url:    fmt.Sprintf("%v/%v", domainApi, id),
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

type pokemonsIDSorter []int

func (pis pokemonsIDSorter) Len() int { return len(pis) }

func (pis pokemonsIDSorter) Less(i, j int) bool { return pis[i] < pis[j] }

func (pis pokemonsIDSorter) Swap(i, j int) { pis[i], pis[j] = pis[j], pis[i] }

func NewPokemonDataService() (*PokemonDataService, error) {
	csvPath, csvError := config.GetEnvVar(constant.PokemonSourceVarName)
	if csvError != nil {
		return &PokemonDataService{}, csvError
	}
	csvSource := data.CsvSource(csvPath)
	httpSource := data.HttpSource{
		Client: &http.Client{Timeout: time.Minute},
	}
	service := &PokemonDataService{
		CsvSource:  csvSource,
		HttpSource: httpSource,
	}

	return service, nil
}
