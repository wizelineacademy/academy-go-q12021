package store

import (
	"github.com/labstack/echo/v4"
	"github.com/pankecho/golang-bootcamp-2020/entity"
)

// Todo: add the data to the database, in the meantime we just return them
func (ps PersonStore) CreatePerson(ctx echo.Context, p *entity.Person) (*entity.Person, error) {
	return p, nil
}