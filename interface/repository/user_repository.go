package repository

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

type userRepository struct {
	db     *services.Database
	logger *services.Logger
}

//NewUserRepository creates a new User Repository
func NewUserRepository(db *services.Database, logger *services.Logger) repository.UserRepository {
	return &userRepository{db, logger}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {

	return nil, nil
}

func (ur *userRepository) Create(u *model.User) (*model.User, error) {
	//u, err := ur.db.Create(u)

	return nil, nil
}
