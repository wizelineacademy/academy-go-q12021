package presenter

import "digimons/domain/model"

type digimonPresenter struct {
}

type DigimonPresenter interface {
	ResponseDigimons(di []*model.Digimon) []*model.Digimon
}

func NewDigimonPresenter() DigimonPresenter {
	return &digimonPresenter{}
}

func (dp *digimonPresenter) ResponseDigimons(ds []*model.Digimon) []*model.Digimon {
	for _, d := range ds {
		d.Name = d.Name
	}

	return ds
}
