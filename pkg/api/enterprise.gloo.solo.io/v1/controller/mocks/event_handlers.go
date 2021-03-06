// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockAuthConfigEventHandler is a mock of AuthConfigEventHandler interface
type MockAuthConfigEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigEventHandlerMockRecorder
}

// MockAuthConfigEventHandlerMockRecorder is the mock recorder for MockAuthConfigEventHandler
type MockAuthConfigEventHandlerMockRecorder struct {
	mock *MockAuthConfigEventHandler
}

// NewMockAuthConfigEventHandler creates a new mock instance
func NewMockAuthConfigEventHandler(ctrl *gomock.Controller) *MockAuthConfigEventHandler {
	mock := &MockAuthConfigEventHandler{ctrl: ctrl}
	mock.recorder = &MockAuthConfigEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthConfigEventHandler) EXPECT() *MockAuthConfigEventHandlerMockRecorder {
	return m.recorder
}

// CreateAuthConfig mocks base method
func (m *MockAuthConfigEventHandler) CreateAuthConfig(obj *v1.AuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAuthConfig indicates an expected call of CreateAuthConfig
func (mr *MockAuthConfigEventHandlerMockRecorder) CreateAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthConfig", reflect.TypeOf((*MockAuthConfigEventHandler)(nil).CreateAuthConfig), obj)
}

// UpdateAuthConfig mocks base method
func (m *MockAuthConfigEventHandler) UpdateAuthConfig(old, new *v1.AuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthConfig", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthConfig indicates an expected call of UpdateAuthConfig
func (mr *MockAuthConfigEventHandlerMockRecorder) UpdateAuthConfig(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthConfig", reflect.TypeOf((*MockAuthConfigEventHandler)(nil).UpdateAuthConfig), old, new)
}

// DeleteAuthConfig mocks base method
func (m *MockAuthConfigEventHandler) DeleteAuthConfig(obj *v1.AuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthConfig indicates an expected call of DeleteAuthConfig
func (mr *MockAuthConfigEventHandlerMockRecorder) DeleteAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthConfig", reflect.TypeOf((*MockAuthConfigEventHandler)(nil).DeleteAuthConfig), obj)
}

// GenericAuthConfig mocks base method
func (m *MockAuthConfigEventHandler) GenericAuthConfig(obj *v1.AuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericAuthConfig indicates an expected call of GenericAuthConfig
func (mr *MockAuthConfigEventHandlerMockRecorder) GenericAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericAuthConfig", reflect.TypeOf((*MockAuthConfigEventHandler)(nil).GenericAuthConfig), obj)
}

// MockAuthConfigEventWatcher is a mock of AuthConfigEventWatcher interface
type MockAuthConfigEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigEventWatcherMockRecorder
}

// MockAuthConfigEventWatcherMockRecorder is the mock recorder for MockAuthConfigEventWatcher
type MockAuthConfigEventWatcherMockRecorder struct {
	mock *MockAuthConfigEventWatcher
}

// NewMockAuthConfigEventWatcher creates a new mock instance
func NewMockAuthConfigEventWatcher(ctrl *gomock.Controller) *MockAuthConfigEventWatcher {
	mock := &MockAuthConfigEventWatcher{ctrl: ctrl}
	mock.recorder = &MockAuthConfigEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthConfigEventWatcher) EXPECT() *MockAuthConfigEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockAuthConfigEventWatcher) AddEventHandler(ctx context.Context, h controller.AuthConfigEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockAuthConfigEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockAuthConfigEventWatcher)(nil).AddEventHandler), varargs...)
}
