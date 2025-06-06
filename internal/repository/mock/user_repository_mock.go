// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/user_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	model "task-tracker/internal/model"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user)
}

// Delete mocks base method.
func (m *MockUserRepository) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepository)(nil).Delete), id)
}

// FindAll mocks base method.
func (m *MockUserRepository) FindAll() ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserRepository)(nil).FindAll))
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), email)
}

// FindByEmailWithRole mocks base method.
func (m *MockUserRepository) FindByEmailWithRole(email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmailWithRole", email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmailWithRole indicates an expected call of FindByEmailWithRole.
func (mr *MockUserRepositoryMockRecorder) FindByEmailWithRole(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmailWithRole", reflect.TypeOf((*MockUserRepository)(nil).FindByEmailWithRole), email)
}

// FindByID mocks base method.
func (m *MockUserRepository) FindByID(id uint) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepository)(nil).FindByID), id)
}

// FindByTokenWithRole mocks base method.
func (m *MockUserRepository) FindByTokenWithRole(token string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTokenWithRole", token)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTokenWithRole indicates an expected call of FindByTokenWithRole.
func (mr *MockUserRepositoryMockRecorder) FindByTokenWithRole(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTokenWithRole", reflect.TypeOf((*MockUserRepository)(nil).FindByTokenWithRole), token)
}

// Update mocks base method.
func (m *MockUserRepository) Update(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), user)
}
