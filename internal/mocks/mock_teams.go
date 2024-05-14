// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/store (interfaces: TeamLister,TeamDescriber,TeamCreator,TeamRenamer,TeamDeleter,TeamAdder,TeamUserRemover,TeamRolesUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115014/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockTeamLister is a mock of TeamLister interface.
type MockTeamLister struct {
	ctrl     *gomock.Controller
	recorder *MockTeamListerMockRecorder
}

// MockTeamListerMockRecorder is the mock recorder for MockTeamLister.
type MockTeamListerMockRecorder struct {
	mock *MockTeamLister
}

// NewMockTeamLister creates a new mock instance.
func NewMockTeamLister(ctrl *gomock.Controller) *MockTeamLister {
	mock := &MockTeamLister{ctrl: ctrl}
	mock.recorder = &MockTeamListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamLister) EXPECT() *MockTeamListerMockRecorder {
	return m.recorder
}

// Teams mocks base method.
func (m *MockTeamLister) Teams(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedTeam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teams", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedTeam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teams indicates an expected call of Teams.
func (mr *MockTeamListerMockRecorder) Teams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teams", reflect.TypeOf((*MockTeamLister)(nil).Teams), arg0, arg1)
}

// MockTeamDescriber is a mock of TeamDescriber interface.
type MockTeamDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockTeamDescriberMockRecorder
}

// MockTeamDescriberMockRecorder is the mock recorder for MockTeamDescriber.
type MockTeamDescriberMockRecorder struct {
	mock *MockTeamDescriber
}

// NewMockTeamDescriber creates a new mock instance.
func NewMockTeamDescriber(ctrl *gomock.Controller) *MockTeamDescriber {
	mock := &MockTeamDescriber{ctrl: ctrl}
	mock.recorder = &MockTeamDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamDescriber) EXPECT() *MockTeamDescriberMockRecorder {
	return m.recorder
}

// TeamByID mocks base method.
func (m *MockTeamDescriber) TeamByID(arg0, arg1 string) (*admin.TeamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByID", arg0, arg1)
	ret0, _ := ret[0].(*admin.TeamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByID indicates an expected call of TeamByID.
func (mr *MockTeamDescriberMockRecorder) TeamByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByID", reflect.TypeOf((*MockTeamDescriber)(nil).TeamByID), arg0, arg1)
}

// TeamByName mocks base method.
func (m *MockTeamDescriber) TeamByName(arg0, arg1 string) (*admin.TeamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByName", arg0, arg1)
	ret0, _ := ret[0].(*admin.TeamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByName indicates an expected call of TeamByName.
func (mr *MockTeamDescriberMockRecorder) TeamByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByName", reflect.TypeOf((*MockTeamDescriber)(nil).TeamByName), arg0, arg1)
}

// MockTeamCreator is a mock of TeamCreator interface.
type MockTeamCreator struct {
	ctrl     *gomock.Controller
	recorder *MockTeamCreatorMockRecorder
}

// MockTeamCreatorMockRecorder is the mock recorder for MockTeamCreator.
type MockTeamCreatorMockRecorder struct {
	mock *MockTeamCreator
}

// NewMockTeamCreator creates a new mock instance.
func NewMockTeamCreator(ctrl *gomock.Controller) *MockTeamCreator {
	mock := &MockTeamCreator{ctrl: ctrl}
	mock.recorder = &MockTeamCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamCreator) EXPECT() *MockTeamCreatorMockRecorder {
	return m.recorder
}

// CreateTeam mocks base method.
func (m *MockTeamCreator) CreateTeam(arg0 string, arg1 *admin.Team) (*admin.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", arg0, arg1)
	ret0, _ := ret[0].(*admin.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *MockTeamCreatorMockRecorder) CreateTeam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockTeamCreator)(nil).CreateTeam), arg0, arg1)
}

// MockTeamRenamer is a mock of TeamRenamer interface.
type MockTeamRenamer struct {
	ctrl     *gomock.Controller
	recorder *MockTeamRenamerMockRecorder
}

// MockTeamRenamerMockRecorder is the mock recorder for MockTeamRenamer.
type MockTeamRenamerMockRecorder struct {
	mock *MockTeamRenamer
}

// NewMockTeamRenamer creates a new mock instance.
func NewMockTeamRenamer(ctrl *gomock.Controller) *MockTeamRenamer {
	mock := &MockTeamRenamer{ctrl: ctrl}
	mock.recorder = &MockTeamRenamerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamRenamer) EXPECT() *MockTeamRenamerMockRecorder {
	return m.recorder
}

// RenameTeam mocks base method.
func (m *MockTeamRenamer) RenameTeam(arg0, arg1 string, arg2 *admin.TeamUpdate) (*admin.TeamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameTeam", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.TeamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameTeam indicates an expected call of RenameTeam.
func (mr *MockTeamRenamerMockRecorder) RenameTeam(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameTeam", reflect.TypeOf((*MockTeamRenamer)(nil).RenameTeam), arg0, arg1, arg2)
}

// MockTeamDeleter is a mock of TeamDeleter interface.
type MockTeamDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockTeamDeleterMockRecorder
}

