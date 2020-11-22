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

// DigimonInteractor contain the methods that the interactor should contain
type DigimonInteractor interface {
	Get(d []*model.Digimon) ([]*model.Digimon, error)
}

// NewDigimonInteractor returns a digimon interactor based on the repository and a presenter.
func NewDigimonInteractor(r repository.DigimonRepository, p presenter.DigimonPresenter) DigimonInteractor {
	return &digimonInteractor{r, p}
}

// Get begin the process of retrieving all the digimon data as it is called from Input Port.
func (di *digimonInteractor) Get(d []*model.Digimon) ([]*model.Digimon, error) {
	d, err := di.DigimonRepository.FindAll(d)
	if err != nil {
		return nil, err
	}

	return di.DigimonPresenter.ResponseDigimons(d), nil
}
