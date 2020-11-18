package router

import (
	"github.com/gorilla/mux"
)

//NewRouter creates a new router instance
/*
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })

	return e
}
*/

//NewRouter creates a new router instance
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
