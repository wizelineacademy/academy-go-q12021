package controller

import (
	"net/http"

	"digimons/domain/model"
	"digimons/usecase/interactor"
)

type digimonController struct {
	digimonController interactor.DigimonInteractor
}

// DigimonController This interface will handle the request for Digimons that comes from outer layers
type DigimonController interface {
	GetDigimons(c Context) error
}

// NewDigimonController This function returns a Digimon controller based on the interactor which catch the request for digimon data
func NewDigimonController(di interactor.DigimonInteractor) DigimonController {
	return &digimonController{di}
}

// GetDigimons retrieves data from Digimons and return response as json or shows and error.
func (di *digimonController) GetDigimons(c Context) error {
	var d []*model.Digimon

	d, err := di.digimonController.Get(d)
	if err != nil {
		return c.JSON(http.StatusNotFound, d)
	}

	return c.JSON(http.StatusOK, d)
}
