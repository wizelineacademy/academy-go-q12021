package routes

import "github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/handlers"

var userRoutes RoutePrefix = RoutePrefix{
	Prefix: "/users",
	Routes: []Route{
		{
			Name:        "Get Users",
			Path:        "/",
			Method:      "GET",
			IsProtected: false,
			Handler:     handlers.GetUsers,
		},
		{
			Name:        "Create User",
			Path:        "/",
			Method:      "POST",
			IsProtected: false,
			Handler:     handlers.CreateUser,
		},
	},
}
