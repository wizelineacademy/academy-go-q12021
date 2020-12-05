package usecase

import (
	"fmt"
	"log"

	"golang-bootcamp-2020/domain/model"
)

// StudentService interface
type StudentService interface {
	GetStudentsService() ([]model.Student, error)
	GetUrlService() ([]model.Student, error)
	SaveToCsv(students []model.Student) (bool, error)
}

// Usecase struct type
type Usecase struct {
	service StudentService
}

// NewUsecase
func NewUsecase(s StudentService) *Usecase {
	return &Usecase{s}
}

// GetStudentsHandler usecase
func (u *Usecase) GetStudentsService() ([]model.Student, error) {
	students, err := u.service.GetStudentsService()
	return students, err
}

func (u *Usecase) GetUrlService() ([]model.Student, error) {
	students, err := u.service.GetUrlService()
	if err != nil {
		log.Fatal("Fallo leer url ", err)
	}
	res, err := u.service.SaveToCsv(students)
	if err != nil {
		log.Fatal("Fallo el salvar csv ", err)
	}
	if res {
		fmt.Println("se guardo con exito")
	}

	return students, err
}
