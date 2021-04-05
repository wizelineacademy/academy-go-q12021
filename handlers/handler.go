package handlers

import (
	"github.com/cesararredondow/academy-go-q12021/models"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

//useCase is an interface of usecases
type UseCase interface {
	GetPokemons() ([]*models.Pokemon, error)
	GetPokemon(string) (*models.Pokemon, error)

	GetPokemonsFromAPI(quantity string) ([]*models.Pokemon_api, error)
	GetPokemonFromAPI(id string) (*models.PokemonResponse, error)

	GetPokemonsConcurrency(odd string, quantity string, workersNumer string) ([]*models.Pokemon, error)
}

// Pokemons struct
type Pokemons struct {
	useCase UseCase
	logger  *logrus.Logger
	render  *render.Render
}

// New returns a controller
func New(
	u UseCase,
	logger *logrus.Logger,
	r *render.Render,
) *Pokemons {
	return &Pokemons{u, logger, r}
}
