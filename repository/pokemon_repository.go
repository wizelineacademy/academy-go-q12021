package repository

import (
	"errors"
	"os"

	"github.com/wizelineacademy/academy-go-q12021/model"

	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

// PokemonRepository structure for repository, contains the csv file's name
type PokemonRepository struct {
	file string
}

// NewPokemonRepository method for create a Repository instance
func NewPokemonRepository() (*PokemonRepository, error) {
	pokemonFile := viper.Get("CSVFile").(string)
	return &PokemonRepository{
		file: pokemonFile,
	}, nil

}

// GetAll get all pokemons from csv file
func (p *PokemonRepository) GetAll() ([]model.Pokemon, error) {
	pokemonFile, err := p.openFile()
	if err != nil {
		return nil, err
	}
	pokemons := []model.Pokemon{}

	if err := gocsv.UnmarshalFile(pokemonFile, &pokemons); err != nil {
		return nil, errors.New("There was a problem parsing the csv file")
	}
	defer pokemonFile.Close()
	return pokemons, nil
}

// openFile open the csv file
func (p *PokemonRepository) openFile() (*os.File, error) {
	filePokemon, err := os.OpenFile(p.file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.New("There was a problem opening the csv file")
	}
	return filePokemon, nil
}

// GetByID get pokemon from csv by id
func (p *PokemonRepository) GetByID(id int) (*model.Pokemon, error) {
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return &pokemon, nil
		}
	}
	return nil, errors.New("the pokemon does not exist")
}

func (p PokemonRepository) StoreByID(id int) (*model.Pokemon, error) {
	/*	pokemonMap, err := p.getCSVDataInMemory()
		if err != nil {
			return nil, err
		}
	*/

	return nil, nil
}

func (p PokemonRepository) getCSVDataInMemory() (map[int]model.Pokemon, error) {
	pokemonMap := make(map[int]model.Pokemon)
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		pokemonMap[pokemon.Id] = pokemon
	}

	return pokemonMap, nil
}
