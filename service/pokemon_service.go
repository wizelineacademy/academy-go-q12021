package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
var isOdd = func(value int) bool {
	return value%2 > 0
}
var isEven = func(value int) bool {
	return value%2 == 0
}
var lookForPokemons = func(data jobInfo, isOddFlag bool) {
	var evaluation func(value int) bool = isEven
	if isOddFlag {
		evaluation = isOdd
	}

	foundKeys := 0
	for _, key := range data.keys {
		if foundKeys < data.items && evaluation(key) {
			log.Printf("%v found %v\n", data.name, key)
			foundKeys++
			*data.results <- key
		} else if foundKeys >= data.items {
			break
		}
	}

	*data.shutdownChannel <- data.name
	return
}

type filterData struct {
	totalItems     int
	itemsPerWorker int
	numJobs        int
	segmentSize    int
	isOdd          bool
}

func (fd *filterData) getSegments(data map[int]model.Pokemon) [][]int {
	total := len(data)
	if fd.totalItems > total {
		fd.totalItems = total
	}

	if fd.itemsPerWorker > fd.totalItems {
		fd.itemsPerWorker = fd.totalItems
	}

	fd.numJobs = fd.totalItems / fd.itemsPerWorker
	fd.segmentSize = total / fd.numJobs
	if cover := fd.numJobs * fd.segmentSize; cover < total {
		fd.numJobs++
	}

	segments := make([][]int, fd.numJobs)
	jobIndex := 0
	for key := range data {
		jobData := len(segments[jobIndex])
		if jobData == fd.segmentSize {
			jobIndex++
			jobData = 0
		}

		segments[jobIndex][jobData] = key
	}

	return segments
}

type jobInfo struct {
	name            string
	items           int
	results         *chan int
	shutdownChannel *chan string
	keys            []int
}

// PokemonDataService is a service layer to work with the data (list, filter, etc.)
type PokemonDataService struct {
	Data       map[int]model.Pokemon
	CsvSource  data.Source
	HttpSource data.Source
}

// Init initiliazes the data layer
func (pds *PokemonDataService) Init() error {
	data, err := pds.CsvSource.GetData()
	if err != nil {
		log.Printf("Error initiating pokemon service: %v\n", err)
		return err
	}

	if pds.Data == nil || len(pds.Data) == 0 {
		pds.Data = make(map[int]model.Pokemon)
	}

	for _, line := range data.CsvData {
		id := strings.Trim(line[0], " ")
		name := strings.Trim(line[1], " ")

		if !isDigit.MatchString(id) || len(name) == 0 {
			log.Println("Header is present")
			continue
		}

		convID, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error parsing %v: '%v'\n", id, err)
			continue
		}

		pokemon := model.Pokemon{Id: convID, Name: name}
		pds.Data[convID] = pokemon
	}

	log.Printf("Pokemon Service initiated: %v\n", *pds)
	return nil
}

// Get returns a pokemon by ID
func (pds *PokemonDataService) Get(id int) model.Response {
	total := len(pds.Data)
	// Look for Pokemon in CSV Data
	pokemon, ok := pds.Data[id]
	if ok {
		pokemons := []model.Pokemon{pokemon}
		response := model.Response{Result: pokemons, Total: total, Items: 1}
		return response
	}
	log.Printf("Pokemon %v not found in CSV source\n", id)
	notFoundError := errs.NotFoundError{Id: id, Datatype: dataType}

	// Look for Pokemon in API
	if httpsource, ok := pds.HttpSource.(data.HttpSource); ok {
		pokemon, apiError := pds.getPokemonFromAPI(id, &httpsource)
		if apiError == nil {
			response := model.Response{Result: []model.Pokemon{pokemon}, Total: total, Items: 1}
			return response
		}
		log.Printf("Pokemon %v not found in API source\n", id)
		notFoundError.TechnicalError = apiError
	} else {
		log.Println("Error converting HttpSource")
		notFoundError.TechnicalError = errors.New("Error converting HttpSource")
	}

	return model.Response{Error: notFoundError}
}

// List returns all the pokemons by page
func (pds *PokemonDataService) List(typeFilter model.TypeFilter, items, itemsPerWorker int) model.Response {
	if pds.Data == nil || len(pds.Data) == 0 {
		emptyError := errs.EmptyDataError(dataType)
		return model.Response{Error: emptyError}
	}

	filterData := &filterData{
		isOdd:          typeFilter.IsOdd(),
		totalItems:     items,
		itemsPerWorker: itemsPerWorker,
	}
	keys := pds.filterPokemons(filterData)
	sorter := pokemonsIDSorter(keys)
	sort.Sort(sorter)

	pokemons := make([]model.Pokemon, len(keys))
	for index, key := range keys {
		pokemons[index] = pds.Data[key]
	}
	return model.Response{Result: pokemons, Total: len(pds.Data), Items: len(pokemons)}
}

func (pds *PokemonDataService) filterPokemons(data *filterData) []int {
	segments := data.getSegments(pds.Data)
	shutdown := make(chan string, data.numJobs)
	results := make(chan int, data.numJobs)

	for index := 0; index < data.numJobs; index++ {
		jobInfo := jobInfo{
			name:            fmt.Sprintf("Job %v", index+1),
			items:           data.itemsPerWorker,
			keys:            segments[index],
			results:         &results,
			shutdownChannel: &shutdown,
		}
		go lookForPokemons(jobInfo, data.isOdd)
	}

	ids := make([]int, data.totalItems)
	finishedJobs := 0
	for finishedJobs < data.numJobs {
		select {
		case <-shutdown:
			finishedJobs++
		case key := <-results:
			currentKey := len(ids)
			ids[currentKey] = key
		}
	}
	close(results)
	close(shutdown)

	return ids
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
