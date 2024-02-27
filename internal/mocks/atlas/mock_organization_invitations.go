// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fmenezes/mongodb-atlas-cli/internal/store/atlas (interfaces: OrganizationInvitationLister,OrganizationInvitationDeleter,OrganizationInvitationDescriber,OrganizationInvitationUpdater,OrganizationInviter)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockOrganizationInvitationLister is a mock of OrganizationInvitationLister interface.
type MockOrganizationInvitationLister struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationInvitationListerMockRecorder
}

// MockOrganizationInvitationListerMockRecorder is the mock recorder for MockOrganizationInvitationLister.
type MockOrganizationInvitationListerMockRecorder struct {
	mock *MockOrganizationInvitationLister
}

// NewMockOrganizationInvitationLister creates a new mock instance.
func NewMockOrganizationInvitationLister(ctrl *gomock.Controller) *MockOrganizationInvitationLister {
	mock := &MockOrganizationInvitationLister{ctrl: ctrl}
	mock.recorder = &MockOrganizationInvitationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationInvitationLister) EXPECT() *MockOrganizationInvitationListerMockRecorder {
	return m.recorder
}

// OrganizationInvitations mocks base method.
func (m *MockOrganizationInvitationLister) OrganizationInvitations(arg0 *admin.ListOrganizationInvitationsApiParams) ([]admin.OrganizationInvitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationInvitations", arg0)
	ret0, _ := ret[0].([]admin.OrganizationInvitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationInvitations indicates an expected call of OrganizationInvitations.
func (mr *MockOrganizationInvitationListerMockRecorder) OrganizationInvitations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationInvitations", reflect.TypeOf((*MockOrganizationInvitationLister)(nil).OrganizationInvitations), arg0)
}

// MockOrganizationInvitationDeleter is a mock of OrganizationInvitationDeleter interface.
type MockOrganizationInvitationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationInvitationDeleterMockRecorder
}

// MockOrganizationInvitationDeleterMockRecorder is the mock recorder for MockOrganizationInvitationDeleter.
type MockOrganizationInvitationDeleterMockRecorder struct {
	mock *MockOrganizationInvitationDeleter
}

// NewMockOrganizationInvitationDeleter creates a new mock instance.
func NewMockOrganizationInvitationDeleter(ctrl *gomock.Controller) *MockOrganizationInvitationDeleter {
	mock := &MockOrganizationInvitationDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationInvitationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationInvitationDeleter) EXPECT() *MockOrganizationInvitationDeleterMockRecorder {
	return m.recorder
}

// DeleteInvitation mocks base method.
func (m *MockOrganizationInvitationDeleter) DeleteInvitation(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInvitation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInvitation indicates an expected call of DeleteInvitation.
func (mr *MockOrganizationInvitationDeleterMockRecorder) DeleteInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInvitation", reflect.TypeOf((*MockOrganizationInvitationDeleter)(nil).DeleteInvitation), arg0, arg1)
}

// MockOrganizationInvitationDescriber is a mock of OrganizationInvitationDescriber interface.
type MockOrganizationInvitationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationInvitationDescriberMockRecorder
}

// MockOrganizationInvitationDescriberMockRecorder is the mock recorder for MockOrganizationInvitationDescriber.
type MockOrganizationInvitationDescriberMockRecorder struct {
	mock *MockOrganizationInvitationDescriber
}

// NewMockOrganizationInvitationDescriber creates a new mock instance.
func NewMockOrganizationInvitationDescriber(ctrl *gomock.Controller) *MockOrganizationInvitationDescriber {
	mock := &MockOrganizationInvitationDescriber{ctrl: ctrl}
	mock.recorder = &MockOrganizationInvitationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationInvitationDescriber) EXPECT() *MockOrganizationInvitationDescriberMockRecorder {
	return m.recorder
}

// OrganizationInvitation mocks base method.
func (m *MockOrganizationInvitationDescriber) OrganizationInvitation(arg0, arg1 string) (*admin.OrganizationInvitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationInvitation", arg0, arg1)
	ret0, _ := ret[0].(*admin.OrganizationInvitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationInvitation indicates an expected call of OrganizationInvitation.
func (mr *MockOrganizationInvitationDescriberMockRecorder) OrganizationInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationInvitation", reflect.TypeOf((*MockOrganizationInvitationDescriber)(nil).OrganizationInvitation), arg0, arg1)
}

// MockOrganizationInvitationUpdater is a mock of OrganizationInvitationUpdater interface.
type MockOrganizationInvitationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationInvitationUpdaterMockRecorder
}

// MockOrganizationInvitationUpdaterMockRecorder is the mock recorder for MockOrganizationInvitationUpdater.
type MockOrganizationInvitationUpdaterMockRecorder struct {
	mock *MockOrganizationInvitationUpdater
}

// NewMockOrganizationInvitationUpdater creates a new mock instance.
func NewMockOrganizationInvitationUpdater(ctrl *gomock.Controller) *MockOrganizationInvitationUpdater {
	mock := &MockOrganizationInvitationUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationInvitationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationInvitationUpdater) EXPECT() *MockOrganizationInvitationUpdaterMockRecorder {
	return m.recorder
}

// UpdateOrganizationInvitation mocks base method.
func (m *MockOrganizationInvitationUpdater) UpdateOrganizationInvitation(arg0, arg1 string, arg2 *admin.OrganizationInvitationRequest) (*admin.OrganizationInvitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrganizationInvitation", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.OrganizationInvitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrganizationInvitation indicates an expected call of UpdateOrganizationInvitation.
func (mr *MockOrganizationInvitationUpdaterMockRecorder) UpdateOrganizationInvitation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganizationInvitation", reflect.TypeOf((*MockOrganizationInvitationUpdater)(nil).UpdateOrganizationInvitation), arg0, arg1, arg2)
}

// MockOrganizationInviter is a mock of OrganizationInviter interface.
type MockOrganizationInviter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationInviterMockRecorder
}

// MockOrganizationInviterMockRecorder is the mock recorder for MockOrganizationInviter.
type MockOrganizationInviterMockRecorder struct {
	mock *MockOrganizationInviter
}

// NewMockOrganizationInviter creates a new mock instance.
func NewMockOrganizationInviter(ctrl *gomock.Controller) *MockOrganizationInviter {
	mock := &MockOrganizationInviter{ctrl: ctrl}
	mock.recorder = &MockOrganizationInviterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationInviter) EXPECT() *MockOrganizationInviterMockRecorder {
	return m.recorder
}

// InviteUser mocks base method.
func (m *MockOrganizationInviter) InviteUser(arg0 string, arg1 *admin.OrganizationInvitationRequest) (*admin.OrganizationInvitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InviteUser", arg0, arg1)
	ret0, _ := ret[0].(*admin.OrganizationInvitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InviteUser indicates an expected call of InviteUser.
func (mr *MockOrganizationInviterMockRecorder) InviteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUser", reflect.TypeOf((*MockOrganizationInviter)(nil).InviteUser), arg0, arg1)
}
