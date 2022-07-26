// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import entity "brodo-demo/entity"
import mock "github.com/stretchr/testify/mock"
import repository "brodo-demo/repository"

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// FindAllAndCount provides a mock function with given fields: params
func (_m *ProductRepository) FindAllAndCount(params repository.FindAllProductsParam) ([]entity.Product, int, error) {
	ret := _m.Called(params)

	var r0 []entity.Product
	if rf, ok := ret.Get(0).(func(repository.FindAllProductsParam) []entity.Product); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(repository.FindAllProductsParam) int); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(repository.FindAllProductsParam) error); ok {
		r2 = rf(params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindById provides a mock function with given fields: productId
func (_m *ProductRepository) FindById(productId int) (entity.Product, error) {
	ret := _m.Called(productId)

	var r0 entity.Product
	if rf, ok := ret.Get(0).(func(int) entity.Product); ok {
		r0 = rf(productId)
	} else {
		r0 = ret.Get(0).(entity.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOne provides a mock function with given fields: product
func (_m *ProductRepository) InsertOne(product entity.Product) (*entity.Product, error) {
	ret := _m.Called(product)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(entity.Product) *entity.Product); ok {
		r0 = rf(product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Product) error); ok {
		r1 = rf(product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
