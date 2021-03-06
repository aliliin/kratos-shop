// Code generated by MockGen. DO NOT EDIT.
// Source: user/internal/biz (interfaces: Transaction)

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// ExecTx mocks base method.
func (m *MockTransaction) ExecTx(arg0 context.Context, arg1 func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecTx indicates an expected call of ExecTx.
func (mr *MockTransactionMockRecorder) ExecTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTx", reflect.TypeOf((*MockTransaction)(nil).ExecTx), arg0, arg1)
}
