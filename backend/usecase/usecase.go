package usecase

import (
	"log"
	"main/model"
)

// UseCase struct
type UseCase struct {
	service Service
}

// Service interface
type Service interface {
	GetMovies() ([]*model.Movie, error)
	GetMovieById(movieID string) (*model.Movie, error)
	GetConcurrently(queryParams model.QueryParameters, complete bool, id string) ([]interface{}, error)
}

// New UseCasequeryParams
func New(service Service) *UseCase {
	return &UseCase{service}
}

// GetConcurrently -
func (u *UseCase) GetConcurrently(queryParams model.QueryParameters, complete bool, id string) ([]interface{}, error) {

	resp, err := u.service.GetConcurrently(queryParams, complete, id)
	if err != nil {
		return nil, err
	}
	log.Println("u.service.GetConcurrently: ", len(resp), queryParams, complete, id)

	return resp, nil
}

// GetMovies -
func (u *UseCase) GetMovies() ([]*model.Movie, error) {
	resp, err := u.service.GetMovies()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetMovieById -
func (u *UseCase) GetMovieById(movieID string) (*model.Movie, error) {
	resp, err := u.service.GetMovieById(movieID)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
