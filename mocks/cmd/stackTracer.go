// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	errors "github.com/pkg/errors"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// stackTracer is an autogenerated mock type for the stackTracer type
type stackTracer struct {
	mock.Mock
}

// StackTrace provides a mock function with given fields:
func (_m *stackTracer) StackTrace() errors.StackTrace {
	ret := _m.Called()

	var r0 errors.StackTrace
	if rf, ok := ret.Get(0).(func() errors.StackTrace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.StackTrace)
		}
	}

	return r0
}

// newStackTracer creates a new instance of stackTracer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newStackTracer(t testing.TB) *stackTracer {
	mock := &stackTracer{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
