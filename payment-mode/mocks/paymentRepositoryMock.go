// Code generated by MockGen. DO NOT EDIT.
// Source: qwik.in/payment-mode/domain/repository (interfaces: PaymentRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	app_errors "qwik.in/payment-mode/app-errors"
	models "qwik.in/payment-mode/domain/models"
)

// MockPaymentRepository is a mock of PaymentRepository interface.
type MockPaymentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentRepositoryMockRecorder
}

// MockPaymentRepositoryMockRecorder is the mock recorder for MockPaymentRepository.
type MockPaymentRepositoryMockRecorder struct {
	mock *MockPaymentRepository
}

// NewMockPaymentRepository creates a new mock instance.
func NewMockPaymentRepository(ctrl *gomock.Controller) *MockPaymentRepository {
	mock := &MockPaymentRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentRepository) EXPECT() *MockPaymentRepositoryMockRecorder {
	return m.recorder
}

// AddPaymentModeToDB mocks base method.
func (m *MockPaymentRepository) AddPaymentModeToDB(arg0 *models.UserPaymentMode) *app_errors.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPaymentModeToDB", arg0)
	ret0, _ := ret[0].(*app_errors.AppError)
	return ret0
}

// AddPaymentModeToDB indicates an expected call of AddPaymentModeToDB.
func (mr *MockPaymentRepositoryMockRecorder) AddPaymentModeToDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPaymentModeToDB", reflect.TypeOf((*MockPaymentRepository)(nil).AddPaymentModeToDB), arg0)
}

// DBHealthCheck mocks base method.
func (m *MockPaymentRepository) DBHealthCheck() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBHealthCheck")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DBHealthCheck indicates an expected call of DBHealthCheck.
func (mr *MockPaymentRepositoryMockRecorder) DBHealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBHealthCheck", reflect.TypeOf((*MockPaymentRepository)(nil).DBHealthCheck))
}

// GetPaymentModeFromDB mocks base method.
func (m *MockPaymentRepository) GetPaymentModeFromDB(arg0 string) (*models.UserPaymentMode, *app_errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentModeFromDB", arg0)
	ret0, _ := ret[0].(*models.UserPaymentMode)
	ret1, _ := ret[1].(*app_errors.AppError)
	return ret0, ret1
}

// GetPaymentModeFromDB indicates an expected call of GetPaymentModeFromDB.
func (mr *MockPaymentRepositoryMockRecorder) GetPaymentModeFromDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentModeFromDB", reflect.TypeOf((*MockPaymentRepository)(nil).GetPaymentModeFromDB), arg0)
}

// UpdatePaymentModeToDB mocks base method.
func (m *MockPaymentRepository) UpdatePaymentModeToDB(arg0 *models.UserPaymentMode) *app_errors.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePaymentModeToDB", arg0)
	ret0, _ := ret[0].(*app_errors.AppError)
	return ret0
}

// UpdatePaymentModeToDB indicates an expected call of UpdatePaymentModeToDB.
func (mr *MockPaymentRepositoryMockRecorder) UpdatePaymentModeToDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePaymentModeToDB", reflect.TypeOf((*MockPaymentRepository)(nil).UpdatePaymentModeToDB), arg0)
}
