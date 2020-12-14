package interactor

import (
	"log"
	"time"

	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/presenter"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

type adminInteractor struct {
	AdminRepository repository.AdminRepository
	AdminPresenter  presenter.AdminPresenter
}

type AdminInteractor interface {
	GetLogs(searchPattern string, startDate time.Time, endDate time.Time) ([]string, error)
}

//NewAdminInteractor generates a new instance of an admin interactor
func NewAdminInteractor(r repository.AdminRepository, p presenter.AdminPresenter) AdminInteractor {
	return &adminInteractor{r, p}
}

func (ai *adminInteractor) GetLogs(searchPattern string, startDate time.Time, endDate time.Time) ([]string, error) {
	log.Println("GetLogs interactor")
	//TODO: Implement checking the start/end date filtering
	logs, err := ai.AdminRepository.FindBy(searchPattern, startDate, endDate)
	if err != nil {
		return nil, err
	}
	log.Println("Retrieved: ", logs)

	return logs, nil
}
