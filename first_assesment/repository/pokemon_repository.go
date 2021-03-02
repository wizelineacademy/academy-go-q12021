package repository

import (
	"errors"
	"os"

	"first/model"

	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

type PokemonRepository struct {
	file string
}

func NewPokemonRepository() (*PokemonRepository, error) {
	pokemonFile := viper.Get("CSVFile").(string)
	return &PokemonRepository{
		file: pokemonFile,
	}, nil

}

func (p *PokemonRepository) GetAll() ([]*model.Pokemon, error) {
	pokemonFile, err := p.openFile()
	if err != nil {
		return nil, err
	}
	pokemons := []*model.Pokemon{}

	if err := gocsv.UnmarshalFile(pokemonFile, &pokemons); err != nil {
		return nil, errors.New("There was a problem parsing the csv file")
	}
	defer pokemonFile.Close()
	return pokemons, nil
}

func (p *PokemonRepository) openFile() (*os.File, error) {
	filePokemon, err := os.OpenFile(p.file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.New("There was a problem opening the csv file")
	}
	return filePokemon, nil
}

func (p *PokemonRepository) GetById(id int) (*model.Pokemon, error) {
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return pokemon, nil
		}
	}
	return nil, errors.New("The pokemon does not exist!")
}
