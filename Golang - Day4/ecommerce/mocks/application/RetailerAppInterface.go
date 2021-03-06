// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "ecommerce/domain/entity"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// RetailerAppInterface is an autogenerated mock type for the RetailerAppInterface type
type RetailerAppInterface struct {
	mock.Mock
}

// AddRetailer provides a mock function with given fields: _a0
func (_m *RetailerAppInterface) AddRetailer(_a0 *entity.Retailer) (*entity.Retailer, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Retailer
	if rf, ok := ret.Get(0).(func(*entity.Retailer) *entity.Retailer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Retailer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Retailer) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRetailerByID provides a mock function with given fields: _a0
func (_m *RetailerAppInterface) GetRetailerByID(_a0 uuid.UUID) (*entity.Retailer, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Retailer
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entity.Retailer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Retailer)
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

type mockConstructorTestingTNewRetailerAppInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewRetailerAppInterface creates a new instance of RetailerAppInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRetailerAppInterface(t mockConstructorTestingTNewRetailerAppInterface) *RetailerAppInterface {
	mock := &RetailerAppInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
