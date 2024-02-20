// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	entity "git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	mock "github.com/stretchr/testify/mock"
)

// HelperInf is an autogenerated mock type for the HelperInf type
type HelperInf struct {
	mock.Mock
}

// CheckPassword provides a mock function with given fields: pwd, hash
func (_m *HelperInf) CheckPassword(pwd string, hash []byte) (bool, error) {
	ret := _m.Called(pwd, hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, []byte) bool); ok {
		r0 = rf(pwd, hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(pwd, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAndSign provides a mock function with given fields: data
func (_m *HelperInf) CreateAndSign(data entity.User) (string, error) {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func(entity.User) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HashPassword provides a mock function with given fields: pwd, cost
func (_m *HelperInf) HashPassword(pwd string, cost int) ([]byte, error) {
	ret := _m.Called(pwd, cost)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, int) []byte); ok {
		r0 = rf(pwd, cost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(pwd, cost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
