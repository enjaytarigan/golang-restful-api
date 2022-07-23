// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AuthenticationTokenManager is an autogenerated mock type for the AuthenticationTokenManager type
type AuthenticationTokenManager struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: userId
func (_m *AuthenticationTokenManager) CreateAccessToken(userId int) (string, error) {
	ret := _m.Called(userId)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyAccessToken provides a mock function with given fields: token
func (_m *AuthenticationTokenManager) VerifyAccessToken(token string) (int, error) {
	ret := _m.Called(token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}