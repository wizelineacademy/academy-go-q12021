package service

import (
	"fmt"
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

func (s *ExternalPokemonAPI) getPokemonFromAPI(id int) (*model.Pokemon, error) {
	response, err := http.Get(fmt.Sprintf("%s/%d", s.url, id))
	if err != nil {
		return nil, err
	}
}
