// Code generated by MockGen. DO NOT EDIT.
// Source: timetables\timetables.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	timetables "github.com/team-gleam/kiwi-basket/server/src/domain/model/timetables"
	username "github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
)

// MockITimetablesRepository is a mock of ITimetablesRepository interface.
type MockITimetablesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITimetablesRepositoryMockRecorder
}

// MockITimetablesRepositoryMockRecorder is the mock recorder for MockITimetablesRepository.
type MockITimetablesRepositoryMockRecorder struct {
	mock *MockITimetablesRepository
}

// NewMockITimetablesRepository creates a new mock instance.
func NewMockITimetablesRepository(ctrl *gomock.Controller) *MockITimetablesRepository {
	mock := &MockITimetablesRepository{ctrl: ctrl}
	mock.recorder = &MockITimetablesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITimetablesRepository) EXPECT() *MockITimetablesRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockITimetablesRepository) Create(arg0 username.Username, arg1 timetables.Timetables) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockITimetablesRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITimetablesRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockITimetablesRepository) Delete(arg0 username.Username) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITimetablesRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITimetablesRepository)(nil).Delete), arg0)
}

// Exists mocks base method.
func (m *MockITimetablesRepository) Exists(arg0 username.Username) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockITimetablesRepositoryMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockITimetablesRepository)(nil).Exists), arg0)
}

// Get mocks base method.
func (m *MockITimetablesRepository) Get(arg0 username.Username) (timetables.Timetables, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(timetables.Timetables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockITimetablesRepositoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockITimetablesRepository)(nil).Get), arg0)
}
