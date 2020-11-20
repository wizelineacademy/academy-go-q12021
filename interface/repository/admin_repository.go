package repository

import (
	"time"

	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

type adminRepository struct {
	logger services.Logger
}

func NewAdminRepository(logger services.Logger) repository.AdminRepository {
	return &adminRepository{logger: logger}
}

func (ar *adminRepository) FindBy(searchPattern string, startDate, endDate time.Time) ([]string, error) {
	ar.logger.Get()
	return nil, nil
}