// MockTeamDeleterMockRecorder is the mock recorder for MockTeamDeleter.
type MockTeamDeleterMockRecorder struct {
	mock *MockTeamDeleter
}

// NewMockTeamDeleter creates a new mock instance.
func NewMockTeamDeleter(ctrl *gomock.Controller) *MockTeamDeleter {
	mock := &MockTeamDeleter{ctrl: ctrl}
	mock.recorder = &MockTeamDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamDeleter) EXPECT() *MockTeamDeleterMockRecorder {
	return m.recorder
}

// DeleteTeam mocks base method.
func (m *MockTeamDeleter) DeleteTeam(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeam", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeam indicates an expected call of DeleteTeam.
func (mr *MockTeamDeleterMockRecorder) DeleteTeam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeam", reflect.TypeOf((*MockTeamDeleter)(nil).DeleteTeam), arg0, arg1)
}

// MockTeamAdder is a mock of TeamAdder interface.
type MockTeamAdder struct {
	ctrl     *gomock.Controller
	recorder *MockTeamAdderMockRecorder
}

// MockTeamAdderMockRecorder is the mock recorder for MockTeamAdder.
type MockTeamAdderMockRecorder struct {
	mock *MockTeamAdder
}

// NewMockTeamAdder creates a new mock instance.
func NewMockTeamAdder(ctrl *gomock.Controller) *MockTeamAdder {
	mock := &MockTeamAdder{ctrl: ctrl}
	mock.recorder = &MockTeamAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamAdder) EXPECT() *MockTeamAdderMockRecorder {
	return m.recorder
}

// AddUsersToTeam mocks base method.
func (m *MockTeamAdder) AddUsersToTeam(arg0, arg1 string, arg2 []admin.AddUserToTeam) (*admin.PaginatedApiAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUsersToTeam", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedApiAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUsersToTeam indicates an expected call of AddUsersToTeam.
func (mr *MockTeamAdderMockRecorder) AddUsersToTeam(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUsersToTeam", reflect.TypeOf((*MockTeamAdder)(nil).AddUsersToTeam), arg0, arg1, arg2)
}

// MockTeamUserRemover is a mock of TeamUserRemover interface.
type MockTeamUserRemover struct {
	ctrl     *gomock.Controller
	recorder *MockTeamUserRemoverMockRecorder
}

// MockTeamUserRemoverMockRecorder is the mock recorder for MockTeamUserRemover.
type MockTeamUserRemoverMockRecorder struct {
	mock *MockTeamUserRemover
}

// NewMockTeamUserRemover creates a new mock instance.
func NewMockTeamUserRemover(ctrl *gomock.Controller) *MockTeamUserRemover {
	mock := &MockTeamUserRemover{ctrl: ctrl}
	mock.recorder = &MockTeamUserRemoverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamUserRemover) EXPECT() *MockTeamUserRemoverMockRecorder {
	return m.recorder
}

// RemoveUserFromTeam mocks base method.
func (m *MockTeamUserRemover) RemoveUserFromTeam(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserFromTeam", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUserFromTeam indicates an expected call of RemoveUserFromTeam.
func (mr *MockTeamUserRemoverMockRecorder) RemoveUserFromTeam(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserFromTeam", reflect.TypeOf((*MockTeamUserRemover)(nil).RemoveUserFromTeam), arg0, arg1, arg2)
}

// MockTeamRolesUpdater is a mock of TeamRolesUpdater interface.
type MockTeamRolesUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockTeamRolesUpdaterMockRecorder
}

// MockTeamRolesUpdaterMockRecorder is the mock recorder for MockTeamRolesUpdater.
type MockTeamRolesUpdaterMockRecorder struct {
	mock *MockTeamRolesUpdater
}

// NewMockTeamRolesUpdater creates a new mock instance.
func NewMockTeamRolesUpdater(ctrl *gomock.Controller) *MockTeamRolesUpdater {
	mock := &MockTeamRolesUpdater{ctrl: ctrl}
	mock.recorder = &MockTeamRolesUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamRolesUpdater) EXPECT() *MockTeamRolesUpdaterMockRecorder {
	return m.recorder
}

// UpdateProjectTeamRoles mocks base method.
func (m *MockTeamRolesUpdater) UpdateProjectTeamRoles(arg0, arg1 string, arg2 *admin.TeamRole) (*admin.PaginatedTeamRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectTeamRoles", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedTeamRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectTeamRoles indicates an expected call of UpdateProjectTeamRoles.
func (mr *MockTeamRolesUpdaterMockRecorder) UpdateProjectTeamRoles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectTeamRoles", reflect.TypeOf((*MockTeamRolesUpdater)(nil).UpdateProjectTeamRoles), arg0, arg1, arg2)
}
