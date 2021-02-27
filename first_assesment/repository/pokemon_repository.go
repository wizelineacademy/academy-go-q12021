package repository

import (
	"os"

	"first/core"

	"github.com/gocarina/gocsv"
)

type Pokemon struct {
	Id            int    `csv:"ID"`
	Name          string `csv:"English"`
	JapaneseName  string `csv:"Japanese"`
	PrimaryType   string `csv:"Primary"`
	SecondaryType string `csv:"Secondary"`
	EvolvesTo     string `csv:"Evolves into"`
	Information   string `csv:"Notes"`
}

type PokemonRepository struct {
	file string
}

func NewPokemonRepository() (*PokemonRepository, error) {

	//defer filePokemons.Close()
	return &PokemonRepository{
		file: "csv/input_file.csv",
	}, nil

}

func (p *PokemonRepository) GetAll() ([]*Pokemon, error) {
	pokemonFile, err := p.openFile()
	if err != nil {
		return nil, err
	}
	pokemons := []*Pokemon{}

	if err := gocsv.UnmarshalFile(pokemonFile, &pokemons); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (p *PokemonRepository) openFile() (*os.File, error) {
	filePokemon, err := os.OpenFile(p.file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return filePokemon, nil
}

func (p *PokemonRepository) GetById(id int) (*Pokemon, error) {
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return pokemon, nil
		}
	}
	return nil, core.NewError("The pokemon does not exist!")
}
