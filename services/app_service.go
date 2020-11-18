package services

import (
	"golang-bootcamp-2020/domain/model"
	"golang-bootcamp-2020/repository/db"
	"golang-bootcamp-2020/repository/rest"
)

type Service interface {
	FetchData() ([]model.Character, error)
	GetCharacterById(id string) (*model.Character, error)
	GetAllCharacters() ([]model.Character, error)
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

func (s *service) FetchData() ([]model.Character, error) {
	//TODO: hanle this error correctly
	ch, err := s.restRepo.FetchData()

	s.dbRepo.CreateCharactersCSV(ch)

	return ch, err
}

func (s *service) GetCharacterById(id string) (*model.Character, error) {
	return s.dbRepo.GetCharacterFromId(id)
}

func (s *service) GetAllCharacters() ([]model.Character, error) {
	return s.dbRepo.GetCharacters()
}
