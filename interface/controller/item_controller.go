package controller

import (
	"bootcamp/domain/model"
	"bootcamp/interface/controller/vo"
	"bootcamp/usecase/interactor"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// itemController struct for ItemInteractor
type itemController struct {
	itemIterator interactor.ItemInteractor
}

// ItemController interface
type ItemController interface {
	GetItems(c echo.Context) error
	GetItem(c echo.Context) error
	CreateItem(c Context) error
}

// NewItemController return an ItemController
func NewItemController(us interactor.ItemInteractor) ItemController {
	return &itemController{us}
}

// GetItems return an array of Item
func (ic *itemController) GetItems(c echo.Context) error {
	var i []*model.Item

	paged, err := getPaged(c)
	if err != nil {
		return err
	}

	fmt.Println("Paged...", paged)

	i, err = ic.itemIterator.GetItems(i, paged)

	if err != nil {
		return err
	}

	fmt.Println("GetItems...")
	return c.JSON(http.StatusOK, i)
}

// getPaged returns a validated paged
func getPaged(c echo.Context) (*vo.Paged, error) {

	// if there are not query params returns nils
	if len(c.QueryParams()) == 0 {
		return nil, nil
	}

	paged := new(vo.Paged)

	if err := c.Bind(paged); err != nil {
		return nil, err
	}

	if err := c.Validate(paged); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return paged, nil
}


// GetItem return an Item
func (ic *itemController) GetItem(c echo.Context) error {
	var i []*model.Item

	id, err := strconv.ParseUint(c.Param("id"),10, 64)

	i = append(i, &model.Item{ID: uint(id)})
	fmt.Println("Param === %s", id)

	i, err = ic.itemIterator.GetItems(i,nil)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, i)
}

// CreateItem creates an item in the data store
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
