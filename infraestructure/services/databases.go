package services

import "github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"

type Database interface {
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete() error
	Get(*model.User) (*model.User, error)
}
