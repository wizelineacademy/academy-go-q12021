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
			Handler:     handlers.GetSongData,
		},
		{
			Name:        "Find",
			Path:        "/artist/{artistId}/album/{albumId}/track/{trackId}",
			Method:      "GET",
			IsProtected: true,
			Handler:     handlers.GetSongLyrics,
		},
	},
}
