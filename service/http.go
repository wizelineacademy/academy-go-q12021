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
	GetPokemonsFromExternalAPI()
}

func NewHttp() *HttpService {
	return &HttpService{}
}

func (h *HttpService) GetPokemonsFromExternalAPI() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Something happened", err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var response model.PokemonExternal
	json.Unmarshal(bodyBytes, &response)

	fmt.Printf("API Response as struct %+v\n", response)
	addLineCsv(&response.Results)

}
