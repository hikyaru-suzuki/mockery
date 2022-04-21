// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	foo "github.com/vektra/mockery/v2/pkg/fixtures/example_project/foo"

	testing "testing"
)

// Foo is an autogenerated mock type for the Foo type
type Foo struct {
	mock.Mock
}

// DoFoo provides a mock function with given fields:
func (_m *Foo) DoFoo() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetBaz provides a mock function with given fields:
func (_m *Foo) GetBaz() (*foo.Baz, error) {
	ret := _m.Called()

	var r0 *foo.Baz
	if rf, ok := ret.Get(0).(func() *foo.Baz); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*foo.Baz)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFoo creates a new instance of Foo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewFoo(t testing.TB) *Foo {
	mock := &Foo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
