package model

import "net/http"

// Router implements a struct with a URL path (route) information
type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

// Routes implements an array of Route
type Routes []Route
