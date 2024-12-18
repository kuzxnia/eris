// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/workflow/workflow_repository.go
//
// Generated by this command:
//
//	mockgen -source=pkg/workflow/workflow_repository.go -destination=tests/unit/mocks/workflow_repository_mock.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	workflow "github.com/kuzxnia/eris/control-plane/pkg/workflow"
	gomock "go.uber.org/mock/gomock"
)

// MockWorkflowRepository is a mock of WorkflowRepository interface.
type MockWorkflowRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowRepositoryMockRecorder
	isgomock struct{}
}

// MockWorkflowRepositoryMockRecorder is the mock recorder for MockWorkflowRepository.
type MockWorkflowRepositoryMockRecorder struct {
	mock *MockWorkflowRepository
}

// NewMockWorkflowRepository creates a new mock instance.
func NewMockWorkflowRepository(ctrl *gomock.Controller) *MockWorkflowRepository {
	mock := &MockWorkflowRepository{ctrl: ctrl}
	mock.recorder = &MockWorkflowRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkflowRepository) EXPECT() *MockWorkflowRepositoryMockRecorder {
	return m.recorder
}

// GetWorkflow mocks base method.
func (m *MockWorkflowRepository) GetWorkflow(arg0 string) (*workflow.Workflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflow", arg0)
	ret0, _ := ret[0].(*workflow.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflow indicates an expected call of GetWorkflow.
func (mr *MockWorkflowRepositoryMockRecorder) GetWorkflow(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflow", reflect.TypeOf((*MockWorkflowRepository)(nil).GetWorkflow), arg0)
}

// SaveWorkflow mocks base method.
func (m *MockWorkflowRepository) SaveWorkflow(arg0 *workflow.Workflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWorkflow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWorkflow indicates an expected call of SaveWorkflow.
func (mr *MockWorkflowRepositoryMockRecorder) SaveWorkflow(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWorkflow", reflect.TypeOf((*MockWorkflowRepository)(nil).SaveWorkflow), arg0)
}
