package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokeapi/models"
)

const uri string = "https://pokeapi.co/api/v2/pokemon?limit=5&offset=300"

type Http struct{}

type IHttp interface {
	GetPokemonsFromExternalAPI()
}

func NewHttpService() *Http {
	return &Http{}
}

func (h *Http) GetPokemonsFromExternalAPI() {
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

	var response models.PokemonExternal
	json.Unmarshal(bodyBytes, &response)

	fmt.Printf("API Response as struct %+v\n", response)
	addLineCsv(&response.Results)

}
