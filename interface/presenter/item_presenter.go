package presenter

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/presenter"
)

type itemPresenter struct {}

func NewItemPresenter() presenter.ItemPresenter {
	return &itemPresenter{}
}

func (ip *itemPresenter) ResponseItems(items []*model.Item) []*model.Item {
	for _, item := range items {
		item.Name = "Mr. " + item.Name
	}

	return items
}