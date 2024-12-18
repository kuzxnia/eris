// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/agent/agent_repository.go
//
// Generated by this command:
//
//	mockgen -source=pkg/agent/agent_repository.go -destination=tests/unit/mocks/agent_repository_mock.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	agent "github.com/kuzxnia/eris/control-plane/pkg/agent"
	gomock "go.uber.org/mock/gomock"
)

// MockAgentRepository is a mock of AgentRepository interface.
type MockAgentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAgentRepositoryMockRecorder
	isgomock struct{}
}

// MockAgentRepositoryMockRecorder is the mock recorder for MockAgentRepository.
type MockAgentRepositoryMockRecorder struct {
	mock *MockAgentRepository
}

// NewMockAgentRepository creates a new mock instance.
func NewMockAgentRepository(ctrl *gomock.Controller) *MockAgentRepository {
	mock := &MockAgentRepository{ctrl: ctrl}
	mock.recorder = &MockAgentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentRepository) EXPECT() *MockAgentRepositoryMockRecorder {
	return m.recorder
}

// GetAgents mocks base method.
func (m *MockAgentRepository) GetAgents() ([]*agent.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgents")
	ret0, _ := ret[0].([]*agent.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgents indicates an expected call of GetAgents.
func (mr *MockAgentRepositoryMockRecorder) GetAgents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgents", reflect.TypeOf((*MockAgentRepository)(nil).GetAgents))
}

// SaveAgent mocks base method.
func (m *MockAgentRepository) SaveAgent(arg0 *agent.Agent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAgent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAgent indicates an expected call of SaveAgent.
func (mr *MockAgentRepositoryMockRecorder) SaveAgent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAgent", reflect.TypeOf((*MockAgentRepository)(nil).SaveAgent), arg0)
}
