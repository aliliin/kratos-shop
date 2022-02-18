// Code generated by MockGen. DO NOT EDIT.
// Source: user/internal/biz (interfaces: AddressRepo)

// Package mrepo is a generated GoMock package.
package mrepo

import (
	context "context"
	reflect "reflect"
	biz "user/internal/biz"

	gomock "github.com/golang/mock/gomock"
)

// MockAddressRepo is a mock of AddressRepo interface.
type MockAddressRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAddressRepoMockRecorder
}

// MockAddressRepoMockRecorder is the mock recorder for MockAddressRepo.
type MockAddressRepoMockRecorder struct {
	mock *MockAddressRepo
}

// NewMockAddressRepo creates a new mock instance.
func NewMockAddressRepo(ctrl *gomock.Controller) *MockAddressRepo {
	mock := &MockAddressRepo{ctrl: ctrl}
	mock.recorder = &MockAddressRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddressRepo) EXPECT() *MockAddressRepoMockRecorder {
	return m.recorder
}

// CreateAddress mocks base method.
func (m *MockAddressRepo) CreateAddress(arg0 context.Context, arg1 *biz.Address) (*biz.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddress", arg0, arg1)
	ret0, _ := ret[0].(*biz.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAddress indicates an expected call of CreateAddress.
func (mr *MockAddressRepoMockRecorder) CreateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddress", reflect.TypeOf((*MockAddressRepo)(nil).CreateAddress), arg0, arg1)
}