package services

import (
	"golang-bootcamp-2020/repository/db"
	"golang-bootcamp-2020/repository/rest"
)

type Service interface {
	FetchCharacters() (interface{}, error)
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

func (s *service) FetchCharacters() (interface{}, error) {
	return s.restRepo.GetCharacters()
}
