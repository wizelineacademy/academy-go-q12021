package persistence

import (
	"context"
	"sync"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// MovieInMemory handles all persistence Movie's operations locally using local memory
type MovieInMemory struct {
	mu sync.RWMutex
	db map[string]*aggregate.Movie
}

// NewMovieInMemory allocates a Movie repository In Memory concrete implementation
func NewMovieInMemory() *MovieInMemory {
	return &MovieInMemory{
		mu: sync.RWMutex{},
		db: make(map[string]*aggregate.Movie),
	}
}

// Get retrieves a Movie by its ID
func (m *MovieInMemory) Get(ctx context.Context, id valueobject.MovieID) (*aggregate.Movie, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.db[string(id)], nil
}

// Search retrieves a set of Movies by the given criteria filters, returns the set, a next page token or an error
func (m *MovieInMemory) Search(ctx context.Context, criteria repository.Criteria) ([]*aggregate.Movie, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	movies := make([]*aggregate.Movie, 0)
	for _, m := range m.db {
		if len(movies) == criteria.Limit {
			break
		}
		movies = append(movies, m)
	}
	return movies, "", nil
}

// Save stores the current state of the given Movie
func (m *MovieInMemory) Save(ctx context.Context, movie aggregate.Movie) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.db[string(movie.ID)] = &movie
	return nil
}
