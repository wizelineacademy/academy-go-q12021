package httpservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"pokeapi/model"
)

type httpService struct{}

const uri string = "https://pokeapi.co/api/v2/pokemon?limit=5&offset=300"

type NewHttpService interface {
	GetPokemons() ([]model.SinglePokeExternal, *model.Error)
}

func New() *httpService {
	return &httpService{}
}

func (h *httpService) GetPokemons() ([]model.SinglePokeExternal, *model.Error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong creating the GET request",
		}
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong sending your request",
		}
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong reading the response",
		}
	}
	var response model.PokemonExternal
	json.Unmarshal(bodyBytes, &response)

	newPokemons := response.Results
	return newPokemons, nil
}
