// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "ecommerce/domain/entity"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// OrderAppInterface is an autogenerated mock type for the OrderAppInterface type
type OrderAppInterface struct {
	mock.Mock
}

// AddOrder provides a mock function with given fields: _a0
func (_m *OrderAppInterface) AddOrder(_a0 *entity.Order) (*entity.Order, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Order
	if rf, ok := ret.Get(0).(func(*entity.Order) *entity.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Order) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddOrderDetail provides a mock function with given fields: _a0
func (_m *OrderAppInterface) AddOrderDetail(_a0 *entity.OrderDetail) (*entity.OrderDetail, error) {
	ret := _m.Called(_a0)

	var r0 *entity.OrderDetail
	if rf, ok := ret.Get(0).(func(*entity.OrderDetail) *entity.OrderDetail); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.OrderDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.OrderDetail) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByID provides a mock function with given fields: _a0
func (_m *OrderAppInterface) GetOrderByID(_a0 uuid.UUID) ([]entity.OrderDetail, error) {
	ret := _m.Called(_a0)

	var r0 []entity.OrderDetail
	if rf, ok := ret.Get(0).(func(uuid.UUID) []entity.OrderDetail); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.OrderDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOrderAppInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderAppInterface creates a new instance of OrderAppInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderAppInterface(t mockConstructorTestingTNewOrderAppInterface) *OrderAppInterface {
	mock := &OrderAppInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}