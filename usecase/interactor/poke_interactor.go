package interactor

import (
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	"github.com/ToteEmmanuel/academy-go-q12021/usecase/presenter"
	"github.com/ToteEmmanuel/academy-go-q12021/usecase/repository"
)

type pokeInteractor struct {
	PokeRepo      repository.PokeRepository
	PokePresenter presenter.PokePresenter
}

type PokeInteractor interface {
	Get(id int32) (*model.Pokemon, error)
	GetAll() []*model.Pokemon
}

func NewPokeInteractor(r repository.PokeRepository, p presenter.PokePresenter) PokeInteractor {
	return &pokeInteractor{
		PokeRepo:      r,
		PokePresenter: p,
	}
}

func (pI *pokeInteractor) Get(id int32) (*model.Pokemon, error) {
	p, err := pI.PokeRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return pI.PokePresenter.ResponsePokemon(p), nil
}

func (pI *pokeInteractor) GetAll() []*model.Pokemon {
	p := pI.PokeRepo.FindAll()
	return pI.PokePresenter.ResponsePokemons(p)
}
