package business

import (
	"log"

	"github.com/wizelineacademy/academy-go-q12021/model"
	"github.com/wizelineacademy/academy-go-q12021/repository"
	"github.com/wizelineacademy/academy-go-q12021/service"
)

type IPokemonBusiness interface {
	GetAll() ([]model.Pokemon, error)
	GetByID(id int) (*model.Pokemon, error)
	StoreByID(id int) (*model.Pokemon, error)
}

// PokemonService dependencies from Pokemon service
type PokemonBusiness struct {
	pokemonRepository repository.IPokemonRepository
	serviceAPI        service.IExternalPokemonAPI
}

// NewPokemonService initializer method for create PokemonService
func NewPokemonBusiness(repository repository.IPokemonRepository, service service.IExternalPokemonAPI) (IPokemonBusiness, error) {
	return &PokemonBusiness{
		pokemonRepository: repository,
		serviceAPI:        service,
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
	pokemonAPI, err := s.serviceAPI.GetPokemonFromAPI(id)
	if err != nil {
		return nil, err
	}
	pokemon, err := s.pokemonRepository.StoreToCSV(*pokemonAPI)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}
