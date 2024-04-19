// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/scope_repository.go

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
	model "github.com/ucho456job/lgtmeme/internal/model"
)

// MockScopeRepository is a mock of ScopeRepository interface.
type MockScopeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockScopeRepositoryMockRecorder
}

// MockScopeRepositoryMockRecorder is the mock recorder for MockScopeRepository.
type MockScopeRepositoryMockRecorder struct {
	mock *MockScopeRepository
}

// NewMockScopeRepository creates a new mock instance.
func NewMockScopeRepository(ctrl *gomock.Controller) *MockScopeRepository {
	mock := &MockScopeRepository{ctrl: ctrl}
	mock.recorder = &MockScopeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScopeRepository) EXPECT() *MockScopeRepositoryMockRecorder {
	return m.recorder
}

// FindByScopesStr mocks base method.
func (m *MockScopeRepository) FindByScopesStr(c echo.Context, scopesStr string) (*[]model.MasterScope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByScopesStr", c, scopesStr)
	ret0, _ := ret[0].(*[]model.MasterScope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByScopesStr indicates an expected call of FindByScopesStr.
func (mr *MockScopeRepositoryMockRecorder) FindByScopesStr(c, scopesStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByScopesStr", reflect.TypeOf((*MockScopeRepository)(nil).FindByScopesStr), c, scopesStr)
}
