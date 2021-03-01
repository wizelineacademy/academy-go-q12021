package controller

import (
	"github.com/maestre3d/academy-go-q12021/internal/module"
	"github.com/maestre3d/academy-go-q12021/pkg/httputil"
	"go.uber.org/fx"
)

// HTTPModule holds a set of controllers for Uber Fx container
var HTTPModule = fx.Options(
	module.Kernel,
	module.Movie,
	fx.Provide(
		NewMovieHTTP,
		httputil.NewGorillaRouter,
	),
)
