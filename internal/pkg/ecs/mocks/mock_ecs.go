// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/ecs/ecs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	ecs "github.com/aws/copilot-cli/internal/pkg/aws/ecs"
	resourcegroups "github.com/aws/copilot-cli/internal/pkg/aws/resourcegroups"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockresourceGetter is a mock of resourceGetter interface
type MockresourceGetter struct {
	ctrl     *gomock.Controller
	recorder *MockresourceGetterMockRecorder
}

// MockresourceGetterMockRecorder is the mock recorder for MockresourceGetter
type MockresourceGetterMockRecorder struct {
	mock *MockresourceGetter
}

// NewMockresourceGetter creates a new mock instance
func NewMockresourceGetter(ctrl *gomock.Controller) *MockresourceGetter {
	mock := &MockresourceGetter{ctrl: ctrl}
	mock.recorder = &MockresourceGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockresourceGetter) EXPECT() *MockresourceGetterMockRecorder {
	return m.recorder
}

// GetResourcesByTags mocks base method
func (m *MockresourceGetter) GetResourcesByTags(resourceType string, tags map[string]string) ([]*resourcegroups.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesByTags", resourceType, tags)
	ret0, _ := ret[0].([]*resourcegroups.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcesByTags indicates an expected call of GetResourcesByTags
func (mr *MockresourceGetterMockRecorder) GetResourcesByTags(resourceType, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesByTags", reflect.TypeOf((*MockresourceGetter)(nil).GetResourcesByTags), resourceType, tags)
}

// MockecsClient is a mock of ecsClient interface
type MockecsClient struct {
	ctrl     *gomock.Controller
	recorder *MockecsClientMockRecorder
}

// MockecsClientMockRecorder is the mock recorder for MockecsClient
type MockecsClientMockRecorder struct {
	mock *MockecsClient
}

// NewMockecsClient creates a new mock instance
func NewMockecsClient(ctrl *gomock.Controller) *MockecsClient {
	mock := &MockecsClient{ctrl: ctrl}
	mock.recorder = &MockecsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockecsClient) EXPECT() *MockecsClientMockRecorder {
	return m.recorder
}

// RunningTasksInFamily mocks base method
func (m *MockecsClient) RunningTasksInFamily(cluster, family string) ([]*ecs.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunningTasksInFamily", cluster, family)
	ret0, _ := ret[0].([]*ecs.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunningTasksInFamily indicates an expected call of RunningTasksInFamily
func (mr *MockecsClientMockRecorder) RunningTasksInFamily(cluster, family interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunningTasksInFamily", reflect.TypeOf((*MockecsClient)(nil).RunningTasksInFamily), cluster, family)
}

// RunningTasks mocks base method
func (m *MockecsClient) RunningTasks(cluster string) ([]*ecs.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunningTasks", cluster)
	ret0, _ := ret[0].([]*ecs.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunningTasks indicates an expected call of RunningTasks
func (mr *MockecsClientMockRecorder) RunningTasks(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunningTasks", reflect.TypeOf((*MockecsClient)(nil).RunningTasks), cluster)
}

// ServiceTasks mocks base method
func (m *MockecsClient) ServiceTasks(clusterName, serviceName string) ([]*ecs.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceTasks", clusterName, serviceName)
	ret0, _ := ret[0].([]*ecs.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceTasks indicates an expected call of ServiceTasks
func (mr *MockecsClientMockRecorder) ServiceTasks(clusterName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceTasks", reflect.TypeOf((*MockecsClient)(nil).ServiceTasks), clusterName, serviceName)
}

// DefaultCluster mocks base method
func (m *MockecsClient) DefaultCluster() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultCluster")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultCluster indicates an expected call of DefaultCluster
func (mr *MockecsClientMockRecorder) DefaultCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultCluster", reflect.TypeOf((*MockecsClient)(nil).DefaultCluster))
}

// StopTasks mocks base method
func (m *MockecsClient) StopTasks(tasks []string, opts ...ecs.StopTasksOpts) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{tasks}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopTasks", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopTasks indicates an expected call of StopTasks
func (mr *MockecsClientMockRecorder) StopTasks(tasks interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tasks}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopTasks", reflect.TypeOf((*MockecsClient)(nil).StopTasks), varargs...)
}
