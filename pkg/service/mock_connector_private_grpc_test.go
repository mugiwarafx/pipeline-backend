// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/instill-ai/protogen-go/vdp/connector/v1alpha (interfaces: ConnectorPrivateServiceClient)

// Package service_test is a generated GoMock package.
package service_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	connectorv1alpha "github.com/instill-ai/protogen-go/vdp/connector/v1alpha"
	grpc "google.golang.org/grpc"
)

// MockConnectorPrivateServiceClient is a mock of ConnectorPrivateServiceClient interface.
type MockConnectorPrivateServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockConnectorPrivateServiceClientMockRecorder
}

// MockConnectorPrivateServiceClientMockRecorder is the mock recorder for MockConnectorPrivateServiceClient.
type MockConnectorPrivateServiceClientMockRecorder struct {
	mock *MockConnectorPrivateServiceClient
}

// NewMockConnectorPrivateServiceClient creates a new mock instance.
func NewMockConnectorPrivateServiceClient(ctrl *gomock.Controller) *MockConnectorPrivateServiceClient {
	mock := &MockConnectorPrivateServiceClient{ctrl: ctrl}
	mock.recorder = &MockConnectorPrivateServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectorPrivateServiceClient) EXPECT() *MockConnectorPrivateServiceClientMockRecorder {
	return m.recorder
}

// CheckDestinationConnector mocks base method.
func (m *MockConnectorPrivateServiceClient) CheckDestinationConnector(arg0 context.Context, arg1 *connectorv1alpha.CheckDestinationConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.CheckDestinationConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckDestinationConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.CheckDestinationConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDestinationConnector indicates an expected call of CheckDestinationConnector.
func (mr *MockConnectorPrivateServiceClientMockRecorder) CheckDestinationConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDestinationConnector", reflect.TypeOf((*MockConnectorPrivateServiceClient)(nil).CheckDestinationConnector), varargs...)
}

// CheckSourceConnector mocks base method.
func (m *MockConnectorPrivateServiceClient) CheckSourceConnector(arg0 context.Context, arg1 *connectorv1alpha.CheckSourceConnectorRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.CheckSourceConnectorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckSourceConnector", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.CheckSourceConnectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSourceConnector indicates an expected call of CheckSourceConnector.
func (mr *MockConnectorPrivateServiceClientMockRecorder) CheckSourceConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSourceConnector", reflect.TypeOf((*MockConnectorPrivateServiceClient)(nil).CheckSourceConnector), varargs...)
}

// ListDestinationConnectorsAdmin mocks base method.
func (m *MockConnectorPrivateServiceClient) ListDestinationConnectorsAdmin(arg0 context.Context, arg1 *connectorv1alpha.ListDestinationConnectorsAdminRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.ListDestinationConnectorsAdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDestinationConnectorsAdmin", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.ListDestinationConnectorsAdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDestinationConnectorsAdmin indicates an expected call of ListDestinationConnectorsAdmin.
func (mr *MockConnectorPrivateServiceClientMockRecorder) ListDestinationConnectorsAdmin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDestinationConnectorsAdmin", reflect.TypeOf((*MockConnectorPrivateServiceClient)(nil).ListDestinationConnectorsAdmin), varargs...)
}

// ListSourceConnectorsAdmin mocks base method.
func (m *MockConnectorPrivateServiceClient) ListSourceConnectorsAdmin(arg0 context.Context, arg1 *connectorv1alpha.ListSourceConnectorsAdminRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.ListSourceConnectorsAdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSourceConnectorsAdmin", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.ListSourceConnectorsAdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSourceConnectorsAdmin indicates an expected call of ListSourceConnectorsAdmin.
func (mr *MockConnectorPrivateServiceClientMockRecorder) ListSourceConnectorsAdmin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSourceConnectorsAdmin", reflect.TypeOf((*MockConnectorPrivateServiceClient)(nil).ListSourceConnectorsAdmin), varargs...)
}

// LookUpDestinationConnectorAdmin mocks base method.
func (m *MockConnectorPrivateServiceClient) LookUpDestinationConnectorAdmin(arg0 context.Context, arg1 *connectorv1alpha.LookUpDestinationConnectorAdminRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.LookUpDestinationConnectorAdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LookUpDestinationConnectorAdmin", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.LookUpDestinationConnectorAdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookUpDestinationConnectorAdmin indicates an expected call of LookUpDestinationConnectorAdmin.
func (mr *MockConnectorPrivateServiceClientMockRecorder) LookUpDestinationConnectorAdmin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookUpDestinationConnectorAdmin", reflect.TypeOf((*MockConnectorPrivateServiceClient)(nil).LookUpDestinationConnectorAdmin), varargs...)
}

// LookUpSourceConnectorAdmin mocks base method.
func (m *MockConnectorPrivateServiceClient) LookUpSourceConnectorAdmin(arg0 context.Context, arg1 *connectorv1alpha.LookUpSourceConnectorAdminRequest, arg2 ...grpc.CallOption) (*connectorv1alpha.LookUpSourceConnectorAdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LookUpSourceConnectorAdmin", varargs...)
	ret0, _ := ret[0].(*connectorv1alpha.LookUpSourceConnectorAdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookUpSourceConnectorAdmin indicates an expected call of LookUpSourceConnectorAdmin.
func (mr *MockConnectorPrivateServiceClientMockRecorder) LookUpSourceConnectorAdmin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookUpSourceConnectorAdmin", reflect.TypeOf((*MockConnectorPrivateServiceClient)(nil).LookUpSourceConnectorAdmin), varargs...)
}
