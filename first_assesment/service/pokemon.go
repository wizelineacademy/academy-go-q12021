package service

import (
	"log"

	"github.com/wizelineacademy/academy-go-q12021/model"
	"github.com/wizelineacademy/academy-go-q12021/repository"
)

// PokemonService dependencies from Pokemon service
type PokemonService struct {
	pokemonRepository *repository.PokemonRepository
}

// NewPokemonService initializer method for create PokemonService
func NewPokemonService() (*PokemonService, error) {
	pokemonRepository, err := repository.NewPokemonRepository()
	if err != nil {
		return nil, err
	}

	return &PokemonService{
		pokemonRepository: pokemonRepository,
	}, nil
}

// GetAll get all pokemons from repository
func (s *PokemonService) GetAll() ([]model.Pokemon, error) {
	log.Println("Enter to get all pokemons!!!")
	pokemons, err := s.pokemonRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

// GetByID get pokemon by his id
func (s *PokemonService) GetByID(id int) (*model.Pokemon, error) {
	log.Println("Enter to get pokemon by id!!!")
	pokemon, err := s.pokemonRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}
