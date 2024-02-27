// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fmenezes/mongodb-atlas-cli/internal/store/atlas (interfaces: PipelineRunsLister,PipelineRunsDescriber)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockPipelineRunsLister is a mock of PipelineRunsLister interface.
type MockPipelineRunsLister struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineRunsListerMockRecorder
}

// MockPipelineRunsListerMockRecorder is the mock recorder for MockPipelineRunsLister.
type MockPipelineRunsListerMockRecorder struct {
	mock *MockPipelineRunsLister
}

// NewMockPipelineRunsLister creates a new mock instance.
func NewMockPipelineRunsLister(ctrl *gomock.Controller) *MockPipelineRunsLister {
	mock := &MockPipelineRunsLister{ctrl: ctrl}
	mock.recorder = &MockPipelineRunsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelineRunsLister) EXPECT() *MockPipelineRunsListerMockRecorder {
	return m.recorder
}

// PipelineRuns mocks base method.
func (m *MockPipelineRunsLister) PipelineRuns(arg0, arg1 string) (*admin.PaginatedPipelineRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineRuns", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedPipelineRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineRuns indicates an expected call of PipelineRuns.
func (mr *MockPipelineRunsListerMockRecorder) PipelineRuns(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineRuns", reflect.TypeOf((*MockPipelineRunsLister)(nil).PipelineRuns), arg0, arg1)
}

// MockPipelineRunsDescriber is a mock of PipelineRunsDescriber interface.
type MockPipelineRunsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineRunsDescriberMockRecorder
}

// MockPipelineRunsDescriberMockRecorder is the mock recorder for MockPipelineRunsDescriber.
type MockPipelineRunsDescriberMockRecorder struct {
	mock *MockPipelineRunsDescriber
}

// NewMockPipelineRunsDescriber creates a new mock instance.
func NewMockPipelineRunsDescriber(ctrl *gomock.Controller) *MockPipelineRunsDescriber {
	mock := &MockPipelineRunsDescriber{ctrl: ctrl}
	mock.recorder = &MockPipelineRunsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelineRunsDescriber) EXPECT() *MockPipelineRunsDescriberMockRecorder {
	return m.recorder
}

// PipelineRun mocks base method.
func (m *MockPipelineRunsDescriber) PipelineRun(arg0, arg1, arg2 string) (*admin.IngestionPipelineRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineRun", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.IngestionPipelineRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineRun indicates an expected call of PipelineRun.
func (mr *MockPipelineRunsDescriberMockRecorder) PipelineRun(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineRun", reflect.TypeOf((*MockPipelineRunsDescriber)(nil).PipelineRun), arg0, arg1, arg2)
}
