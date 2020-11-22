package repository

import "digimons/domain/model"

// DigimonRepository this is the interface that digimon repository should implement.
type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
}
