package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-bootcamp-2020/services"
	"net/http"
)

type AppController interface {
	FetchData(c *gin.Context)
	GetHealth(c *gin.Context)
	GetCharacter(c *gin.Context)
	GetCharacters(c *gin.Context)
	GetCharacterIdByName(c *gin.Context)
}

type appController struct {
	service services.Service
}

func NewAppController(service services.Service) AppController {
	return &appController{service}
}

func (ac *appController) GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func (ac *appController) FetchData(c *gin.Context) {

	var response interface{}
	response = nil

	res, err := ac.service.FetchData()
	if err == nil {
		response = res
	}

	c.JSON(http.StatusOK, response)
}

func (ac *appController) GetCharacter(c *gin.Context) {

	characterId := c.Query("id")
	if characterId == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		return
	}

	ch, err := ac.service.GetCharacterById(characterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ch)
}

func (ac *appController) GetCharacters(c *gin.Context) {

	characters, err := ac.service.GetAllCharacters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, characters)
}

func (ac *appController) GetCharacterIdByName(c *gin.Context) {

	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, "name is required")
	}

	character, err := ac.service.GetCharacterIdByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, character)
}
