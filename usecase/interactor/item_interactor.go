package interactor

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/presenter"
	"bootcamp/usecase/repository"
	"errors"
)

type itemInteractor struct {
	ItemRepository repository.ItemRepository
	ItemPresenter presenter.ItemPresenter
	DBRepository repository.DBRepository
}

type ItemInteractor interface {
	Get(u []*model.Item) ([]*model.Item, error)
	Create(u *model.Item) (*model.Item, error)
}

func NewItemInteractor(r repository.ItemRepository, p presenter.ItemPresenter, d repository.DBRepository) ItemInteractor {
	return &itemInteractor{r,p, d}
}

func (us *itemInteractor) Get(u []*model.Item) ([]*model.Item, error) {
	u, err := us.ItemRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return us.ItemPresenter.ResponseItems(u), nil
}

func (us *itemInteractor) Create(u *model.Item) (*model.Item, error) {
	data, err := us.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := us.ItemRepository.Create(u)
		return u, err
	})
	item, ok := data.(*model.Item)
	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return item, nil
}