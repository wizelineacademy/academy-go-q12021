package routes

import "github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/handlers"

var songRoutes RoutePrefix = RoutePrefix{
	Prefix: "/songs",
	Routes: []Route{
		{
			Name:        "Find",
			Path:        "/",
			Method:      "GET",
			IsProtected: true,
			Handler:     handlers.GetSong,
		},
		{
			Name:        "Find",
			Path:        "/artist/",
			Method:      "GET",
			IsProtected: true,
			Handler:     handlers.GetSong,
		},
		{
			Name:        "Find",
			Path:        "/lyric",
			Method:      "GET",
			IsProtected: true,
			Handler:     handlers.GetSong,
		},
	},
}
