package repository

import "bootcamp/domain/model"

// ItemRepository interface for Items
type ItemRepository interface {
	// FindAll return all Items
	FindAll(u []*model.Item) ([]*model.Item, error)
	// Create adds new items to the datastore
	Create(u *model.Item) (*model.Item, error)
}
