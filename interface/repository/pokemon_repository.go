package repository

import (
	"encoding/csv"
	"os"

	"github.com/AlejandroSeguraWIZ/academy-go-q12021/domain/model"
)

type pokemonRepository struct {
	fileName string
}

type PokemonRepository interface {
	FetchAll() ([]model.Pokemon, error)
}

func NewPokemonRepository(fileName string) PokemonRepository {
	return &pokemonRepository{
		fileName: fileName,
	}
}

func (pr *pokemonRepository) FetchAll() ([]model.Pokemon, error) {
	f, err := os.Open(pr.fileName)
	if err != nil {
		return []model.Pokemon{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return []model.Pokemon{}, err
	}

	result := []model.Pokemon{}
	for _, line := range lines[1:] {
		pokemon := model.BuildPokemonFromStore(line)
		result = append(result, pokemon)
	}
	return result, nil
}
