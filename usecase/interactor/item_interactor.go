package interactor

import (
	"bootcamp/domain/model"
	"bootcamp/interface/controller/vo"
	"bootcamp/usecase/presenter"
	"bootcamp/usecase/repository"
	"errors"
)

// itemInteractor struct for Items model
type itemInteractor struct {
	ItemRepository repository.ItemRepository
	ItemPresenter presenter.ItemPresenter
	DBRepository repository.DBRepository
}

// ItemInteractor
type ItemInteractor interface {
	// Get return all Items
	GetItems(i []*model.Item, paged *vo.Paged) ([]*model.Item, error)

	// Create func creates a new Items
	Create(i *model.Item) (*model.Item, error)
}

// NewItemInteractor returns a new interactor struct
func NewItemInteractor(r repository.ItemRepository, p presenter.ItemPresenter, d repository.DBRepository) ItemInteractor {
	return &itemInteractor{r,p, d}
}

// Get func returns all Items.
func (ii *itemInteractor) GetItems(items []*model.Item, paged *vo.Paged) ([]*model.Item, error) {

	var err error = nil
	if paged == nil {
		items, err = ii.ItemRepository.FindAll(items)
	} else {
		items, err = ii.ItemRepository.FindAllPaged(items, paged)
	}

	if err != nil {
		return nil, err
	}
	return ii.ItemPresenter.ResponseItems(items), nil
}

// Create func add a new Item in the datastore.
func (ii *itemInteractor) Create(item *model.Item) (*model.Item, error) {
	data, err := ii.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := ii.ItemRepository.Create(item)
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