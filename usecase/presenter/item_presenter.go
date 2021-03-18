package presenter

import "bootcamp/domain/model"

type ItemPresenter interface {
	ResponseItems(u []*model.Item) []*model.Item
}
