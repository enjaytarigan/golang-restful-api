// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PasswordHash is an autogenerated mock type for the PasswordHash type
type PasswordHash struct {
	mock.Mock
}

// Hash provides a mock function with given fields: password
func (_m *PasswordHash) Hash(password string) (string, error) {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsMatch provides a mock function with given fields: password, hashedPassword
func (_m *PasswordHash) IsMatch(password string, hashedPassword string) bool {
	ret := _m.Called(password, hashedPassword)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(password, hashedPassword)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
