package module

import (
	"github.com/maestre3d/academy-go-q12021/internal/application"
	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
	"github.com/maestre3d/academy-go-q12021/internal/persistence"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Movie module
var Movie = fx.Provide(
	func(cfg infrastructure.Configuration, logger *zap.Logger) repository.Movie {
		return repository.NewMovie(persistence.NewMovieCSV(cfg), logger)
	},
	application.NewMovie,
)
