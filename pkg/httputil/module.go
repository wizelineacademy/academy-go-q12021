package httputil

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

// Controller HTTP controller
type Controller interface {
	// Route maps exposed use cases from the current aggregate using the given mux.Router
	Route(r *mux.Router)
}

// ControllersGatewayFx Uber FX's HTTP controller input container
type ControllersGatewayFx struct {
	fx.In
	Controllers []Controller `group:"http_server"`
}

// ControllersFx Uber FX's HTTP controller output container
type ControllersFx struct {
	fx.Out
	Controller Controller `group:"http_server"`
}
