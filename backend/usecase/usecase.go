package usecase

import (
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
}

// New UseCase
func New(service Service) *UseCase {
	return &UseCase{service}
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
