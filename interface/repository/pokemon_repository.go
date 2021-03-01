package repository

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/Topi99/academy-go-q12021/domain/model"
	"github.com/Topi99/academy-go-q12021/usecase/repository"
)

type pokemonRepository struct {
}

// PokemonRepository interface
type PokemonRepository interface {
	FindOne(id uint) (*model.Pokemon, error)
}

// NewPokemonRepository returns new PokemonRepository
func NewPokemonRepository() PokemonRepository {
	return &pokemonRepository{}
}

func (po *pokemonRepository) FindOne(id uint) (*model.Pokemon, error) {
	f, err := os.Open("../../infrastructure/datastore/pokemons.csv")

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		recordID, err := strconv.ParseUint(record[0], 10, 32)

		if err != nil {
			return nil, err
		}

		uintID := uint(recordID)

		if uintID == id {
			p := &model.Pokemon{
				ID:   uintID,
				Name: record[1],
				URL:  record[2],
			}
			return p, nil
		}
	}

	return nil, repository.ErrNotFound
}
