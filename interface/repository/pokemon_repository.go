package repository

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/Topi99/academy-go-q12021/domain/model"
	"github.com/Topi99/academy-go-q12021/usecase/repository"
)

const file_ext = ".csv"

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
	var p *model.Pokemon

	f, err := os.Open(p.TableName() + file_ext)

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
			p.ID = uintID
			p.Name = record[1]
			p.URL = record[2]
			return p, nil
		}
	}

	return nil, repository.ErrNotFound
}
