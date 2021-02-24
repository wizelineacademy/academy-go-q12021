package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokeapi/model"
)

const uri string = "https://pokeapi.co/api/v2/pokemon?limit=5&offset=300"

type HttpService struct{}

type NewHttpService interface {
	GetPokemonsFromExternalAPI() *model.Error
}

func NewHttp() *HttpService {
	return &HttpService{}
}

func (h *HttpService) GetPokemonsFromExternalAPI() *model.Error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Something happened", err.Error())

		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return &err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())

		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return &err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())

		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return &err
	}

	var response model.PokemonExternal
	json.Unmarshal(bodyBytes, &response)

	fmt.Printf("API Response as struct %+v\n", response)
	addLineCsv(&response.Results)

	return nil
}
