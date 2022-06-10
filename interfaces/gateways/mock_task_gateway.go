// Code generated by MockGen. DO NOT EDIT.
// Source: task_gateway.go

// Package mock_gateways is a generated GoMock package.
package gateways

import (
	entities "api-bff-golang/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskGatewayInterface is a mock of TaskGatewayInterface interface.
type MockTaskGatewayInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTaskGatewayInterfaceMockRecorder
}

// MockTaskGatewayInterfaceMockRecorder is the mock recorder for MockTaskGatewayInterface.
type MockTaskGatewayInterfaceMockRecorder struct {
	mock *MockTaskGatewayInterface
}

// NewMockTaskGatewayInterface creates a new mock instance.
func NewMockTaskGatewayInterface(ctrl *gomock.Controller) *MockTaskGatewayInterface {
	mock := &MockTaskGatewayInterface{ctrl: ctrl}
	mock.recorder = &MockTaskGatewayInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskGatewayInterface) EXPECT() *MockTaskGatewayInterfaceMockRecorder {
	return m.recorder
}

// DeleteByTitle mocks base method.
func (m *MockTaskGatewayInterface) DeleteByTitle(title string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByTitle", title)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByTitle indicates an expected call of DeleteByTitle.
func (mr *MockTaskGatewayInterfaceMockRecorder) DeleteByTitle(title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByTitle", reflect.TypeOf((*MockTaskGatewayInterface)(nil).DeleteByTitle), title)
}

// FindAll mocks base method.
func (m *MockTaskGatewayInterface) FindAll() ([]entities.TaskEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]entities.TaskEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockTaskGatewayInterfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTaskGatewayInterface)(nil).FindAll))
}

// GetByTitle mocks base method.
func (m *MockTaskGatewayInterface) GetByTitle(title string) (entities.TaskEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTitle", title)
	ret0, _ := ret[0].(entities.TaskEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTitle indicates an expected call of GetByTitle.
func (mr *MockTaskGatewayInterfaceMockRecorder) GetByTitle(title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTitle", reflect.TypeOf((*MockTaskGatewayInterface)(nil).GetByTitle), title)
}

// SaveTask mocks base method.
func (m *MockTaskGatewayInterface) SaveTask(task *entities.TaskEntity) (entities.TaskEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTask", task)
	ret0, _ := ret[0].(entities.TaskEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveTask indicates an expected call of SaveTask.
func (mr *MockTaskGatewayInterfaceMockRecorder) SaveTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTask", reflect.TypeOf((*MockTaskGatewayInterface)(nil).SaveTask), task)
}