package transport

import "github.com/labstack/echo/v4"

type PersonTransport interface {
	UploadCSV(c echo.Context) error
}

func NewRouter(pT PersonTransport) *echo.Echo {
	e := echo.New()

	person := e.Group("/person")
	person.POST("/csv", pT.UploadCSV)

	return e
}