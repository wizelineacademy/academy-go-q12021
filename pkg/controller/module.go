package controller

import (
	"github.com/maestre3d/academy-go-q12021/internal/application"
	"github.com/maestre3d/academy-go-q12021/internal/domain"
	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
	"github.com/maestre3d/academy-go-q12021/internal/persistence"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/pkg/httputil"
	"go.uber.org/fx"
)

// HTTPModule holds a set of controllers for Uber Fx container
var HTTPModule = fx.Provide(
	infrastructure.NewConfiguration,
	infrastructure.NewZapLogger,
	func(cfg infrastructure.Configuration) repository.Movie {
		return persistence.NewMovieCSV(cfg)
	},
	func() domain.EventBus {
		return nil
	},
	application.NewMovie,
	NewMovieHTTP,
	httputil.NewGorillaRouter,
)
