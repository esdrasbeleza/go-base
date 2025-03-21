// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=interface_mock.go -package=main
//

// Package main is a generated GoMock package.
package main

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMyInterface is a mock of MyInterface interface.
type MockMyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMyInterfaceMockRecorder
	isgomock struct{}
}

// MockMyInterfaceMockRecorder is the mock recorder for MockMyInterface.
type MockMyInterfaceMockRecorder struct {
	mock *MockMyInterface
}

// NewMockMyInterface creates a new mock instance.
func NewMockMyInterface(ctrl *gomock.Controller) *MockMyInterface {
	mock := &MockMyInterface{ctrl: ctrl}
	mock.recorder = &MockMyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMyInterface) EXPECT() *MockMyInterfaceMockRecorder {
	return m.recorder
}

// SomeFunction mocks base method.
func (m *MockMyInterface) SomeFunction(ctx context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SomeFunction", ctx, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SomeFunction indicates an expected call of SomeFunction.
func (mr *MockMyInterfaceMockRecorder) SomeFunction(ctx, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SomeFunction", reflect.TypeOf((*MockMyInterface)(nil).SomeFunction), ctx, arg1, arg2)
}
