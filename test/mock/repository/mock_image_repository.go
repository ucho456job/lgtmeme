// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/image_repository.go

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	dto "github.com/ucho456job/lgtmeme/internal/dto"
	model "github.com/ucho456job/lgtmeme/internal/model"
)

// MockImageRepository is a mock of ImageRepository interface.
type MockImageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockImageRepositoryMockRecorder
}

// MockImageRepositoryMockRecorder is the mock recorder for MockImageRepository.
type MockImageRepositoryMockRecorder struct {
	mock *MockImageRepository
}

// NewMockImageRepository creates a new mock instance.
func NewMockImageRepository(ctrl *gomock.Controller) *MockImageRepository {
	mock := &MockImageRepository{ctrl: ctrl}
	mock.recorder = &MockImageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageRepository) EXPECT() *MockImageRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockImageRepository) Create(c echo.Context, id uuid.UUID, url, keyword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, id, url, keyword)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockImageRepositoryMockRecorder) Create(c, id, url, keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockImageRepository)(nil).Create), c, id, url, keyword)
}

// Delete mocks base method.
func (m *MockImageRepository) Delete(c echo.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockImageRepositoryMockRecorder) Delete(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockImageRepository)(nil).Delete), c, id)
}

// ExistsByID mocks base method.
func (m *MockImageRepository) ExistsByID(c echo.Context, id uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByID", c, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByID indicates an expected call of ExistsByID.
func (mr *MockImageRepositoryMockRecorder) ExistsByID(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByID", reflect.TypeOf((*MockImageRepository)(nil).ExistsByID), c, id)
}

// FindByGetImagesQuery mocks base method.
func (m *MockImageRepository) FindByGetImagesQuery(c echo.Context, q dto.GetImagesQuery) (*[]model.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGetImagesQuery", c, q)
	ret0, _ := ret[0].(*[]model.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGetImagesQuery indicates an expected call of FindByGetImagesQuery.
func (mr *MockImageRepositoryMockRecorder) FindByGetImagesQuery(c, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGetImagesQuery", reflect.TypeOf((*MockImageRepository)(nil).FindByGetImagesQuery), c, q)
}

// FirstByID mocks base method.
func (m *MockImageRepository) FirstByID(c echo.Context, id uuid.UUID, columns []string) (*model.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstByID", c, id, columns)
	ret0, _ := ret[0].(*model.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstByID indicates an expected call of FirstByID.
func (mr *MockImageRepositoryMockRecorder) FirstByID(c, id, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstByID", reflect.TypeOf((*MockImageRepository)(nil).FirstByID), c, id, columns)
}

// Update mocks base method.
func (m *MockImageRepository) Update(c echo.Context, id uuid.UUID, reqType dto.PatchImageReqType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, id, reqType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockImageRepositoryMockRecorder) Update(c, id, reqType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockImageRepository)(nil).Update), c, id, reqType)
}