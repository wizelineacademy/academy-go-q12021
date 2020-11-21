package interactor

import (
	"digimons/domain/model"
	"digimons/usecase/presenter"
	"digimons/usecase/repository"
)

type digimonInteractor struct {
	DigimonRepository repository.DigimonRepository
	DigimonPresenter  presenter.DigimonPresenter
}

type DigimonInteractor interface {
	Get(d []*model.Digimon) ([]*model.Digimon, error)
}

func NewDigimonInteractor(r repository.DigimonRepository, p presenter.DigimonPresenter) DigimonInteractor {
	return &digimonInteractor{r, p}
}

func (di *digimonInteractor) Get(d []*model.Digimon) ([]*model.Digimon, error) {
	d, err := di.DigimonRepository.FindAll(d)
	if err != nil {
		return nil, err
	}

	return di.DigimonPresenter.ResponseDigimons(d), nil
}
