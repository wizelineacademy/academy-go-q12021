package rest

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"golang-bootcamp-2020/domain/model"
	_errors "golang-bootcamp-2020/utils/error"
)

var (
	apiCharacters = "https://rickandmortyapi.com/api/character/"
)

type restResponse struct {
	Info    paginationInfo `json:"info"`
	Results []interface{}  `json:"results"`
}

type paginationInfo struct {
	Next string `json:"next"`
}

type rickAndMortyApi struct {
}

type RickAndMortyApiRepository interface {
	FetchData(maxPages int) ([]model.Character, _errors.RestError)
}

//Return new rest repository
func NewRickAndMortyApiRepository() RickAndMortyApiRepository {
	return &rickAndMortyApi{}
}

//Fetch data from rick and morty api
func (api *rickAndMortyApi) FetchData(maxPages int) ([]model.Character, _errors.RestError) {
	// fetching characters
	var characters []model.Character
	resp, err := processRequest(apiCharacters, maxPages)

	if err != nil {
		return nil, err
	}

	for i := range resp {
		var ch []model.Character
		json.Unmarshal(resp[i], &ch)
		characters = append(characters, ch...)
	}

	return characters, nil
}

func processRequest(url string, maxPages int) ([][]byte, _errors.RestError) {

	var response [][]byte
	endpoint := url
	count := 1

	if maxPages == 0 {
		maxPages = viper.GetInt("rest.maxPagesByDefault")
	}

	for {
		client := resty.New()

		resp, err := client.R().Get(endpoint)
		if err != nil || !resp.IsSuccess() {
			return nil, _errors.NewInternalServerError("error in rest response")
		}

		restR := &restResponse{}
		json.Unmarshal([]byte(resp.String()), restR)
		jsonResult, _ := json.Marshal(restR.Results)

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
