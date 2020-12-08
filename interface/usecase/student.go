/**
Student usecase
*/
package usecase

import (
	"fmt"

	"golang-bootcamp-2020/domain/model"
)

// StudentService interface
type StudentService interface {
	ReadStudentsService() ([]model.Student, error)
	StoreURLService() ([]model.Student, error)
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

// ReadStudentsService in usecase
func (u *Usecase) ReadStudentsService() ([]model.Student, error) {
	students, err := u.service.ReadStudentsService()
	return students, err
}

// StoreURLService in usecase
func (u *Usecase) StoreURLService() ([]model.Student, error) {
	// get data from api into []students
	students, err := u.service.StoreURLService()
	if err != nil {
		return students, fmt.Errorf("the URL could not be obtained")
	}
	// Save students in csv file
	res, err := u.service.SaveToCsv(students)
	if err != nil || !res {
		return students, fmt.Errorf("failed to save csv")
	}
	return students, nil
}
