package datastore

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
)

type MySQL struct { //Implements Service, Database
	db     int //TODO: Implement MySQL DB
	status int
}

var users = []*model.User{ //TODO: Remove this when connected to real DB
	{ID: "1",
		Password: "dpaoidjkpaosijda",
	},
	{
		ID:       "2",
		Password: "dpaoidjkpaosijda",
	},
}

func InitializeDB() *MySQL {
	//TODO: Implement connecting to DB
	return &MySQL{}
}

func (mysql *MySQL) Start() error {
	return nil
}
func (mysql *MySQL) Stop() error {
	return nil
}

func (mysql *MySQL) Status() int {
	return mysql.status
}

func (mysql *MySQL) Get(*model.User) (*model.User, error) {
	return users[0], nil
}
func (mysql *MySQL) Create(*model.User) (*model.User, error) {
	user := &model.User{
		ID:       "3",
		Password: "453ea",
	}
	users = append(users, user)
	return user, nil
}
func (mysql *MySQL) Update(*model.User) (*model.User, error) {
	return nil, nil
}
func (mysql *MySQL) Delete() error {
	return nil
}
