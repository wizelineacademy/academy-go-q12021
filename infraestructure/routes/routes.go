package routes

import (
	"net/http"

	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/middleware"
	"github.com/gorilla/mux"
)

//RoutePrefix maps a subroute of the system. Eg: /users/...
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

			if subroute.IsProtected {
				subroute.Handler = middleware.ValidateToken(subroute.Handler)
			}

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
	AppRoutes = append(AppRoutes, songRoutes)
}
