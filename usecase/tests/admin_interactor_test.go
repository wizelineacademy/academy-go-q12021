package interactor_test

import (
	"testing"
	"time"

	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/interactor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type adminInteractorMock struct {
	mock.Mock
}
type adminRepositoryMock struct {
	mock.Mock
}
type adminPresenterMock struct {
	mock.Mock
}

func (ar *adminRepositoryMock) FindBy(searchPattern string, startDate time.Time, endDate time.Time) ([]string, error) {
	ar.Called(searchPattern, startDate, endDate)
	return []string{"{this is a date, 1092092}", "{another date, 09320929}"}, nil
}

func (ap *adminPresenterMock) ResponseLogs(logRecords []string) []string {
	return nil
}

func TestGetLogs(t *testing.T) {
	adminRepository := new(adminRepositoryMock)
	adminPresenter := new(adminPresenterMock)
	expectedLogs := []string{"{this is a date, 1092092}", "{another date, 09320929}"}
	currentDate := time.Now()
	adminRepository.On("FindBy", "", currentDate, currentDate).Return(expectedLogs)

	adminInteractor := interactor.NewAdminInteractor(adminRepository, adminPresenter)

	logs, err := adminInteractor.GetLogs("", currentDate, currentDate)
	assert.NoError(t, err, "No error")
	assert.Equal(t, expectedLogs, logs)
}
