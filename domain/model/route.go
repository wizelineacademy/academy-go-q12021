package model

import "net/http"

type Route struct{
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route
