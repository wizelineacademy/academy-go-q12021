// usecase test
package usecase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"golang-bootcamp-2020/infrastructure/services"
)
type MyMockedObject struct{
	mock.Mock
}

func TestUsecase_ReadStudentsService1(t *testing.T) {
	s := services.NewClient()
	u := NewUsecase(s)
	stu,err:= u.ReadStudentsService("../../dataFile.csv")
	if err != nil {
		t.Logf("fail readstudents services")
	}
	fmt.Println(stu)
}

func TestUsecase_StoreURLService(t *testing.T) {
	s := services.NewClient()
	u := NewUsecase(s)
	apiURL:= "https://login-app-crud.firebaseio.com/.json"
	_,err:= u.StoreURLService(apiURL)
	if err != nil {
		t.Logf("fail get students url")
	}
}

func TestUsecase_FailStoreURLService(t *testing.T) {
	s := services.NewClient()
	u := NewUsecase(s)
	apiURL:= "https://login-app-crud.firebaseio.com"
	_,err:= u.StoreURLService(apiURL)
	if err != nil {
		t.Logf("fail get students url")
	}
}