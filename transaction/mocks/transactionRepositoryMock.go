// Code generated by MockGen. DO NOT EDIT.
// Source: qwik.in/transaction/domain/repository (interfaces: TransactionRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	app_erros "qwik.in/transaction/app-erros"
	models "qwik.in/transaction/domain/models"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// AddTransactionPointsFromDB mocks base method.
func (m *MockTransactionRepository) AddTransactionPointsFromDB(arg0 *models.Transaction) *app_erros.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransactionPointsFromDB", arg0)
	ret0, _ := ret[0].(*app_erros.AppError)
	return ret0
}

// AddTransactionPointsFromDB indicates an expected call of AddTransactionPointsFromDB.
func (mr *MockTransactionRepositoryMockRecorder) AddTransactionPointsFromDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransactionPointsFromDB", reflect.TypeOf((*MockTransactionRepository)(nil).AddTransactionPointsFromDB), arg0)
}

// DBHealthCheck mocks base method.
func (m *MockTransactionRepository) DBHealthCheck() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBHealthCheck")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DBHealthCheck indicates an expected call of DBHealthCheck.
func (mr *MockTransactionRepositoryMockRecorder) DBHealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBHealthCheck", reflect.TypeOf((*MockTransactionRepository)(nil).DBHealthCheck))
}

// GetTransactionPointsByUserIdFromDB mocks base method.
func (m *MockTransactionRepository) GetTransactionPointsByUserIdFromDB(arg0 string) (int, *app_erros.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionPointsByUserIdFromDB", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*app_erros.AppError)
	return ret0, ret1
}

// GetTransactionPointsByUserIdFromDB indicates an expected call of GetTransactionPointsByUserIdFromDB.
func (mr *MockTransactionRepositoryMockRecorder) GetTransactionPointsByUserIdFromDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionPointsByUserIdFromDB", reflect.TypeOf((*MockTransactionRepository)(nil).GetTransactionPointsByUserIdFromDB), arg0)
}

// UpdateTransactionPointsToDB mocks base method.
func (m *MockTransactionRepository) UpdateTransactionPointsToDB(arg0 *models.Transaction) *app_erros.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionPointsToDB", arg0)
	ret0, _ := ret[0].(*app_erros.AppError)
	return ret0
}

// UpdateTransactionPointsToDB indicates an expected call of UpdateTransactionPointsToDB.
func (mr *MockTransactionRepositoryMockRecorder) UpdateTransactionPointsToDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionPointsToDB", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateTransactionPointsToDB), arg0)
}