package repository

import (
	"context"
	"time"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"

	"go.uber.org/zap"
)

// Movie lets interactions between Movie's persistence layer
type Movie interface {
	// Get retrieves a Movie by its ID
	Get(context.Context, valueobject.MovieID) (*aggregate.Movie, error)
	// Search retrieves a set of Movies by the given criteria filters, returns the set, a next page token or an error
	Search(context.Context, Criteria) ([]*aggregate.Movie, string, error)
	// Save stores the current state of the given Movie
	Save(context.Context, aggregate.Movie) error
}

type movielogger struct {
	logger *zap.Logger
	next   Movie
}

// NewMovie wraps the given Movie repository with observability (logging, *metrics and *tracing)
// and *caching
//	Took as reference gokit middleware pattern
//
//	* Not yet implemented
func NewMovie(root Movie, logger *zap.Logger) Movie {
	return &movielogger{
		logger: logger,
		next:   root,
	}
}

func (l *movielogger) Get(ctx context.Context, id valueobject.MovieID) (movie *aggregate.Movie, err error) {
	defer func(startTime time.Time) {
		fields := []zap.Field{zap.String("movie_id", string(id)), zap.Duration("took", time.Since(startTime))}
		if err != nil {
			fields = append(fields, zap.Error(err))
			l.logger.Error("failed to fetch movie", fields...)
			return
		}

		fields = append(fields, zap.String("movie_imdb_id", string(movie.IMDbID)),
			zap.String("movie_display_name", string(movie.DisplayName)))
		l.logger.Info("fetched movie", fields...)
	}(time.Now())

	movie, err = l.next.Get(ctx, id)
	return
}

func (l *movielogger) Search(ctx context.Context, criteria Criteria) (movies []*aggregate.Movie,
	nextToken string, err error) {
	defer func(startTime time.Time) {
		fields := marshalCriteriaFieldsLog(criteria)
		if err != nil {
			fields = append(fields, zap.Error(err), zap.Duration("took", time.Since(startTime)))
			l.logger.Error("failed to fetch movies", fields...)
			return
		}

		fields = append(fields, zap.String("next_page", nextToken), zap.Int("total_items", len(movies)),
			zap.Duration("took", time.Since(startTime)))
		l.logger.Info("fetched movies", fields...)
	}(time.Now())
	movies, nextToken, err = l.next.Search(ctx, criteria)
	return
}

func (l *movielogger) Save(ctx context.Context, movie aggregate.Movie) (err error) {
	defer func(startTime time.Time) {
		fields := []zap.Field{
			zap.String("movie_id", string(movie.ID)),
			zap.String("movie_imdb_id", string(movie.IMDbID)),
			zap.String("movie_display_name", string(movie.DisplayName)),
			zap.Duration("took", time.Since(startTime)),
		}
		if err != nil {
			fields = append(fields, zap.Error(err))
			l.logger.Error("failed to save movie state", fields...)
			return
		}

		l.logger.Info("saved movie state", fields...)
	}(time.Now())

	err = l.next.Save(ctx, movie)
	return
}
