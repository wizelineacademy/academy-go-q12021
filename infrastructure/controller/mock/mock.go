// Source: infrastructure/controller/student.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	gomock "github.com/golang/mock/gomock"
	model "golang-bootcamp-2020/domain/model"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// ReadStudentsService mocks base method
func (m *MockUsecase) ReadStudentsService(filePath string) ([]model.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadStudentsService", filePath)
	ret0, _ := ret[0].([]model.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadStudentsService indicates an expected call of ReadStudentsService
func (mr *MockUsecaseMockRecorder) ReadStudentsService(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadStudentsService", reflect.TypeOf((*MockUsecase)(nil).ReadStudentsService), filePath)
}

// StoreURLService mocks base method
func (m *MockUsecase) StoreURLService(apiURL string) ([]model.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreURLService", apiURL)
	ret0, _ := ret[0].([]model.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreURLService indicates an expected call of StoreURLService
func (mr *MockUsecaseMockRecorder) StoreURLService(apiURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreURLService", reflect.TypeOf((*MockUsecase)(nil).StoreURLService), apiURL)
}

