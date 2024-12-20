// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/resource/resource_selector.go
//
// Generated by this command:
//
//	mockgen -source=pkg/resource/resource_selector.go -destination=tests/unit/mocks/resource_selector_mock.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	resource "github.com/kuzxnia/eris/control-plane/pkg/resource"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceSelector is a mock of ResourceSelector interface.
type MockResourceSelector struct {
	ctrl     *gomock.Controller
	recorder *MockResourceSelectorMockRecorder
	isgomock struct{}
}

// MockResourceSelectorMockRecorder is the mock recorder for MockResourceSelector.
type MockResourceSelectorMockRecorder struct {
	mock *MockResourceSelector
}

// NewMockResourceSelector creates a new mock instance.
func NewMockResourceSelector(ctrl *gomock.Controller) *MockResourceSelector {
	mock := &MockResourceSelector{ctrl: ctrl}
	mock.recorder = &MockResourceSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceSelector) EXPECT() *MockResourceSelectorMockRecorder {
	return m.recorder
}

// GetSelectorType mocks base method.
func (m *MockResourceSelector) GetSelectorType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelectorType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSelectorType indicates an expected call of GetSelectorType.
func (mr *MockResourceSelectorMockRecorder) GetSelectorType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelectorType", reflect.TypeOf((*MockResourceSelector)(nil).GetSelectorType))
}

// Select mocks base method.
func (m *MockResourceSelector) Select(arg0 *resource.Selector) ([]*resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", arg0)
	ret0, _ := ret[0].([]*resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockResourceSelectorMockRecorder) Select(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockResourceSelector)(nil).Select), arg0)
}
