package main

import (
	"github.com/maestre3d/academy-go-q12021/pkg/controller"
	"github.com/maestre3d/academy-go-q12021/pkg/httputil"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.NopLogger,
		controller.HTTPModule,
		fx.Invoke(
			httputil.StartServer,
		),
	)
	app.Run()
}
