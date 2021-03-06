// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"
	controller "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMultiClusterRoleEventHandler is a mock of MultiClusterRoleEventHandler interface
type MockMultiClusterRoleEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMultiClusterRoleEventHandlerMockRecorder
}

// MockMultiClusterRoleEventHandlerMockRecorder is the mock recorder for MockMultiClusterRoleEventHandler
type MockMultiClusterRoleEventHandlerMockRecorder struct {
	mock *MockMultiClusterRoleEventHandler
}

// NewMockMultiClusterRoleEventHandler creates a new mock instance
func NewMockMultiClusterRoleEventHandler(ctrl *gomock.Controller) *MockMultiClusterRoleEventHandler {
	mock := &MockMultiClusterRoleEventHandler{ctrl: ctrl}
	mock.recorder = &MockMultiClusterRoleEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMultiClusterRoleEventHandler) EXPECT() *MockMultiClusterRoleEventHandlerMockRecorder {
	return m.recorder
}

// CreateMultiClusterRole mocks base method
func (m *MockMultiClusterRoleEventHandler) CreateMultiClusterRole(obj *v1alpha1.MultiClusterRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMultiClusterRole", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMultiClusterRole indicates an expected call of CreateMultiClusterRole
func (mr *MockMultiClusterRoleEventHandlerMockRecorder) CreateMultiClusterRole(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleEventHandler)(nil).CreateMultiClusterRole), obj)
}

// UpdateMultiClusterRole mocks base method
func (m *MockMultiClusterRoleEventHandler) UpdateMultiClusterRole(old, new *v1alpha1.MultiClusterRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMultiClusterRole", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiClusterRole indicates an expected call of UpdateMultiClusterRole
func (mr *MockMultiClusterRoleEventHandlerMockRecorder) UpdateMultiClusterRole(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleEventHandler)(nil).UpdateMultiClusterRole), old, new)
}

// DeleteMultiClusterRole mocks base method
func (m *MockMultiClusterRoleEventHandler) DeleteMultiClusterRole(obj *v1alpha1.MultiClusterRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMultiClusterRole", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiClusterRole indicates an expected call of DeleteMultiClusterRole
func (mr *MockMultiClusterRoleEventHandlerMockRecorder) DeleteMultiClusterRole(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleEventHandler)(nil).DeleteMultiClusterRole), obj)
}

// GenericMultiClusterRole mocks base method
func (m *MockMultiClusterRoleEventHandler) GenericMultiClusterRole(obj *v1alpha1.MultiClusterRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericMultiClusterRole", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericMultiClusterRole indicates an expected call of GenericMultiClusterRole
func (mr *MockMultiClusterRoleEventHandlerMockRecorder) GenericMultiClusterRole(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleEventHandler)(nil).GenericMultiClusterRole), obj)
}

// MockMultiClusterRoleEventWatcher is a mock of MultiClusterRoleEventWatcher interface
type MockMultiClusterRoleEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMultiClusterRoleEventWatcherMockRecorder
}

// MockMultiClusterRoleEventWatcherMockRecorder is the mock recorder for MockMultiClusterRoleEventWatcher
type MockMultiClusterRoleEventWatcherMockRecorder struct {
	mock *MockMultiClusterRoleEventWatcher
}

// NewMockMultiClusterRoleEventWatcher creates a new mock instance
func NewMockMultiClusterRoleEventWatcher(ctrl *gomock.Controller) *MockMultiClusterRoleEventWatcher {
	mock := &MockMultiClusterRoleEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMultiClusterRoleEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMultiClusterRoleEventWatcher) EXPECT() *MockMultiClusterRoleEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockMultiClusterRoleEventWatcher) AddEventHandler(ctx context.Context, h controller.MultiClusterRoleEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockMultiClusterRoleEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMultiClusterRoleEventWatcher)(nil).AddEventHandler), varargs...)
}
