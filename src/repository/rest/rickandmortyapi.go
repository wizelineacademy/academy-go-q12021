package rest

import (
	"encoding/json"
	"fmt"

	"golang-bootcamp-2020/domain/model"
	_errors "golang-bootcamp-2020/utils/error"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var (
	apiCharacters string
)

type restResponse struct {
	Info    paginationInfo `json:"info"`
	Results []interface{}  `json:"results"`
}

type paginationInfo struct {
	Next string `json:"next"`
}

type rickAndMortyAPI struct {
	restClient *resty.Client
}

//RickAndMortyAPIRepository - rickandmorty repository methods
type RickAndMortyAPIRepository interface {
	FetchData(maxPages int) ([]model.Character, _errors.RestError)
}

//Init - init repo data
func Init() {
	apiCharacters = viper.GetString("rest.host") + "character/"
}

//NewRickAndMortyAPIRepository - Return new rest repository
func NewRickAndMortyAPIRepository(restClient *resty.Client) RickAndMortyAPIRepository {
	Init()
	return &rickAndMortyAPI{restClient}
}

//FetchData - Fetch data from rick and morty API
func (api *rickAndMortyAPI) FetchData(maxPages int) ([]model.Character, _errors.RestError) {
	// fetching characters
	var characters []model.Character
	resp, err := api.processRequest(apiCharacters, maxPages)

	if err != nil {
		return nil, err
	}

	for i := range resp {
		var ch []model.Character
		if err := json.Unmarshal(resp[i], &ch); err != nil {
			return nil, _errors.NewInternalServerError("error when trying to unmarshal results")
		}
		characters = append(characters, ch...)
	}

	return characters, nil
}

func (api *rickAndMortyAPI) processRequest(endpoint string, maxPages int) ([][]byte, _errors.RestError) {

	var response [][]byte
	count := 1

	if maxPages == 0 {
		maxPages = viper.GetInt("rest.maxPagesByDefault")
	}

	for {
		fmt.Println(endpoint)
		resp, err := api.restClient.R().Get(endpoint)
		if err != nil || !resp.IsSuccess() {
			return nil, _errors.NewInternalServerError("error in rest response")
		}

		restR := &restResponse{}
		if err := json.Unmarshal([]byte(resp.String()), restR); err != nil {
			return nil, _errors.NewInternalServerError("error when trying to unmarshal results")
		}
		jsonResult, err := json.Marshal(restR.Results)
		if err != nil {
			return nil, _errors.NewInternalServerError("error when trying to marshal results")
		}

		response = append(response, jsonResult)

		if restR.Info.Next == "" || count >= maxPages {
			break
		} else {
			endpoint = restR.Info.Next
			count++
		}
	}

	return response, nil
}
