package repository

import "digimons/domain/model"

type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
}
