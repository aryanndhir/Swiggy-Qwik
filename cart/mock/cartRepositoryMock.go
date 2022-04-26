// Code generated by MockGen. DO NOT EDIT.
// Source: cartService/domain/repository (interfaces: CartRepositoryDB)

// Package mocks is a generated GoMock package.
package mocks

import (
	model "cartService/domain/model"
	error "cartService/internal/error"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCartRepositoryDB is a mock of CartRepositoryDB interface
type MockCartRepositoryDB struct {
	ctrl     *gomock.Controller
	recorder *MockCartRepositoryDBMockRecorder
}

// MockCartRepositoryDBMockRecorder is the mock recorder for MockCartRepositoryDB
type MockCartRepositoryDBMockRecorder struct {
	mock *MockCartRepositoryDB
}

// NewMockCartRepositoryDB creates a new mock instance
func NewMockCartRepositoryDB(ctrl *gomock.Controller) *MockCartRepositoryDB {
	mock := &MockCartRepositoryDB{ctrl: ctrl}
	mock.recorder = &MockCartRepositoryDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCartRepositoryDB) EXPECT() *MockCartRepositoryDBMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCartRepositoryDB) Create(arg0 model.Cart) *error.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*error.AppError)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCartRepositoryDBMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCartRepositoryDB)(nil).Create), arg0)
}

// DBHealthCheck mocks base method
func (m *MockCartRepositoryDB) DBHealthCheck() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBHealthCheck")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DBHealthCheck indicates an expected call of DBHealthCheck
func (mr *MockCartRepositoryDBMockRecorder) DBHealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBHealthCheck", reflect.TypeOf((*MockCartRepositoryDB)(nil).DBHealthCheck))
}

// Delete mocks base method
func (m *MockCartRepositoryDB) Delete(arg0 string) *error.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*error.AppError)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCartRepositoryDBMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCartRepositoryDB)(nil).Delete), arg0)
}

// Read mocks base method
func (m *MockCartRepositoryDB) Read(arg0 string) (*model.Cart, *error.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(*model.Cart)
	ret1, _ := ret[1].(*error.AppError)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockCartRepositoryDBMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCartRepositoryDB)(nil).Read), arg0)
}

// ReadAll mocks base method
func (m *MockCartRepositoryDB) ReadAll() (*[]model.Cart, *error.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll")
	ret0, _ := ret[0].(*[]model.Cart)
	ret1, _ := ret[1].(*error.AppError)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll
func (mr *MockCartRepositoryDBMockRecorder) ReadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockCartRepositoryDB)(nil).ReadAll))
}

// UpdateExisting mocks base method
func (m *MockCartRepositoryDB) UpdateExisting(arg0 *model.Cart) *error.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExisting", arg0)
	ret0, _ := ret[0].(*error.AppError)
	return ret0
}

// UpdateExisting indicates an expected call of UpdateExisting
func (mr *MockCartRepositoryDBMockRecorder) UpdateExisting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExisting", reflect.TypeOf((*MockCartRepositoryDB)(nil).UpdateExisting), arg0)
}
