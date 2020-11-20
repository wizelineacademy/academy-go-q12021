package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-bootcamp-2020/services"
	_errors "golang-bootcamp-2020/utils/error"
	"net/http"
	"strconv"
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

//Return new pointer to application controller
func NewAppController(service services.Service) AppController {
	return &appController{service}
}

//Return ok if api works correctly
func (ac *appController) GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

//Fetch data from rick and morty api
func (ac *appController) FetchData(c *gin.Context) {

	var maxPages int
	maxPagesParam := c.Query("maxPages")
	if maxPagesParam == "" {
		maxPages = 0
	} else {
		var convErr error
		maxPages, convErr = strconv.Atoi(maxPagesParam)

		if convErr != nil || maxPages < 1 {
			restErr := _errors.NewBadRequestError("id param must be integer bigger than 0")
			c.JSON(restErr.Code(), restErr)
			return
		}
	}

	res, err := ac.service.FetchData(maxPages)
	if err != nil {
		c.JSON(err.Code(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}

//Get character by id
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

//Get all characters from map
func (ac *appController) GetCharacters(c *gin.Context) {

	characters, err := ac.service.GetAllCharacters()
	if err != nil {
		c.JSON(err.Code(), err)
		return
	}

	c.JSON(http.StatusOK, characters)
}

//Get character id by name from csv map
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
