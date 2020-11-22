package presenter

import "digimons/domain/model"

type digimonPresenter struct{}

// DigimonPresenter contain the methods that are called before massing it to the views that refer to Digimons.
type DigimonPresenter interface {
	ResponseDigimons(di []*model.Digimon) []*model.Digimon
}

// NewDigimonPresenter returns a pointer to DigimonPresenter
func NewDigimonPresenter() DigimonPresenter {
	return &digimonPresenter{}
}

// ResponseDigimons This method handle all the data before passing it to view.
func (dp *digimonPresenter) ResponseDigimons(ds []*model.Digimon) []*model.Digimon {
	return ds
}
