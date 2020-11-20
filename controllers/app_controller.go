package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-bootcamp-2020/services"
	_errors "golang-bootcamp-2020/utils/error"
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

	res, err := ac.service.FetchData()
	if err != nil {
		c.JSON(err.Code(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ac *appController) GetCharacter(c *gin.Context) {

	characterId := c.Query("id")
	if characterId == "" {
		err := _errors.NewBadRequestError("id is required")
		c.JSON(err.Code(), err)
		return
	}

	ch, err := ac.service.GetCharacterById(characterId)
	if err != nil {
		c.JSON(err.Code(), err)
		return
	}

	c.JSON(http.StatusOK, ch)
}

func (ac *appController) GetCharacters(c *gin.Context) {

	characters, err := ac.service.GetAllCharacters()
	if err != nil {
		c.JSON(err.Code(), err)
		return
	}

	c.JSON(http.StatusOK, characters)
}

func (ac *appController) GetCharacterIdByName(c *gin.Context) {

	name := c.Query("name")
	if name == "" {
		err := _errors.NewBadRequestError("name is required")
		c.JSON(err.Code(), err)
		return
	}

	character, err := ac.service.GetCharacterIdByName(name)
	if err != nil {
		c.JSON(err.Code(), err)
		return
	}

	c.JSON(http.StatusOK, character)
}
