package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppController interface {
	FetchData(c *gin.Context)
	GetHealth(c *gin.Context)
}

type appController struct{}

func NewAppController() AppController {
	return &appController{}
}

func (h *appController) GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func (h *appController) FetchData(c *gin.Context) {
	c.String(http.StatusOK, "fetching data...")
}
