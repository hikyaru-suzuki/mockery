// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ExpecterTest is an autogenerated mock type for the ExpecterTest type
type ExpecterTest struct {
	mock.Mock
}

// ManyArgsReturns provides a mock function with given fields: str, i
func (_m *ExpecterTest) ManyArgsReturns(str string, i int) ([]string, error) {
	ret := _m.Called(str, i)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, int) []string); ok {
		r0 = rf(str, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(str, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoArg provides a mock function with given fields:
func (_m *ExpecterTest) NoArg() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NoReturn provides a mock function with given fields: str
func (_m *ExpecterTest) NoReturn(str string) {
	_m.Called(str)
}

// Variadic provides a mock function with given fields: ints
func (_m *ExpecterTest) Variadic(ints ...int) error {
	_va := make([]interface{}, len(ints))
	for _i := range ints {
		_va[_i] = ints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...int) error); ok {
		r0 = rf(ints...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VariadicMany provides a mock function with given fields: i, a, intfs
func (_m *ExpecterTest) VariadicMany(i int, a string, intfs ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, i, a)
	_ca = append(_ca, intfs...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, ...interface{}) error); ok {
		r0 = rf(i, a, intfs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewExpecterTest creates a new instance of ExpecterTest. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewExpecterTest(t testing.TB) *ExpecterTest {
	mock := &ExpecterTest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
