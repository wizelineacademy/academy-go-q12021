package presenter

import "bootcamp/domain/model"

// ItemPresenter interface for Items
type ItemPresenter interface {
	// ResponseItems returns manipulated Items
	ResponseItems(u []*model.Item) []*model.Item
}
