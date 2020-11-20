package services

import (
	"golang-bootcamp-2020/domain/model"
	"golang-bootcamp-2020/repository/db"
	"golang-bootcamp-2020/repository/rest"
	_errors "golang-bootcamp-2020/utils/error"
)

type Service interface {
	FetchData() ([]model.Character, _errors.RestError)
	GetCharacterById(id string) (*model.Character, _errors.RestError)
	GetAllCharacters() ([]model.Character, _errors.RestError)
	GetCharacterIdByName(name string) (string, _errors.RestError)
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

func (s *service) FetchData() ([]model.Character, _errors.RestError) {
	//TODO: hanle this error correctly
	ch, err := s.restRepo.FetchData()

	s.dbRepo.CreateCharactersCSV(ch)

	return ch, err
}

func (s *service) GetCharacterById(id string) (*model.Character, _errors.RestError) {
	return s.dbRepo.GetCharacterFromId(id)
}

func (s *service) GetAllCharacters() ([]model.Character, _errors.RestError) {
	return s.dbRepo.GetCharacters()
}

func (s *service) GetCharacterIdByName(name string) (string, _errors.RestError) {
	return s.dbRepo.GetCharacterIdByName(name)
}
