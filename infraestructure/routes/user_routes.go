package router

import "github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/handlers/users"

var userRoutes RoutePrefix = RoutePrefix{
	Prefix: "/users",
	Routes: []Route{
		{
			Path:        "/",
			Method:      "GET",
			IsProtected: false,
			Handler:     users.GetUsers,
		},
		{
			Name:        "Create User",
			Path:        "/",
			Method:      "POST",
			IsProtected: false,
			Handler:     users.GetUsers,
		},
	},
}
