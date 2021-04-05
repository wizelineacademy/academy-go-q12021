package services

import (
	"encoding/csv"
	"errors"
	"strconv"

	"github.com/cesararredondow/academy-go-q12021/models"
)

//GetPokemons read all the information in the csv y return an array of them
func (s *Service) GetPokemons() ([]*models.Pokemon, error) {
	pokemons := []*models.Pokemon{}
	r := csv.NewReader(s.reader)
	records, err := r.ReadAll()
	if err != nil {
		return nil, errors.New("error Reading the file")
	}

	if records != nil {
		records := records[1:]
		for _, rec := range records {
			p := new(models.Pokemon)
			val, err := strconv.Atoi(rec[0])
			if err != nil {
				return nil, errors.New("error Converting the id to int from csv file")
			}
			p.ID = val
			p.Name = rec[1]
			pokemons = append(pokemons, p)
		}
	}
	s.reader.Seek(0, 0)
	return pokemons, nil
}

//GetPokemon read all the information in the csv y return an pokemon of them
func (s *Service) GetPokemon(pokemonID string) (*models.Pokemon, error) {
	pokemon := &models.Pokemon{}
	r := csv.NewReader(s.reader)
	records, err := r.ReadAll()
	if err != nil {
		return nil, errors.New("error Reading the file")
	}

	found := false

	if records != nil {
		records := records[1:]
		for _, rec := range records {
			if rec[0] == pokemonID {
				val, err := strconv.Atoi(rec[0])
				if err != nil {
					return nil, errors.New("error Converting the id to int from csv file")
				}
				pokemon.ID = val
				pokemon.Name = rec[1]
				found = true
				break
			}
		}
	}
	if found {
		return pokemon, nil
	} else {
		return nil, nil
	}
}
