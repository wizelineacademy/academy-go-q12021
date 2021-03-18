package repository

import "bootcamp/domain/model"

type ItemRepository interface {
	FindAll(u []*model.Item) ([]*model.Item, error)
	Create(u *model.Item) (*model.Item, error)
}