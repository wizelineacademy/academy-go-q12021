package routes

import "github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/handlers"

var adminRoutes RoutePrefix = RoutePrefix{
	Prefix: "/admin",
	Routes: []Route{
		{
			Name:        "Find",
			Path:        "/logs/",
			Method:      "GET",
			IsProtected: true,
			Handler:     handlers.GetLogs,
		},
	},
}
