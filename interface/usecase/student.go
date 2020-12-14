// Usecase package
package usecase

import (
	"fmt"
	"path/filepath"

	"github.com/ruvaz/golang-bootcamp-2020/config"
	"github.com/ruvaz/golang-bootcamp-2020/domain/model"
)

// StudentService interface
type StudentService interface {
	ReadStudentsService(filePath string) ([]model.Student, error)
	StoreURLService(apiURL string) ([]model.Student, error)
	SaveToCsv(students []model.Student, filePath string) (bool, error)
}

// Usecase struct
type Usecase struct {
	service StudentService
}

// NewUsecase using student service interface
func NewUsecase(s StudentService) *Usecase {
	return &Usecase{s}
}

// ReadStudentsService: usecase to read students from csv
func (u *Usecase) ReadStudentsService(filePath string) ([]model.Student, error) {
	students, err := u.service.ReadStudentsService(filePath)
	return students, err
}

// StoreURLService usecase to store students from api to csv
func (u *Usecase) StoreURLService(apiURL string) ([]model.Student, error) {
	config.ReadConfig()

	// get data from api into []students
	students, err := u.service.StoreURLService(apiURL)
	if err != nil {
		return students, fmt.Errorf("the URL could not be obtained")
	}
	filePath, err := filepath.Abs(config.C.CsvPath.Prod)
	if err != nil {
		return students, fmt.Errorf("failed to get the file")
	}

	// Save students in csv file
	res, err := u.service.SaveToCsv(students, filePath)
	if err != nil || !res {
		return students, fmt.Errorf("failed to save csv")
	}
	return students, nil
}
