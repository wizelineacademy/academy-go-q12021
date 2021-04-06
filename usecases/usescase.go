package usecases

import "github.com/cesararredondow/academy-go-q12021/models"

//UseCase struct
type UseCase struct {
	service Service
}

//Service interface is contains the service
type Service interface {
	GetPokemons() ([]*models.Pokemon, error)
	GetPokemon(string) (*models.Pokemon, error)

	GetPokemonsFromAPI(string) ([]*models.Pokemon_api, error)
	GetPokemonFromAPI(string) (*models.PokemonResponse, error)

	GetRegistries(odd bool, itemsNumber int, workers int, pokemons []*models.Pokemon) ([]*models.Pokemon, error)
}

//function to init the usecase
func New(service Service) *UseCase {
	return &UseCase{service}
}
