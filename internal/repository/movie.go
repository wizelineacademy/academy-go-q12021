package repository

import (
	"context"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// Movie lets interactions between Movie's persistence layer
type Movie interface {
	// Find retrieves a Movie by its ID
	Find(context.Context, valueobject.MovieID) (*aggregate.Movie, error)
	// Search retrieves a set of Movies by the given criteria filters, returns the set, a next page token or an error
	Search(context.Context, Criteria) ([]*aggregate.Movie, string, error)
	// Save stores the current state of the given Movie
	Save(context.Context, aggregate.Movie) error
}
