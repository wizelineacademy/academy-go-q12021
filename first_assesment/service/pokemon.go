package service

import (
	"first/model"
	"first/repository"
	"fmt"
)

type PokemonService struct {
	pokemonRepository *repository.PokemonRepository
}

func NewPokemonService() (*PokemonService, error) {
	pokemonRepository, err := repository.NewPokemonRepository()
	if err != nil {
		return nil, err
	}

	return &PokemonService{
		pokemonRepository: pokemonRepository,
	}, nil
}

func (s *PokemonService) Saludo() {
	fmt.Println("Hello world")
}

func (s *PokemonService) GetAll() ([]*model.Pokemon, error) {
	fmt.Println("Enter to get all pokemons!!!")
	pokemons, err := s.pokemonRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (s *PokemonService) GetById(id int) (*model.Pokemon, error) {
	fmt.Println("Enter to get pokemon by id!!!")
	pokemon, err := s.pokemonRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}
