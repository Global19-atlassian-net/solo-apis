// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"
	controller "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterMultiClusterRoleReconciler is a mock of MulticlusterMultiClusterRoleReconciler interface
type MockMulticlusterMultiClusterRoleReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMultiClusterRoleReconcilerMockRecorder
}

// MockMulticlusterMultiClusterRoleReconcilerMockRecorder is the mock recorder for MockMulticlusterMultiClusterRoleReconciler
type MockMulticlusterMultiClusterRoleReconcilerMockRecorder struct {
	mock *MockMulticlusterMultiClusterRoleReconciler
}

// NewMockMulticlusterMultiClusterRoleReconciler creates a new mock instance
func NewMockMulticlusterMultiClusterRoleReconciler(ctrl *gomock.Controller) *MockMulticlusterMultiClusterRoleReconciler {
	mock := &MockMulticlusterMultiClusterRoleReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMultiClusterRoleReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterMultiClusterRoleReconciler) EXPECT() *MockMulticlusterMultiClusterRoleReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMultiClusterRole mocks base method
func (m *MockMulticlusterMultiClusterRoleReconciler) ReconcileMultiClusterRole(clusterName string, obj *v1alpha1.MultiClusterRole) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMultiClusterRole", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMultiClusterRole indicates an expected call of ReconcileMultiClusterRole
func (mr *MockMulticlusterMultiClusterRoleReconcilerMockRecorder) ReconcileMultiClusterRole(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMultiClusterRole", reflect.TypeOf((*MockMulticlusterMultiClusterRoleReconciler)(nil).ReconcileMultiClusterRole), clusterName, obj)
}

// MockMulticlusterMultiClusterRoleDeletionReconciler is a mock of MulticlusterMultiClusterRoleDeletionReconciler interface
type MockMulticlusterMultiClusterRoleDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMultiClusterRoleDeletionReconcilerMockRecorder
}

// MockMulticlusterMultiClusterRoleDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterMultiClusterRoleDeletionReconciler
type MockMulticlusterMultiClusterRoleDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterMultiClusterRoleDeletionReconciler
}

// NewMockMulticlusterMultiClusterRoleDeletionReconciler creates a new mock instance
func NewMockMulticlusterMultiClusterRoleDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterMultiClusterRoleDeletionReconciler {
	mock := &MockMulticlusterMultiClusterRoleDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMultiClusterRoleDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterMultiClusterRoleDeletionReconciler) EXPECT() *MockMulticlusterMultiClusterRoleDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMultiClusterRoleDeletion mocks base method
func (m *MockMulticlusterMultiClusterRoleDeletionReconciler) ReconcileMultiClusterRoleDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMultiClusterRoleDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileMultiClusterRoleDeletion indicates an expected call of ReconcileMultiClusterRoleDeletion
func (mr *MockMulticlusterMultiClusterRoleDeletionReconcilerMockRecorder) ReconcileMultiClusterRoleDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMultiClusterRoleDeletion", reflect.TypeOf((*MockMulticlusterMultiClusterRoleDeletionReconciler)(nil).ReconcileMultiClusterRoleDeletion), clusterName, req)
}

// MockMulticlusterMultiClusterRoleReconcileLoop is a mock of MulticlusterMultiClusterRoleReconcileLoop interface
type MockMulticlusterMultiClusterRoleReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMultiClusterRoleReconcileLoopMockRecorder
}

// MockMulticlusterMultiClusterRoleReconcileLoopMockRecorder is the mock recorder for MockMulticlusterMultiClusterRoleReconcileLoop
type MockMulticlusterMultiClusterRoleReconcileLoopMockRecorder struct {
	mock *MockMulticlusterMultiClusterRoleReconcileLoop
}

// NewMockMulticlusterMultiClusterRoleReconcileLoop creates a new mock instance
func NewMockMulticlusterMultiClusterRoleReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterMultiClusterRoleReconcileLoop {
	mock := &MockMulticlusterMultiClusterRoleReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMultiClusterRoleReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterMultiClusterRoleReconcileLoop) EXPECT() *MockMulticlusterMultiClusterRoleReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterMultiClusterRoleReconciler mocks base method
func (m *MockMulticlusterMultiClusterRoleReconcileLoop) AddMulticlusterMultiClusterRoleReconciler(ctx context.Context, rec controller.MulticlusterMultiClusterRoleReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterMultiClusterRoleReconciler", varargs...)
}

// AddMulticlusterMultiClusterRoleReconciler indicates an expected call of AddMulticlusterMultiClusterRoleReconciler
func (mr *MockMulticlusterMultiClusterRoleReconcileLoopMockRecorder) AddMulticlusterMultiClusterRoleReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterMultiClusterRoleReconciler", reflect.TypeOf((*MockMulticlusterMultiClusterRoleReconcileLoop)(nil).AddMulticlusterMultiClusterRoleReconciler), varargs...)
}
