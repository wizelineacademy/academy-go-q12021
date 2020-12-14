// usecase uni test
package usecase

import (
	"testing"

	"golang-bootcamp-2020/infrastructure/services"
)

// TestUsecaseReadStudentsService: test usecase  for read students service
func TestUsecaseReadStudentsService(t *testing.T) {
	s := services.NewClient()
	u := NewUsecase(s)
	_, err := u.ReadStudentsService("../../dataFile.csv")
	if err != nil {
		t.Logf("fail readstudents services")
	}
}

// TestUsecaseStoreURLService: test usecase for StoreURLService
func TestUsecaseStoreURLService(t *testing.T) {
	s := services.NewClient()
	u := NewUsecase(s)
	apiURL := "https://login-app-crud.firebaseio.com/.json"
	_, err := u.StoreURLService(apiURL)
	if err != nil {
		t.Logf("fail get students url")
	}
}

// TestUsecaseFailStoreURLService: test usecase for fail StoreURLService
func TestUsecaseFailStoreURLService(t *testing.T) {
	s := services.NewClient()
	u := NewUsecase(s)
	apiURL := "https://login-app-crud.firebaseio.com"
	_, err := u.StoreURLService(apiURL)
	if err != nil {
		t.Logf("fail get students url")
	}
}
