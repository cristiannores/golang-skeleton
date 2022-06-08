// Code generated by MockGen. DO NOT EDIT.
// Source: consumer_client.go

// Package mock_consumer is a generated GoMock package.
package consumer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConsumerClientInterface is a mock of ConsumerClientInterface interface.
type MockConsumerClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerClientInterfaceMockRecorder
}

// MockConsumerClientInterfaceMockRecorder is the mock recorder for MockConsumerClientInterface.
type MockConsumerClientInterfaceMockRecorder struct {
	mock *MockConsumerClientInterface
}

// NewMockConsumerClientInterface creates a new mock instance.
func NewMockConsumerClientInterface(ctrl *gomock.Controller) *MockConsumerClientInterface {
	mock := &MockConsumerClientInterface{ctrl: ctrl}
	mock.recorder = &MockConsumerClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerClientInterface) EXPECT() *MockConsumerClientInterfaceMockRecorder {
	return m.recorder
}

// Consumer mocks base method.
func (m *MockConsumerClientInterface) Consumer(ctx context.Context) <-chan IncomingMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consumer", ctx)
	ret0, _ := ret[0].(<-chan IncomingMessage)
	return ret0
}

// Consumer indicates an expected call of Consumer.
func (mr *MockConsumerClientInterfaceMockRecorder) Consumer(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consumer", reflect.TypeOf((*MockConsumerClientInterface)(nil).Consumer), ctx)
}
