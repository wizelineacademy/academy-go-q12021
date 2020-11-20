package interactor

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/presenter"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
	Create(user *model.User) (*model.User, error)
}

//NewUserInteractor creates a new user iteractor
func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(user []*model.User) ([]*model.User, error) {
	user, err := us.UserRepository.FindAll(user)

	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(user), nil
}

func (us *userInteractor) Create(user *model.User) (*model.User, error) {
	u, err := us.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return us.UserPresenter.ResponseUser(u), nil
}
