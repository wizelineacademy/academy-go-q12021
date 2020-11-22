package repository

import (
	"fmt"
	"log"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

type userRepository struct {
	db     services.Database
	logger services.Logger
}

//NewUserRepository creates a new User Repository
func NewUserRepository(db services.Database, logger services.Logger) repository.UserRepository {
	return &userRepository{db, logger}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	users := []*model.User{}
	for i := 0; i < 2; i++ {
		us, _ := ur.db.Get(nil)
		users = append(users, us)
		log.Println("Estoy aquí")
		ur.logger.Append(fmt.Sprintf("Retrieved %v", us))
	}
	return users, nil
}

func (ur *userRepository) Create(u *model.User) (*model.User, error) {
	u, err := ur.db.Create(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
