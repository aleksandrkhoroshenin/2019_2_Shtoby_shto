// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package user is a generated GoMock package.
package user

import (
	. "2019_2_Shtoby_shto/src/customType"
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockHandlerUserService is a mock of HandlerUserService interface
type MockHandlerUserService struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerUserServiceMockRecorder
}

// MockHandlerUserServiceMockRecorder is the mock recorder for MockHandlerUserService
type MockHandlerUserServiceMockRecorder struct {
	mock *MockHandlerUserService
}

// NewMockHandlerUserService creates a new mock instance
func NewMockHandlerUserService(ctrl *gomock.Controller) *MockHandlerUserService {
	mock := &MockHandlerUserService{ctrl: ctrl}
	mock.recorder = &MockHandlerUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandlerUserService) EXPECT() *MockHandlerUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockHandlerUserService) CreateUser(data []byte) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", data)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockHandlerUserServiceMockRecorder) CreateUser(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockHandlerUserService)(nil).CreateUser), data)
}

// UpdateUser mocks base method
func (m *MockHandlerUserService) UpdateUser(data []byte, id StringUUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", data, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockHandlerUserServiceMockRecorder) UpdateUser(data, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockHandlerUserService)(nil).UpdateUser), data, id)
}

// GetUserById mocks base method
func (m *MockHandlerUserService) GetUserById(id StringUUID) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockHandlerUserServiceMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockHandlerUserService)(nil).GetUserById), id)
}

// GetUserByLogin mocks base method
func (m *MockHandlerUserService) GetUserByLogin(data []byte) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLogin", data)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByLogin indicates an expected call of GetUserByLogin
func (mr *MockHandlerUserServiceMockRecorder) GetUserByLogin(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLogin", reflect.TypeOf((*MockHandlerUserService)(nil).GetUserByLogin), data)
}

// FetchUsers mocks base method
func (m *MockHandlerUserService) FetchUsers(limit, offset int) ([]User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUsers", limit, offset)
	ret0, _ := ret[0].([]User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUsers indicates an expected call of FetchUsers
func (mr *MockHandlerUserServiceMockRecorder) FetchUsers(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUsers", reflect.TypeOf((*MockHandlerUserService)(nil).FetchUsers), limit, offset)
}
