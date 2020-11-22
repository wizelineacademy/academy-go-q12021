package repository

import (
	"digimons/domain/model"

	"gorm.io/gorm"
)

type digimonRepository struct {
	db *gorm.DB
}

// DigimonRepository shows all the methods to be implemented by a digimon repository
type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
}

// NewDigimonRepository Returns an instance of a digimon repository
func NewDigimonRepository(db *gorm.DB) DigimonRepository {
	return &digimonRepository{db}
}

// FindAll retrieve all the digimons from the database
func (dr *digimonRepository) FindAll(d []*model.Digimon) ([]*model.Digimon, error) {
	err := dr.db.Find(&d).Error

	if err != nil {
		return nil, err
	}

	return d, nil
}
