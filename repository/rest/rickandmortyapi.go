package rest

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"golang-bootcamp-2020/domain/model"
)

const (
	apiCharacters = "https://rickandmortyapi.com/api/character/"
	//API_LOCATIONS  = "https://rickandmortyapi.com/api/location/"
	//API_EPISODES   = "https://rickandmortyapi.com/api/episode/"
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
	GetCharacters() (interface{}, error)
}

func NewRickAndMortyApiRepository() RickAndMortyApiRepository {
	return &rickAndMortyApi{}
}

func (api *rickAndMortyApi) GetCharacters() (interface{}, error) {
	resp, err := processRequest(apiCharacters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func processRequest(url string) (interface{}, error) {
	client := resty.New()

	resp, err := client.R().Get(url)
	if err != nil || !resp.IsSuccess() {
		return nil, errors.New("error in rest response")
	}

	parsed := parseResponse(resp.String())
	return parsed, nil
}

func parseResponse(response string) interface{} {

	res := &restResponse{}
	var characters []model.Character
	json.Unmarshal([]byte(response), res)

	jsonResults, _ := json.Marshal(res.Results)
	json.Unmarshal(jsonResults, &characters)
	return characters
}
