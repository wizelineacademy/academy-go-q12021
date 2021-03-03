package persistence

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"sync"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
	"github.com/maestre3d/academy-go-q12021/internal/marshal"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// MovieCSV handles all persistence Movie's operations locally using an specific `.csv` file
//	Implements Movie repository
type MovieCSV struct {
	mu  sync.RWMutex
	cfg infrastructure.Configuration
}

// NewMovieCSV allocates a Movie repository CSV concrete implementation
func NewMovieCSV(config infrastructure.Configuration) *MovieCSV {
	return &MovieCSV{
		mu:  sync.RWMutex{},
		cfg: config,
	}
}

// Get retrieves a Movie by its ID
func (m *MovieCSV) Get(ctx context.Context, id valueobject.MovieID) (*aggregate.Movie, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	file, err := os.Open(m.cfg.MoviesDataset)
	if err != nil {
		return nil, err
	}

	return m.searchMovieOnFile(csv.NewReader(file), id)
}

func (m *MovieCSV) searchMovieOnFile(r *csv.Reader, id valueobject.MovieID) (*aggregate.Movie, error) {
	isHeader := true
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		} else if isHeader {
			isHeader = false
			continue
		} else if err != nil {
			return nil, err
		}

		movie := aggregate.NewEmptyMovie()
		if err = marshal.UnmarshalMovieCSV(movie, records...); err != nil {
			return nil, err
		} else if movie.ID == id {
			return movie, nil
		}
	}

	return nil, nil
}

// Search retrieves a set of Movies by the given criteria filters, returns the set, a next page token or an error
func (m *MovieCSV) Search(ctx context.Context, criteria repository.Criteria) ([]*aggregate.Movie, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	file, err := os.Open(m.cfg.MoviesDataset)
	if err != nil {
		return nil, "", err
	}
	return m.searchMoviesOnFile(csv.NewReader(file), criteria)
}

func (m *MovieCSV) searchMoviesOnFile(r *csv.Reader, criteria repository.Criteria) ([]*aggregate.Movie, string, error) {
	isHeader := true
	totalCount := 0
	allowFetch := false
	movies := make([]*aggregate.Movie, 0)
	nextPageToken := ""
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, "", err
		} else if isHeader {
			isHeader = false
			continue
		} else if criteria.NextPage != "" && records[0] != criteria.NextPage && allowFetch == false {
			continue
		}

		movie := aggregate.NewEmptyMovie()
		if err = marshal.UnmarshalMovieCSV(movie, records...); err != nil {
			return nil, "", err
		} else if totalCount >= criteria.Limit { // fetch one more item to get next page
			nextPageToken = string(movie.ID)
			break
		}

		movies = append(movies, movie)
		totalCount++
		allowFetch = true
	}

	return movies, nextPageToken, nil
}

// Save stores the current state of the given Movie
func (m *MovieCSV) Save(ctx context.Context, movie aggregate.Movie) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	panic("not implemented")
}
