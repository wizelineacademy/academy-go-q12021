package controller

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/interactor"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type itemController struct {
	itemIterator interactor.ItemInteractor
}

type ItemController interface {
	GetItems(c Context) error
	GetItem(c echo.Context) error
	CreateItem(c Context) error
}

func NewItemController(us interactor.ItemInteractor) ItemController {
	return &itemController{us}
}

func (ic *itemController) GetItems(c Context) error {
	var i []*model.Item

	i, err := ic.itemIterator.Get(i)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, i)
}

func (ic *itemController) GetItem(c echo.Context) error {
	var i []*model.Item

	id, err := strconv.ParseUint(c.Param("id"),10, 64)

	i = append(i, &model.Item{ID: uint(id)})
	fmt.Println("Param === %s", id)

	i, err = ic.itemIterator.Get(i)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, i)
}

func (ic *itemController) CreateItem(c Context) error {
	var params model.Item

	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	i, err := ic.itemIterator.Create(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSON(http.StatusOK, i)
}