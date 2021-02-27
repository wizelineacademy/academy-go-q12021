package service

import (
	"first/repository"
	"fmt"
)

type PokemonService struct {
	pokemonRepository *repository.PokemonRepository
}

func NewPokemonService() (*PokemonService, error) {
	pokemonRepository, err := repository.NewPokemonRepository()
	if err != nil {
		panic(err)
	}

	return &PokemonService{
		pokemonRepository: pokemonRepository,
	}, nil
}

func (s *PokemonService) Saludo() {
	fmt.Println("Hello world")
}
