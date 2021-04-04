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
	GetMovies() ([]*model.MovieSummary, error)
	GetMovieById(movieID string) (*model.Movie, error)
	GetMoviesConcurrently(queryParams model.QueryParameters, complete bool, id string) ([]interface{}, error)
}

// New UseCasequeryParams
func New(service Service) *UseCase {
	return &UseCase{service}
}

// GetMoviesConcurrently -
func (u *UseCase) GetMoviesConcurrently(queryParams model.QueryParameters, complete bool, id string) ([]interface{}, error) {

	resp, err := u.service.GetMoviesConcurrently(queryParams, complete, id)
	if err != nil {
		return nil, err
	}
	log.Println("u.service.GetMoviesConcurrently: ", len(resp), queryParams, complete, id)

	return resp, nil
}

// GetMovies -
func (u *UseCase) GetMovies() ([]*model.MovieSummary, error) {
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
