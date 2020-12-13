package services

import (
	"golang-bootcamp-2020/domain/model"
	"golang-bootcamp-2020/repository/db"
	"golang-bootcamp-2020/repository/rest"
	_errors "golang-bootcamp-2020/utils/error"
)

//Service - service methods
type Service interface {
	FetchData(maxPages int) ([]model.Character, _errors.RestError)
	GetCharacterById(id string) (*model.Character, _errors.RestError)
	GetAllCharacters() ([]model.Character, _errors.RestError)
	GetCharacterIdByName(name string) (string, _errors.RestError)
}

type service struct {
	restRepo rest.RickAndMortyAPIRepository
	dbRepo   db.DataBaseRepository
}

//NewService - return new service
func NewService(restRepo rest.RickAndMortyAPIRepository, dbRepo db.DataBaseRepository) Service {
	return &service{
		restRepo: restRepo,
		dbRepo:   dbRepo,
	}
}

//FetchData - Fetch data from rest repository and then make csv file
func (s *service) FetchData(maxPages int) ([]model.Character, _errors.RestError) {
	ch, err := s.restRepo.FetchData(maxPages)

	if err := s.dbRepo.CreateCharactersCSV(ch); err != nil {
		return nil, _errors.NewInternalServerError("error when trying to save data in csv")
	}

	return ch, err
}

//GetCharacterById - Get character from map (complex O(1) )
func (s *service) GetCharacterById(id string) (*model.Character, _errors.RestError) {
	return s.dbRepo.GetCharacterFromId(id)
}

//GetAllCharacters - Get all characters from map
func (s *service) GetAllCharacters() ([]model.Character, _errors.RestError) {
	return s.dbRepo.GetCharacters()
}

//GetCharacterIdByName - Get character id from csv map (complex O(n) )
func (s *service) GetCharacterIdByName(name string) (string, _errors.RestError) {
	return s.dbRepo.GetCharacterIdByName(name)
}
