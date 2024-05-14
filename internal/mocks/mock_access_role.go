// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/store (interfaces: CloudProviderAccessRoleCreator,CloudProviderAccessRoleAuthorizer,CloudProviderAccessRoleLister,CloudProviderAccessRoleDeauthorizer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115014/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockCloudProviderAccessRoleCreator is a mock of CloudProviderAccessRoleCreator interface.
type MockCloudProviderAccessRoleCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderAccessRoleCreatorMockRecorder
}

// MockCloudProviderAccessRoleCreatorMockRecorder is the mock recorder for MockCloudProviderAccessRoleCreator.
type MockCloudProviderAccessRoleCreatorMockRecorder struct {
	mock *MockCloudProviderAccessRoleCreator
}

// NewMockCloudProviderAccessRoleCreator creates a new mock instance.
func NewMockCloudProviderAccessRoleCreator(ctrl *gomock.Controller) *MockCloudProviderAccessRoleCreator {
	mock := &MockCloudProviderAccessRoleCreator{ctrl: ctrl}
	mock.recorder = &MockCloudProviderAccessRoleCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudProviderAccessRoleCreator) EXPECT() *MockCloudProviderAccessRoleCreatorMockRecorder {
	return m.recorder
}

// CreateCloudProviderAccessRole mocks base method.
func (m *MockCloudProviderAccessRoleCreator) CreateCloudProviderAccessRole(arg0, arg1 string) (*admin.CloudProviderAccessRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCloudProviderAccessRole", arg0, arg1)
	ret0, _ := ret[0].(*admin.CloudProviderAccessRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCloudProviderAccessRole indicates an expected call of CreateCloudProviderAccessRole.
func (mr *MockCloudProviderAccessRoleCreatorMockRecorder) CreateCloudProviderAccessRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCloudProviderAccessRole", reflect.TypeOf((*MockCloudProviderAccessRoleCreator)(nil).CreateCloudProviderAccessRole), arg0, arg1)
}

// MockCloudProviderAccessRoleAuthorizer is a mock of CloudProviderAccessRoleAuthorizer interface.
type MockCloudProviderAccessRoleAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderAccessRoleAuthorizerMockRecorder
}

// MockCloudProviderAccessRoleAuthorizerMockRecorder is the mock recorder for MockCloudProviderAccessRoleAuthorizer.
type MockCloudProviderAccessRoleAuthorizerMockRecorder struct {
	mock *MockCloudProviderAccessRoleAuthorizer
}

// NewMockCloudProviderAccessRoleAuthorizer creates a new mock instance.
func NewMockCloudProviderAccessRoleAuthorizer(ctrl *gomock.Controller) *MockCloudProviderAccessRoleAuthorizer {
	mock := &MockCloudProviderAccessRoleAuthorizer{ctrl: ctrl}
	mock.recorder = &MockCloudProviderAccessRoleAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudProviderAccessRoleAuthorizer) EXPECT() *MockCloudProviderAccessRoleAuthorizerMockRecorder {
	return m.recorder
}

// AuthorizeCloudProviderAccessRole mocks base method.
func (m *MockCloudProviderAccessRoleAuthorizer) AuthorizeCloudProviderAccessRole(arg0, arg1 string, arg2 *mongodbatlas.CloudProviderAccessRoleRequest) (*admin.CloudProviderAccessRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeCloudProviderAccessRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.CloudProviderAccessRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizeCloudProviderAccessRole indicates an expected call of AuthorizeCloudProviderAccessRole.
func (mr *MockCloudProviderAccessRoleAuthorizerMockRecorder) AuthorizeCloudProviderAccessRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeCloudProviderAccessRole", reflect.TypeOf((*MockCloudProviderAccessRoleAuthorizer)(nil).AuthorizeCloudProviderAccessRole), arg0, arg1, arg2)
}

// MockCloudProviderAccessRoleLister is a mock of CloudProviderAccessRoleLister interface.
type MockCloudProviderAccessRoleLister struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderAccessRoleListerMockRecorder
}

// MockCloudProviderAccessRoleListerMockRecorder is the mock recorder for MockCloudProviderAccessRoleLister.
type MockCloudProviderAccessRoleListerMockRecorder struct {
	mock *MockCloudProviderAccessRoleLister
}

// NewMockCloudProviderAccessRoleLister creates a new mock instance.
func NewMockCloudProviderAccessRoleLister(ctrl *gomock.Controller) *MockCloudProviderAccessRoleLister {
	mock := &MockCloudProviderAccessRoleLister{ctrl: ctrl}
	mock.recorder = &MockCloudProviderAccessRoleListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudProviderAccessRoleLister) EXPECT() *MockCloudProviderAccessRoleListerMockRecorder {
	return m.recorder
}

// CloudProviderAccessRoles mocks base method.
func (m *MockCloudProviderAccessRoleLister) CloudProviderAccessRoles(arg0 string) (*admin.CloudProviderAccessRoles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderAccessRoles", arg0)
	ret0, _ := ret[0].(*admin.CloudProviderAccessRoles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderAccessRoles indicates an expected call of CloudProviderAccessRoles.
func (mr *MockCloudProviderAccessRoleListerMockRecorder) CloudProviderAccessRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderAccessRoles", reflect.TypeOf((*MockCloudProviderAccessRoleLister)(nil).CloudProviderAccessRoles), arg0)
}

// MockCloudProviderAccessRoleDeauthorizer is a mock of CloudProviderAccessRoleDeauthorizer interface.
type MockCloudProviderAccessRoleDeauthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderAccessRoleDeauthorizerMockRecorder
}

// MockCloudProviderAccessRoleDeauthorizerMockRecorder is the mock recorder for MockCloudProviderAccessRoleDeauthorizer.
type MockCloudProviderAccessRoleDeauthorizerMockRecorder struct {
	mock *MockCloudProviderAccessRoleDeauthorizer
}

// NewMockCloudProviderAccessRoleDeauthorizer creates a new mock instance.
func NewMockCloudProviderAccessRoleDeauthorizer(ctrl *gomock.Controller) *MockCloudProviderAccessRoleDeauthorizer {
	mock := &MockCloudProviderAccessRoleDeauthorizer{ctrl: ctrl}
	mock.recorder = &MockCloudProviderAccessRoleDeauthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudProviderAccessRoleDeauthorizer) EXPECT() *MockCloudProviderAccessRoleDeauthorizerMockRecorder {
	return m.recorder
}

// DeauthorizeCloudProviderAccessRoles mocks base method.
func (m *MockCloudProviderAccessRoleDeauthorizer) DeauthorizeCloudProviderAccessRoles(arg0 *mongodbatlas.CloudProviderDeauthorizationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeauthorizeCloudProviderAccessRoles", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeauthorizeCloudProviderAccessRoles indicates an expected call of DeauthorizeCloudProviderAccessRoles.
func (mr *MockCloudProviderAccessRoleDeauthorizerMockRecorder) DeauthorizeCloudProviderAccessRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeauthorizeCloudProviderAccessRoles", reflect.TypeOf((*MockCloudProviderAccessRoleDeauthorizer)(nil).DeauthorizeCloudProviderAccessRoles), arg0)
}
