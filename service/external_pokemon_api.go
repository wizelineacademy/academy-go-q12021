package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wizelineacademy/academy-go-q12021/model"
)

// ExternalPokemonAPI dependencies from Pokemon service
type ExternalPokemonAPI struct {
	url string
}

// NewExternalPokemonAPI initializer method for create ExternalPokemonAPI
func NewExternalPokemonAPI() *ExternalPokemonAPI {
	return &ExternalPokemonAPI{
		url: "https://pokeapi.co/api/v2/pokemon",
	}
}

func (s *ExternalPokemonAPI) GetPokemonFromAPI(id int) (*model.PokemonAPI, error) {
	response, err := http.Get(fmt.Sprintf("%s/%d", s.url, id))
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusNotFound {
		return nil, errors.New("The Pokemon does not exist")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var pokemonAPI *model.PokemonAPI
	json.Unmarshal(responseData, &pokemonAPI)
	return pokemonAPI, nil
}
