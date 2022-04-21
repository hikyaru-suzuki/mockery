// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// FuncArgsCollision is an autogenerated mock type for the FuncArgsCollision type
type FuncArgsCollision struct {
	mock.Mock
}

// Foo provides a mock function with given fields: ret
func (_m *FuncArgsCollision) Foo(ret interface{}) error {
	ret_1 := _m.Called(ret)

	var r0 error
	if rf, ok := ret_1.Get(0).(func(interface{}) error); ok {
		r0 = rf(ret)
	} else {
		r0 = ret_1.Error(0)
	}

	return r0
}

// NewFuncArgsCollision creates a new instance of FuncArgsCollision. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewFuncArgsCollision(t testing.TB) *FuncArgsCollision {
	mock := &FuncArgsCollision{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
