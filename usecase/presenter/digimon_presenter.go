package presenter

import "digimons/domain/model"

type DigimonPresenter interface {
	ResponseDigimons(d []*model.Digimon) []*model.Digimon
}
