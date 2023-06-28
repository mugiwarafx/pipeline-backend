// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/instill-ai/protogen-go/vdp/connector/v1alpha (interfaces: ConnectorPublicServiceClient)

// Package service_test is a generated GoMock package.
package service_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	connectorv1alpha "github.com/instill-ai/protogen-go/vdp/connector/v1alpha"
	grpc "google.golang.org/grpc"
)

// MockConnectorPublicServiceClient is a mock of ConnectorPublicServiceClient interface.
type MockConnectorPublicServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockConnectorPublicServiceClientMockRecorder
}

// MockConnectorPublicServiceClientMockRecorder is the mock recorder for MockConnectorPublicServiceClient.
type MockConnectorPublicServiceClientMockRecorder struct {
	mock *MockConnectorPublicServiceClient
}

// NewMockConnectorPublicServiceClient creates a new mock instance.
func NewMockConnectorPublicServiceClient(ctrl *gomock.Controller) *MockConnectorPublicServiceClient {
	mock := &MockConnectorPublicServiceClient{ctrl: ctrl}
	mock.recorder = &MockConnectorPublicServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectorPublicServiceClient) EXPECT() *MockConnectorPublicServiceClientMockRecorder {
	return m.recorder
}

// ConnectConnector mocks base method.
func (m *MockConnectorPublicServiceClient) ConnectConnector(arg0 context.Context, arg1 *connectorv1alpha.ConnectConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.ConnectConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConnectConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.ConnectConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectConnector indicates an expected call of ConnectConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) ConnectConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).ConnectConnector), varargs...)
}

// CreateConnector mocks base method.
func (m *MockConnectorPublicServiceClient) CreateConnector(arg0 context.Context, arg1 *connectorv1alpha.CreateConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.CreateConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.CreateConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConnector indicates an expected call of CreateConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) CreateConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).CreateConnector), varargs...)
}

// DeleteConnector mocks base method.
func (m *MockConnectorPublicServiceClient) DeleteConnector(arg0 context.Context, arg1 *connectorv1alpha.DeleteConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.DeleteConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.DeleteConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteConnector indicates an expected call of DeleteConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) DeleteConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).DeleteConnector), varargs...)
}

// DisconnectConnector mocks base method.
func (m *MockConnectorPublicServiceClient) DisconnectConnector(arg0 context.Context, arg1 *connectorv1alpha.DisconnectConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.DisconnectConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisconnectConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.DisconnectConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisconnectConnector indicates an expected call of DisconnectConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) DisconnectConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisconnectConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).DisconnectConnector), varargs...)
}

// ExecuteConnector mocks base method.
func (m *MockConnectorPublicServiceClient) ExecuteConnector(arg0 context.Context, arg1 *connectorv1alpha.ExecuteConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.ExecuteConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.ExecuteConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteConnector indicates an expected call of ExecuteConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) ExecuteConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).ExecuteConnector), varargs...)
}

// GetConnector mocks base method.
func (m *MockConnectorPublicServiceClient) GetConnector(arg0 context.Context, arg1 *connectorv1alpha.GetConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.GetConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.GetConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnector indicates an expected call of GetConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) GetConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).GetConnector), varargs...)
}

// GetConnectorDefinition mocks base method.
func (m *MockConnectorPublicServiceClient) GetConnectorDefinition(arg0 context.Context, arg1 *connectorv1alpha.GetConnectorDefinitionRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.GetConnectorDefinitionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConnectorDefinition", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.GetConnectorDefinitionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnectorDefinition indicates an expected call of GetConnectorDefinition.
func (mr *MockConnectorPublicServiceClientMockRecorder) GetConnectorDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectorDefinition", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).GetConnectorDefinition), varargs...)
}

// ListConnectorDefinitions mocks base method.
func (m *MockConnectorPublicServiceClient) ListConnectorDefinitions(arg0 context.Context, arg1 *connectorv1alpha.ListConnectorDefinitionsRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.ListConnectorDefinitionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConnectorDefinitions", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.ListConnectorDefinitionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectorDefinitions indicates an expected call of ListConnectorDefinitions.
func (mr *MockConnectorPublicServiceClientMockRecorder) ListConnectorDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectorDefinitions", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).ListConnectorDefinitions), varargs...)
}

// ListConnectors mocks base method.
func (m *MockConnectorPublicServiceClient) ListConnectors(arg0 context.Context, arg1 *connectorv1alpha.ListConnectorsRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.ListConnectorsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConnectors", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.ListConnectorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectors indicates an expected call of ListConnectors.
func (mr *MockConnectorPublicServiceClientMockRecorder) ListConnectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectors", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).ListConnectors), varargs...)
}

// Liveness mocks base method.
func (m *MockConnectorPublicServiceClient) Liveness(arg0 context.Context, arg1 *connectorv1alpha.LivenessRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.LivenessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Liveness", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.LivenessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Liveness indicates an expected call of Liveness.
func (mr *MockConnectorPublicServiceClientMockRecorder) Liveness(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Liveness", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).Liveness), varargs...)
}

// LookUpConnector mocks base method.
func (m *MockConnectorPublicServiceClient) LookUpConnector(arg0 context.Context, arg1 *connectorv1alpha.LookUpConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.LookUpConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LookUpConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.LookUpConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookUpConnector indicates an expected call of LookUpConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) LookUpConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookUpConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).LookUpConnector), varargs...)
}

// Readiness mocks base method.
func (m *MockConnectorPublicServiceClient) Readiness(arg0 context.Context, arg1 *connectorv1alpha.ReadinessRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.ReadinessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Readiness", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.ReadinessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Readiness indicates an expected call of Readiness.
func (mr *MockConnectorPublicServiceClientMockRecorder) Readiness(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readiness", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).Readiness), varargs...)
}

// RenameConnector mocks base method.
func (m *MockConnectorPublicServiceClient) RenameConnector(arg0 context.Context, arg1 *connectorv1alpha.RenameConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.RenameConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenameConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.RenameConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameConnector indicates an expected call of RenameConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) RenameConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).RenameConnector), varargs...)
}

// TestConnector mocks base method.
func (m *MockConnectorPublicServiceClient) TestConnector(arg0 context.Context, arg1 *connectorv1alpha.TestConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.TestConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TestConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.TestConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TestConnector indicates an expected call of TestConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) TestConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).TestConnector), varargs...)
}

// UpdateConnector mocks base method.
func (m *MockConnectorPublicServiceClient) UpdateConnector(arg0 context.Context, arg1 *connectorv1alpha.UpdateConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.UpdateConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.UpdateConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConnector indicates an expected call of UpdateConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) UpdateConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).UpdateConnector), varargs...)
}

// WatchConnector mocks base method.
func (m *MockConnectorPublicServiceClient) WatchConnector(arg0 context.Context, arg1 *connectorv1alpha.WatchConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.WatchConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.WatchConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchConnector indicates an expected call of WatchConnector.
func (mr *MockConnectorPublicServiceClientMockRecorder) WatchConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchConnector", reflect.TypeOf((*MockConnectorPublicServiceClient)(nil).WatchConnector), varargs...)
}
