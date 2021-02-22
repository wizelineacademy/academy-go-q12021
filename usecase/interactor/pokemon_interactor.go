package interactor

import (
	"pokeapi/domain/model"
	"pokeapi/usecase/presenter"
	"pokeapi/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	Get(u []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (ps *pokemonInteractor) Get(p []*model.Pokemon) ([]*model.Pokemon, error) {
	p, err := ps.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}

	return ps.PokemonPresenter.ResponsePokemons(p), nil
}
