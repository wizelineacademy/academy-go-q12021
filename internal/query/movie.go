package query

import (
	"context"

	"github.com/maestre3d/academy-go-q12021/internal/application"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// GetMovieByID request a fetch of a Movie by its unique identifier
type GetMovieByID struct {
	ID string `json:"movie_id"`
}

// HandleGetMovieByID GetMovieByID query handler
func HandleGetMovieByID(ctx context.Context, app *application.Movie, q GetMovieByID) (MovieResponse, error) {
	id, err := valueobject.NewMovieID(q.ID)
	if err != nil {
		return MovieResponse{}, err
	}

	movie, err := app.GetByID(ctx, id)
	if err != nil {
		return MovieResponse{}, err
	}
	return marshalMovieResponse(movie), nil
}

// ListMovies request a fetch of a list of Movies by the given criteria filters
type ListMovies struct {
	Criteria repository.Criteria `json:"criteria"`
}

// HandleListMovies ListMovies query handler
func HandleListMovies(ctx context.Context, app *application.Movie, q ListMovies) (MoviesResponse, error) {
	movies, nextPage, err := app.List(ctx, q.Criteria)
	if err != nil {
		return MoviesResponse{}, err
	}
	return marshalMoviesResponse(nextPage, movies...), nil
}
