package rest

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"golang-bootcamp-2020/domain/model"
)

const (
	apiCharacters = "https://rickandmortyapi.com/api/character/"
	maxPages      = 1
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
	GetCharacters() ([]model.Character, error)
	FetchData() ([]model.Character, error)
}

func NewRickAndMortyApiRepository() RickAndMortyApiRepository {
	return &rickAndMortyApi{}
}

func (api *rickAndMortyApi) FetchData() ([]model.Character, error) {
	var err error
	// fetching characters
	var characters []model.Character
	characters, err = api.GetCharacters()
	if err != nil {
		return nil, errors.New("error fetching characters")
	}

	return characters, nil
}

func (api *rickAndMortyApi) GetCharacters() ([]model.Character, error) {
	var characters []model.Character

	resp, err := processRequest(apiCharacters)

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

func processRequest(url string) ([][]byte, error) {

	var response [][]byte
	endpoint := url
	count := 1

	for {
		client := resty.New()

		resp, err := client.R().Get(endpoint)
		if err != nil || !resp.IsSuccess() {
			return nil, errors.New("error in rest response")
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
