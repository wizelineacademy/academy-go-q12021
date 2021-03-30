package presenter

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/presenter"
)

// itemPresenter struct for ItemPresenter
type itemPresenter struct {}

// NewItemPresenter return an ItemPresenter
func NewItemPresenter() presenter.ItemPresenter {
	return &itemPresenter{}
}

// ResponseItems return an array of Item
func (ip *itemPresenter) ResponseItems(items []*model.Item) []*model.Item {
	for _, item := range items {
		item.Name = "Mr. " + item.Name
	}

	return items
}