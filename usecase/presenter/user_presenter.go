package presenter

import "github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
	ResponseUser(u *model.User) *model.User
}
