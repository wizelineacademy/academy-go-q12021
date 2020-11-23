package repository

import (
	"api-booking-time/domain/model"
	"api-booking-time/usecase/repository"
)

type centreRepository struct {
	centres *[]*model.Centre
}

func OpenCentreRepository(centres *[]*model.Centre) repository.CentreRepository {
	return &centreRepository{centres}
}

func (cr *centreRepository) GetAll() (*[]*model.Centre, error) {
	return cr.centres, nil
}