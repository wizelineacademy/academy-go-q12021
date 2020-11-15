package services

import (
	"golang-bootcamp-2020/repository/db"
	"golang-bootcamp-2020/repository/rest"
)

type Service interface {
}

type service struct {
	restRepo rest.RickAndMortyApiRepository
	dbRepo   db.DataBaseRepository
}

func NewService(restRepo rest.RickAndMortyApiRepository, dbRepo db.DataBaseRepository) Service {
	return &service{
		restRepo: restRepo,
		dbRepo:   dbRepo,
	}
}
