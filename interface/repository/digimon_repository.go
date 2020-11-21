package repository

import (
	"digimons/domain/model"

	"gorm.io/gorm"
)

type digimonRepository struct {
	db *gorm.DB
}

type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
}

func NewDigimonRepository(db *gorm.DB) DigimonRepository {
	return &digimonRepository{db}
}

func (dr *digimonRepository) FindAll(d []*model.Digimon) ([]*model.Digimon, error) {
	err := dr.db.Find(&d).Error

	if err != nil {
		return nil, err
	}

	return d, nil
}
