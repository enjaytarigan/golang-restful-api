// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import entity "brodo-demo/entity"
import mock "github.com/stretchr/testify/mock"

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// FindByUsername provides a mock function with given fields: username
func (_m *UserRepository) FindByUsername(username string) (entity.User, error) {
	ret := _m.Called(username)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: user
func (_m *UserRepository) Insert(user entity.User) (int, error) {
	ret := _m.Called(user)

	var r0 int
	if rf, ok := ret.Get(0).(func(entity.User) int); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyAvailableUsername provides a mock function with given fields: username
func (_m *UserRepository) VerifyAvailableUsername(username string) bool {
	ret := _m.Called(username)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}