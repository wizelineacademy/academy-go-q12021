package repository

import "github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}
