package presenter

import "digimons/domain/model"

// DigimonPresenter this is the interface of the methods that digimon presenter should has
type DigimonPresenter interface {
	ResponseDigimons(d []*model.Digimon) []*model.Digimon
}
