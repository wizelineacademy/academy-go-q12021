package usecase

import (
	"golang-bootcamp-2020/domain/model"
)

// StudentController interface
type StudentController interface {
	GetStudentsFromCsv() ([]model.Student, error)
}

// Usecase struct type
type Usecase struct {
	service StudentController
}

// NewUsecase
func NewUsecase(s StudentController) *Usecase {
	return &Usecase{s}
}

// GetStudents usecase
func (u *Usecase) GetStudentsFromCsv() ([]model.Student, error) {
	students, err := u.service.GetStudentsFromCsv()
	return students, err
}
