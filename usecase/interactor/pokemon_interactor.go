package interactor

import (
	"github.com/AlejandroSeguraWIZ/academy-go-q12021/domain/model"
	"github.com/AlejandroSeguraWIZ/academy-go-q12021/usecase/presenter"
	"github.com/AlejandroSeguraWIZ/academy-go-q12021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonPresenter  presenter.PokemonPresenter
	PokemonRepository repository.PokemonRepository
}

type PokemonInteractor interface {
	Get() ([]model.PokemonResponse, error)
}

func NewPokemonInteractor(repository repository.PokemonRepository, presenter presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{
		presenter,
		repository,
	}
}

func (i *pokemonInteractor) Get() ([]model.PokemonResponse, error) {
	pokemonList, err := i.PokemonRepository.FetchAll()
	if err != nil {
		return []model.PokemonResponse{}, err
	}
	return i.PokemonPresenter.ResponsePokemons(pokemonList), nil
}
