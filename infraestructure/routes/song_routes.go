package routes

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/handlers/users"
	"github.com/alexis-aguirre/golang-bootcamp-2020/interface/controllers"
)

var songRoutes RoutePrefix = RoutePrefix{
	Prefix: "/songs",
	Routes: []Route{
		{
			Name:        "Get Song",
			Path:        "/",
			Method:      "GET",
			IsProtected: true,
			Handler:     controllers.GetSong
		},
	},
}
