// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Cleanup is an autogenerated mock type for the Cleanup type
type Cleanup struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *Cleanup) Execute() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}