// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fmenezes/mongodb-atlas-cli/v2/internal/store (interfaces: MonitoringStarter,MonitoringStopper)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	opsmngr "go.mongodb.org/ops-manager/opsmngr"
)

// MockMonitoringStarter is a mock of MonitoringStarter interface.
type MockMonitoringStarter struct {
	ctrl     *gomock.Controller
	recorder *MockMonitoringStarterMockRecorder
}

// MockMonitoringStarterMockRecorder is the mock recorder for MockMonitoringStarter.
type MockMonitoringStarterMockRecorder struct {
	mock *MockMonitoringStarter
}

// NewMockMonitoringStarter creates a new mock instance.
func NewMockMonitoringStarter(ctrl *gomock.Controller) *MockMonitoringStarter {
	mock := &MockMonitoringStarter{ctrl: ctrl}
	mock.recorder = &MockMonitoringStarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitoringStarter) EXPECT() *MockMonitoringStarterMockRecorder {
	return m.recorder
}

// StartMonitoring mocks base method.
func (m *MockMonitoringStarter) StartMonitoring(arg0 string, arg1 *opsmngr.Host) (*opsmngr.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartMonitoring", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartMonitoring indicates an expected call of StartMonitoring.
func (mr *MockMonitoringStarterMockRecorder) StartMonitoring(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMonitoring", reflect.TypeOf((*MockMonitoringStarter)(nil).StartMonitoring), arg0, arg1)
}

// MockMonitoringStopper is a mock of MonitoringStopper interface.
type MockMonitoringStopper struct {
	ctrl     *gomock.Controller
	recorder *MockMonitoringStopperMockRecorder
}

// MockMonitoringStopperMockRecorder is the mock recorder for MockMonitoringStopper.
type MockMonitoringStopperMockRecorder struct {
	mock *MockMonitoringStopper
}

// NewMockMonitoringStopper creates a new mock instance.
func NewMockMonitoringStopper(ctrl *gomock.Controller) *MockMonitoringStopper {
	mock := &MockMonitoringStopper{ctrl: ctrl}
	mock.recorder = &MockMonitoringStopperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitoringStopper) EXPECT() *MockMonitoringStopperMockRecorder {
	return m.recorder
}

// StopMonitoring mocks base method.
func (m *MockMonitoringStopper) StopMonitoring(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopMonitoring", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopMonitoring indicates an expected call of StopMonitoring.
func (mr *MockMonitoringStopperMockRecorder) StopMonitoring(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopMonitoring", reflect.TypeOf((*MockMonitoringStopper)(nil).StopMonitoring), arg0, arg1)
}
