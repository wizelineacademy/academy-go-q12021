package application

import (
	"context"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/domain"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// Movie holds all Movie's business use cases
type Movie struct {
	repo     repository.Movie
	eventBus domain.EventBus
}

// NewMovie allocates a movie application
func NewMovie(r repository.Movie, eventBus domain.EventBus) *Movie {
	return &Movie{
		repo:     r,
		eventBus: eventBus,
	}
}

// GetByID fetches a movie (if any) by the given unique identifier
func (m Movie) GetByID(ctx context.Context, id valueobject.MovieID) (*aggregate.Movie, error) {
	movie, err := m.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	} else if movie == nil {
		return nil, aggregate.ErrMovieNotFound
	}

	return movie, nil
}

// List feteches a list of movies (if any) using the given criteria filters
func (m Movie) List(ctx context.Context, criteria repository.Criteria) (movies []*aggregate.Movie,
	nextPage string, err error) {
	movies, nextPage, err = m.repo.Search(ctx, criteria)
	if err != nil {
		return nil, "", err
	} else if len(movies) == 0 {
		return nil, "", aggregate.ErrMovieNotFound
	}

	return movies, nextPage, nil
}
