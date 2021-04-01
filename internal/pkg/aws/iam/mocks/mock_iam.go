// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/iam/iam.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	iam "github.com/aws/aws-sdk-go/service/iam"
	gomock "github.com/golang/mock/gomock"
)

// Mockapi is a mock of api interface.
type Mockapi struct {
	ctrl     *gomock.Controller
	recorder *MockapiMockRecorder
}

// MockapiMockRecorder is the mock recorder for Mockapi.
type MockapiMockRecorder struct {
	mock *Mockapi
}

// NewMockapi creates a new mock instance.
func NewMockapi(ctrl *gomock.Controller) *Mockapi {
	mock := &Mockapi{ctrl: ctrl}
	mock.recorder = &MockapiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockapi) EXPECT() *MockapiMockRecorder {
	return m.recorder
}

// CreateServiceLinkedRole mocks base method.
func (m *Mockapi) CreateServiceLinkedRole(input *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceLinkedRole", input)
	ret0, _ := ret[0].(*iam.CreateServiceLinkedRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServiceLinkedRole indicates an expected call of CreateServiceLinkedRole.
func (mr *MockapiMockRecorder) CreateServiceLinkedRole(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceLinkedRole", reflect.TypeOf((*Mockapi)(nil).CreateServiceLinkedRole), input)
}

// DeleteRole mocks base method.
func (m *Mockapi) DeleteRole(input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", input)
	ret0, _ := ret[0].(*iam.DeleteRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockapiMockRecorder) DeleteRole(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*Mockapi)(nil).DeleteRole), input)
}

// DeleteRolePolicy mocks base method.
func (m *Mockapi) DeleteRolePolicy(input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRolePolicy", input)
	ret0, _ := ret[0].(*iam.DeleteRolePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRolePolicy indicates an expected call of DeleteRolePolicy.
func (mr *MockapiMockRecorder) DeleteRolePolicy(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRolePolicy", reflect.TypeOf((*Mockapi)(nil).DeleteRolePolicy), input)
}

// ListRolePolicies mocks base method.
func (m *Mockapi) ListRolePolicies(input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRolePolicies", input)
	ret0, _ := ret[0].(*iam.ListRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRolePolicies indicates an expected call of ListRolePolicies.
func (mr *MockapiMockRecorder) ListRolePolicies(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRolePolicies", reflect.TypeOf((*Mockapi)(nil).ListRolePolicies), input)
}

// ListRoleTags mocks base method.
func (m *Mockapi) ListRoleTags(input *iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleTags", input)
	ret0, _ := ret[0].(*iam.ListRoleTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoleTags indicates an expected call of ListRoleTags.
func (mr *MockapiMockRecorder) ListRoleTags(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleTags", reflect.TypeOf((*Mockapi)(nil).ListRoleTags), input)
}
