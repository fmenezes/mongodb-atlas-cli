// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fmenezes/mongodb-atlas-cli/v2/internal/store (interfaces: AlertDescriber,AlertLister,AlertAcknowledger)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockAlertDescriber is a mock of AlertDescriber interface.
type MockAlertDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockAlertDescriberMockRecorder
}

// MockAlertDescriberMockRecorder is the mock recorder for MockAlertDescriber.
type MockAlertDescriberMockRecorder struct {
	mock *MockAlertDescriber
}

// NewMockAlertDescriber creates a new mock instance.
func NewMockAlertDescriber(ctrl *gomock.Controller) *MockAlertDescriber {
	mock := &MockAlertDescriber{ctrl: ctrl}
	mock.recorder = &MockAlertDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertDescriber) EXPECT() *MockAlertDescriberMockRecorder {
	return m.recorder
}

// Alert mocks base method.
func (m *MockAlertDescriber) Alert(arg0, arg1 string) (*mongodbatlas.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alert", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Alert indicates an expected call of Alert.
func (mr *MockAlertDescriberMockRecorder) Alert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alert", reflect.TypeOf((*MockAlertDescriber)(nil).Alert), arg0, arg1)
}

// MockAlertLister is a mock of AlertLister interface.
type MockAlertLister struct {
	ctrl     *gomock.Controller
	recorder *MockAlertListerMockRecorder
}

// MockAlertListerMockRecorder is the mock recorder for MockAlertLister.
type MockAlertListerMockRecorder struct {
	mock *MockAlertLister
}

// NewMockAlertLister creates a new mock instance.
func NewMockAlertLister(ctrl *gomock.Controller) *MockAlertLister {
	mock := &MockAlertLister{ctrl: ctrl}
	mock.recorder = &MockAlertListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertLister) EXPECT() *MockAlertListerMockRecorder {
	return m.recorder
}

// Alerts mocks base method.
func (m *MockAlertLister) Alerts(arg0 string, arg1 *mongodbatlas.AlertsListOptions) (*mongodbatlas.AlertsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alerts", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.AlertsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Alerts indicates an expected call of Alerts.
func (mr *MockAlertListerMockRecorder) Alerts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alerts", reflect.TypeOf((*MockAlertLister)(nil).Alerts), arg0, arg1)
}

// MockAlertAcknowledger is a mock of AlertAcknowledger interface.
type MockAlertAcknowledger struct {
	ctrl     *gomock.Controller
	recorder *MockAlertAcknowledgerMockRecorder
}

// MockAlertAcknowledgerMockRecorder is the mock recorder for MockAlertAcknowledger.
type MockAlertAcknowledgerMockRecorder struct {
	mock *MockAlertAcknowledger
}

// NewMockAlertAcknowledger creates a new mock instance.
func NewMockAlertAcknowledger(ctrl *gomock.Controller) *MockAlertAcknowledger {
	mock := &MockAlertAcknowledger{ctrl: ctrl}
	mock.recorder = &MockAlertAcknowledgerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertAcknowledger) EXPECT() *MockAlertAcknowledgerMockRecorder {
	return m.recorder
}

// AcknowledgeAlert mocks base method.
func (m *MockAlertAcknowledger) AcknowledgeAlert(arg0, arg1 string, arg2 *mongodbatlas.AcknowledgeRequest) (*mongodbatlas.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcknowledgeAlert", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcknowledgeAlert indicates an expected call of AcknowledgeAlert.
func (mr *MockAlertAcknowledgerMockRecorder) AcknowledgeAlert(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcknowledgeAlert", reflect.TypeOf((*MockAlertAcknowledger)(nil).AcknowledgeAlert), arg0, arg1, arg2)
}
