package interactor

import (
	"github.com/Topi99/academy-go-q12021/domain/model"
	"github.com/Topi99/academy-go-q12021/usecase/presenter"
	"github.com/Topi99/academy-go-q12021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

// PokemonInteractor interface
type PokemonInteractor interface {
	GetOne(id uint) (*model.Pokemon, error)
}

// NewPokemonInteractor returns new PokemonInteractor
func NewPokemonInteractor(
	r repository.PokemonRepository, p presenter.PokemonPresenter,
) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (po *pokemonInteractor) GetOne(id uint) (*model.Pokemon, error) {
	p, err := po.PokemonRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return p, nil
}
