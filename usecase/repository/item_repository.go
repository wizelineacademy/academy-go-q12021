package repository

import (
	"bootcamp/domain/model"
	"bootcamp/interface/controller/vo"
)

// ItemRepository interface for Items
type ItemRepository interface {
	// FindAll return all Items
	FindAll(u []*model.Item) ([]*model.Item, error)
	// FindAllPaged return all Items paged
	FindAllPaged(u []*model.Item, paged *vo.Paged) ([]*model.Item, error)
	// Create adds new items to the datastore
	Create(u *model.Item) (*model.Item, error)
}
