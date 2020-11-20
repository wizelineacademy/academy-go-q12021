package presenter

import "github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
	ResponseUser(us *model.User) *model.User
}

//NewUserPresenter generates a new instance of UserPresenter
func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr. " + u.Name
	}
	return us
}

func (up *userPresenter) ResponseUser(us *model.User) *model.User {
	return us
}
