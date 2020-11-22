package repository

import (
	"digimons/domain/model"

	"gorm.io/gorm"
)

type digimonRepository struct {
	db *gorm.DB
}

// DigimonRepository
type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
}

// NewDigimonRepository
func NewDigimonRepository(db *gorm.DB) DigimonRepository {
	return &digimonRepository{db}
}

// FindAll
func (dr *digimonRepository) FindAll(d []*model.Digimon) ([]*model.Digimon, error) {
	err := dr.db.Find(&d).Error

	if err != nil {
		return nil, err
	}

	return d, nil
}
