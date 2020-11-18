package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RoutePrefix struct {
	Prefix string
	Routes []Route
}

type Route struct {
	Path        string
	Method      string
	IsProtected bool
	Handler     http.HandlerFunc
}

var AppRoutes []RoutePrefix

func AddRoutes(router *mux.Router) {
	AppRoutes = append(AppRoutes, userRoutes)

	for _, route := range AppRoutes {
		subrouter := router.NewRoute().Subrouter()
		subrouter.PathPrefix(route.Prefix)

		for _, subroute := range route.Routes {

		}
	}
}
