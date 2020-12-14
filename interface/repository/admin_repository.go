package repository

import (
	"log"
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
	log.Println("FindBy")
	records, err := ar.logger.Get()
	if err != nil {
		return nil, err
	}
	return records, nil
}
