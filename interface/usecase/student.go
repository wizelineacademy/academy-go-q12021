// Student usecase
package usecase

import (
	"fmt"

	"golang-bootcamp-2020/domain/model"
)

// StudentService interface
type StudentService interface {
	GetStudentsService() ([]model.Student, error)
	GetURLService() ([]model.Student, error)
	SaveToCsv(students []model.Student) (bool, error)
}

// Usecase struct
type Usecase struct {
	service StudentService
}

// NewUsecase using student service interface
func NewUsecase(s StudentService) *Usecase {
	return &Usecase{s}
}

// GetStudentsService in usecase
func (u *Usecase) GetStudentsService() ([]model.Student, error) {
	students, err := u.service.GetStudentsService()
	return students, err
}

// GetURLService in usecase
func (u *Usecase) GetURLService() ([]model.Student, error) {
	// get data from api into []students
	students, err := u.service.GetURLService()
	if err != nil {
		return students, fmt.Errorf("The URL could not be obtained")
	}
	// Save students in csv file
	res, err2 := u.service.SaveToCsv(students)
	if err2 != nil || res != true {
		return students, fmt.Errorf("Failed to save csv")
	}
	return students, err
}
