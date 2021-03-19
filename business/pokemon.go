package business

import (
	"log"

	"github.com/wizelineacademy/academy-go-q12021/model"
	"github.com/wizelineacademy/academy-go-q12021/repository"
	"github.com/wizelineacademy/academy-go-q12021/service"
)

// PokemonService dependencies from Pokemon service
type PokemonBusiness struct {
	pokemonRepository *repository.PokemonRepository
}

// NewPokemonService initializer method for create PokemonService
func NewPokemonBusiness() (*PokemonBusiness, error) {
	pokemonRepository, err := repository.NewPokemonRepository()
	if err != nil {
		return nil, err
	}

	return &PokemonBusiness{
		pokemonRepository: pokemonRepository,
	}, nil
}

// GetAll get all pokemons from repository
func (s *PokemonBusiness) GetAll() ([]model.Pokemon, error) {
	log.Println("Enter to get all pokemons!!!")
	pokemons, err := s.pokemonRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

// GetByID get pokemon by his id
func (s *PokemonBusiness) GetByID(id int) (*model.Pokemon, error) {
	log.Println("Enter to get pokemon by id!!!")
	pokemon, err := s.pokemonRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// StoreByID get pokemon by his id
func (s *PokemonBusiness) StoreByID(id int) (*model.Pokemon, error) {
	log.Println("Enter to search and store pokemon by id!!!")

	pokemonService := service.NewExternalPokemonAPI()

	pokemonAPI, err := pokemonService.GetPokemonFromAPI(id)
	if err != nil {
		return nil, err
	}
	pokemon, err := s.pokemonRepository.StoreToCSV(*pokemonAPI)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}
