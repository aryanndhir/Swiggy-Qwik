// Code generated by MockGen. DO NOT EDIT.
// Source: qwik.in/transaction/domain/services (interfaces: TransactionService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	app_erros "qwik.in/transaction/app-erros"
	models "qwik.in/transaction/domain/models"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// AddTransactionPoints mocks base method.
func (m *MockTransactionService) AddTransactionPoints(arg0 *models.TransactionDetails) *app_erros.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransactionPoints", arg0)
	ret0, _ := ret[0].(*app_erros.AppError)
	return ret0
}

// AddTransactionPoints indicates an expected call of AddTransactionPoints.
func (mr *MockTransactionServiceMockRecorder) AddTransactionPoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransactionPoints", reflect.TypeOf((*MockTransactionService)(nil).AddTransactionPoints), arg0)
}

// CalculateTransactionPoints mocks base method.
func (m *MockTransactionService) CalculateTransactionPoints(arg0 *models.TransactionDetails) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateTransactionPoints", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// CalculateTransactionPoints indicates an expected call of CalculateTransactionPoints.
func (mr *MockTransactionServiceMockRecorder) CalculateTransactionPoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateTransactionPoints", reflect.TypeOf((*MockTransactionService)(nil).CalculateTransactionPoints), arg0)
}

// GetTransactionPointsByUserId mocks base method.
func (m *MockTransactionService) GetTransactionPointsByUserId(arg0 string) (int, *app_erros.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionPointsByUserId", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*app_erros.AppError)
	return ret0, ret1
}

// GetTransactionPointsByUserId indicates an expected call of GetTransactionPointsByUserId.
func (mr *MockTransactionServiceMockRecorder) GetTransactionPointsByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionPointsByUserId", reflect.TypeOf((*MockTransactionService)(nil).GetTransactionPointsByUserId), arg0)
}

// UpdateTransactionPoints mocks base method.
func (m *MockTransactionService) UpdateTransactionPoints(arg0 int, arg1 string) *app_erros.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionPoints", arg0, arg1)
	ret0, _ := ret[0].(*app_erros.AppError)
	return ret0
}

// UpdateTransactionPoints indicates an expected call of UpdateTransactionPoints.
func (mr *MockTransactionServiceMockRecorder) UpdateTransactionPoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionPoints", reflect.TypeOf((*MockTransactionService)(nil).UpdateTransactionPoints), arg0, arg1)
}

// UseTransactionPoints mocks base method.
func (m *MockTransactionService) UseTransactionPoints(arg0 *models.TransactionDetails) (bool, *models.TransactionDetails, *app_erros.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseTransactionPoints", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*models.TransactionDetails)
	ret2, _ := ret[2].(*app_erros.AppError)
	return ret0, ret1, ret2
}

// UseTransactionPoints indicates an expected call of UseTransactionPoints.
func (mr *MockTransactionServiceMockRecorder) UseTransactionPoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseTransactionPoints", reflect.TypeOf((*MockTransactionService)(nil).UseTransactionPoints), arg0)
}