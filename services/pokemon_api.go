package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/cesararredondow/academy-go-q12021/models"
)

//GetPokemons get an amount of pokemons that you request from pokemon API
func (s *Service) GetPokemonsFromAPI(quantity string) ([]*models.Pokemon_api, error) {
	pokemons := []*models.Pokemon_api{}
	url := "/pokemon/?limit="
	url += quantity

	response, err := http.Get(s.pokemonAPI + url)
	if err != nil {
		return nil, errors.New("Error")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error")
	}

	jsonData := models.Response{}
	json.Unmarshal(responseData, &jsonData)

	for i, pokemon := range jsonData.Results {
		p := new(models.Pokemon_api)
		p.ID = i + 1
		p.Name = pokemon.Name
		p.URL = pokemon.URL
		pokemons = append(pokemons, p)
	}

	if errCSV := s.writeJSONInCSV(jsonData); errCSV != nil {
		return nil, errCSV
	}

	return pokemons, nil
}

//GetPokemon get the pokemons that you request from pokemon API
func (s *Service) GetPokemonFromAPI(pokemonID string) (*models.PokemonResponse, error) {
	response, err := http.Get(s.pokemonAPI + "/pokemon/" + pokemonID)

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error")
	}

	jsonData := models.PokemonResponse{}
	json.Unmarshal(responseData, &jsonData)

	return &jsonData, nil
}

func (s *Service) writeJSONInCSV(jsonData models.Response) error {
	_, err := os.Stat(s.pathFile)
	if os.IsNotExist(err) {
		s.writer.Write([]string{"ID", "NAME"})
	}
	for i, p := range jsonData.Results {
		var row []string
		row = append(row, strconv.Itoa(i+1))
		row = append(row, p.Name)
		s.writer.Write(row)
	}

	s.writer.Flush()
	return nil
}
