package controller

import (
	"net/http"

	"digimons/domain/model"
	"digimons/usecase/interactor"
)

type digimonController struct {
	digimonController interactor.DigimonInteractor
}

type DigimonController interface {
	GetDigimons(c Context) error
}

func NewDigimonController(di interactor.DigimonInteractor) DigimonController {
	return &digimonController{di}
}

func (di *digimonController) GetDigimons(c Context) error {
	var d []*model.Digimon

	d, err := di.digimonController.Get(d)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, d)
}
