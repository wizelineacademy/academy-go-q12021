package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RoutePrefix struct {
	Prefix string
	Routes []Route
}

//Route maps the structure of an application route
type Route struct {
	Name        string
	Path        string
	Method      string
	Handler     http.HandlerFunc
	IsProtected bool
}

//AppRoutes stores the application routes
var AppRoutes []RoutePrefix

//AddRoutes adds the routes to the main router
func AddRoutes(router *mux.Router) {

	for _, route := range AppRoutes {
		subrouter := router.PathPrefix(route.Prefix).Subrouter()

		for _, subroute := range route.Routes {
			subrouter.
				Path(subroute.Path).
				Handler(subroute.Handler).
				Methods(subroute.Method).
				Name(subroute.Name)
		}
	}
}

func init() {
	AppRoutes = append(AppRoutes, userRoutes)
}
