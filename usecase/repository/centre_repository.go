package repository

import "api-booking-time/domain/model"

type CentreRepository interface {
	GetAll() (*[]*model.Centre, error)
}