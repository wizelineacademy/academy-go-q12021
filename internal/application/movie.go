package application

import (
	"context"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/domain"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/internal/service"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// Movie holds all Movie's business use cases
type Movie struct {
	repo     repository.Movie
	eventBus domain.EventBus
	crawler  service.MovieCrawler
}

// NewMovie allocates a movie application
func NewMovie(r repository.Movie, eventBus domain.EventBus, crawler service.MovieCrawler) *Movie {
	return &Movie{
		repo:     r,
		eventBus: eventBus,
		crawler:  crawler,
	}
}

// GetByID fetches a movie (if any) by the given unique identifier
func (m *Movie) GetByID(ctx context.Context, id valueobject.MovieID) (*aggregate.Movie, error) {
	movie, err := m.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if movie == nil {
		return nil, aggregate.ErrMovieNotFound
	}

	return movie, nil
}

// List fetches a list of movies (if any) using the given criteria filters
func (m *Movie) List(ctx context.Context, criteria repository.Criteria) (movies []*aggregate.Movie,
	nextPage string, err error) {
	movies, nextPage, err = m.repo.Search(ctx, criteria)
	if err != nil {
		return nil, "", err
	} else if len(movies) == 0 {
		return nil, "", aggregate.ErrMovieNotFound
	}

	return movies, nextPage, nil
}

// Create add a new movie to the platform
func (m *Movie) Create(ctx context.Context, id valueobject.MovieID, title valueobject.DisplayName,
	releaseYear valueobject.ReleaseYear, imdbID valueobject.MovieID, directors ...valueobject.DisplayName) error {
	movie := aggregate.NewMovie(id, title, releaseYear, imdbID, directors...)

	if err := m.repo.Save(ctx, *movie); err != nil {
		return err
	}

	if m.eventBus != nil {
		err := m.eventBus.PublishEvents(movie.PullEvents()...)
		if err != nil {
			return err
		}
	}
	return nil
}

// CrawlAndSave fetches a Movie using iMDb ID and stores it locally
func (m *Movie) CrawlAndSave(ctx context.Context, id, imdbID valueobject.MovieID) error {
	movie, err := m.crawler.Fetch(ctx, imdbID)
	if err != nil {
		return err
	} else if movie == nil {
		return aggregate.ErrMovieNotFound
	}
	movie.ID = id

	return m.Create(ctx, movie.ID, movie.DisplayName, movie.ReleaseYear, movie.IMDbID, movie.Directors...)
}
