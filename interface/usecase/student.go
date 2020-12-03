package usecase

import (
	"golang-bootcamp-2020/domain/model"
)

// StudentService interface
type StudentService interface {
	GetStudentsService() ([]model.Student, error)
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
