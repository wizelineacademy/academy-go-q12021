package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
	"github.com/maestre3d/academy-go-q12021/internal/marshal"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// MovieCrawlerOmdb implements MovieCrawler service using the OMDb public API
type MovieCrawlerOmdb struct {
	cfg infrastructure.Configuration
	mu  sync.RWMutex
}

// NewMovieCrawlerOmdb allocates a MovieCrawlerOmdb implementation
func NewMovieCrawlerOmdb(cfg infrastructure.Configuration) *MovieCrawlerOmdb {
	return &MovieCrawlerOmdb{
		cfg: cfg,
		mu:  sync.RWMutex{},
	}
}

func (c *MovieCrawlerOmdb) Fetch(ctx context.Context, imdbID valueobject.MovieID) (*aggregate.Movie, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	resp, err := http.Get("http://www.omdbapi.com/?apikey=" + c.cfg.OmdbAPIKey + "&i=" + string(imdbID))
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
	}()
	return c.marshalFromHTTPResponse(resp.Body)
}

func (c *MovieCrawlerOmdb) marshalFromHTTPResponse(res io.ReadCloser) (*aggregate.Movie, error) {
	movieOmdb := marshal.MovieOmdb{}
	err := json.NewDecoder(res).Decode(&movieOmdb)
	if err != nil {
		return nil, err
	}

	movie := aggregate.NewEmptyMovie()
	err = marshal.UnmarshalMovieOmdb(movieOmdb, movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
