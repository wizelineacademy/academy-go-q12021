package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oscarSantoyo/academy-go-q12021/controller"
)

type errorMessage struct {
	Message string `json:"message"`
	ErrorId int    `json:"errorId"`
}

func SetRouter(e *echo.Echo) {
	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusAccepted, "I'm Up")
	})

	e.GET("/filter", func(c echo.Context) error {
		id := c.QueryParam("id")
		if id == "" {
			return reportError(c, "Missing parameter 'id'")
		}

		result := controller.SearchById(id)
		if len(result) == 0 {
			return reportError(c, "No records found")
		}
		return c.JSON(http.StatusAccepted, result)
	})

	e.GET("/loadData", func(c echo.Context) error{
		err := controller.LoadCsvData()
		if err != nil {
			return reportError(c, "Data was not able to load")
		}
		return c.JSON(http.StatusAccepted, "loadded correctly")
	})
}

func reportError(c echo.Context, msg string) error {
	return c.JSON(http.StatusInternalServerError, &errorMessage{
		Message: msg,
		ErrorId: http.StatusInternalServerError,
	})
}
