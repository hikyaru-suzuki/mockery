// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// AsyncProducer is an autogenerated mock type for the AsyncProducer type
type AsyncProducer struct {
	mock.Mock
}

// Input provides a mock function with given fields:
func (_m *AsyncProducer) Input() chan<- bool {
	ret := _m.Called()

	var r0 chan<- bool
	if rf, ok := ret.Get(0).(func() chan<- bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- bool)
		}
	}

	return r0
}

// Output provides a mock function with given fields:
func (_m *AsyncProducer) Output() <-chan bool {
	ret := _m.Called()

	var r0 <-chan bool
	if rf, ok := ret.Get(0).(func() <-chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan bool)
		}
	}

	return r0
}

// Whatever provides a mock function with given fields:
func (_m *AsyncProducer) Whatever() chan bool {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	return r0
}

// NewAsyncProducer creates a new instance of AsyncProducer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewAsyncProducer(t testing.TB) *AsyncProducer {
	mock := &AsyncProducer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
