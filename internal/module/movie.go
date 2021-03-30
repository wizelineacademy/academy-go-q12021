package module

import (
	"github.com/maestre3d/academy-go-q12021/internal/application"
	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
	"github.com/maestre3d/academy-go-q12021/internal/persistence"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/internal/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Movie module
var Movie = fx.Provide(
	func(cfg infrastructure.Configuration, logger *zap.Logger) repository.Movie {
		return repository.NewMovie(persistence.NewMovieCSV(cfg), logger)
	},
	func(cfg infrastructure.Configuration, logger *zap.Logger) service.MovieCrawler {
		return service.NewMovieCrawler(service.NewMovieCrawlerOmdb(cfg), logger)
	},
	application.NewMovie,
)
