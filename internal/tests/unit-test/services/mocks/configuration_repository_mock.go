// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/configuration_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/manabie-com/togo/internal/models"
)

// MockConfigurationRepository is a mock of ConfigurationRepository interface.
type MockConfigurationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationRepositoryMockRecorder
}

// MockConfigurationRepositoryMockRecorder is the mock recorder for MockConfigurationRepository.
type MockConfigurationRepositoryMockRecorder struct {
	mock *MockConfigurationRepository
}

// NewMockConfigurationRepository creates a new mock instance.
func NewMockConfigurationRepository(ctrl *gomock.Controller) *MockConfigurationRepository {
	mock := &MockConfigurationRepository{ctrl: ctrl}
	mock.recorder = &MockConfigurationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationRepository) EXPECT() *MockConfigurationRepositoryMockRecorder {
	return m.recorder
}

// CreateConfiguration mocks base method.
func (m *MockConfigurationRepository) CreateConfiguration(ctx context.Context, configuration *models.Configuration) (*models.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConfiguration", ctx, configuration)
	ret0, _ := ret[0].(*models.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConfiguration indicates an expected call of CreateConfiguration.
func (mr *MockConfigurationRepositoryMockRecorder) CreateConfiguration(ctx, configuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfiguration", reflect.TypeOf((*MockConfigurationRepository)(nil).CreateConfiguration), ctx, configuration)
}

// GetAll mocks base method.
func (m *MockConfigurationRepository) GetAll(ctx context.Context) ([]*models.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*models.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockConfigurationRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockConfigurationRepository)(nil).GetAll), ctx)
}

// GetById mocks base method.
func (m *MockConfigurationRepository) GetById(ctx context.Context, configurationId string) (*models.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, configurationId)
	ret0, _ := ret[0].(*models.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockConfigurationRepositoryMockRecorder) GetById(ctx, configurationId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockConfigurationRepository)(nil).GetById), ctx, configurationId)
}

// GetCapacity mocks base method.
func (m *MockConfigurationRepository) GetCapacity(ctx context.Context, userID, date string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapacity", ctx, userID, date)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCapacity indicates an expected call of GetCapacity.
func (mr *MockConfigurationRepositoryMockRecorder) GetCapacity(ctx, userID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapacity", reflect.TypeOf((*MockConfigurationRepository)(nil).GetCapacity), ctx, userID, date)
}