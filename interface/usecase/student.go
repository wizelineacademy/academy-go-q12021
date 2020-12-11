/**
Student usecase
*/
package usecase

import (
	"fmt"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/domain/model"
)

// StudentService interface
type StudentService interface {
	ReadStudentsService(filePath string) ([]model.Student, error)
	StoreURLService(apiURL string) ([]model.Student, error)
	SaveToCsv(students []model.Student,  filePath string) (bool, error)
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
func (u *Usecase) ReadStudentsService(filePath string) ([]model.Student, error) {
	students, err := u.service.ReadStudentsService(filePath)
	return students, err
}

// StoreURLService in usecase
func (u *Usecase) StoreURLService(apiURL string) ([]model.Student, error) {
	// get data from api into []students
	students, err := u.service.StoreURLService(apiURL)
	if err != nil {
		return students, fmt.Errorf("the URL could not be obtained")
	}
	filePath :=config.C.CsvPath.Prod
	// Save students in csv file
	res, err := u.service.SaveToCsv(students, filePath)
	if err != nil || !res {
		return students, fmt.Errorf("failed to save csv")
	}
	return students, nil
}
