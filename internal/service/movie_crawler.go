package service

import (
	"context"
	"time"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
	"go.uber.org/zap"
)

// MovieCrawler service connects the current application with external Movie APIs
type MovieCrawler interface {
	Fetch(context.Context, valueobject.MovieID) (*aggregate.Movie, error)
}

// NewMovieCrawler wraps the given Movie crawler with observability (logging, *metrics and *tracing)
// and *caching
//	Took as reference gokit middleware pattern
//
//	* Not yet implemented
func NewMovieCrawler(root MovieCrawler, logger *zap.Logger) MovieCrawler {
	return &movieCrawlerLogger{
		logger: logger,
		next:   root,
	}
}

type movieCrawlerLogger struct {
	logger *zap.Logger
	next   MovieCrawler
}

func (c *movieCrawlerLogger) Fetch(ctx context.Context, id valueobject.MovieID) (movie *aggregate.Movie, err error) {
	defer func(startTime time.Time) {
		fields := []zap.Field{
			zap.String("movie_imdb_id", string(id)),
			zap.Duration("took", time.Since(startTime)),
			zap.Bool("was_found", movie != nil),
		}
		if err != nil {
			fields = append(fields, zap.Error(err))
			c.logger.Error("failed to crawl movie from external API", fields...)
			return
		}

		c.logger.Info("fetched movie from external API", fields...)
	}(time.Now())

	movie, err = c.next.Fetch(ctx, id)
	return
}
