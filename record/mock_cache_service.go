// Code generated by mockery v1.0.0. DO NOT EDIT.

package record

import (
	"time"

	mock "github.com/stretchr/testify/mock"
)

// mockCacheService is an autogenerated mock type for the CacheService type
type mockCacheService struct {
	mock.Mock
}

// InsertUserTask provides a mock function with given fields: key, value, expire
func (_m *mockCacheService) SetExpire(key string, value int, expire time.Duration) error {
	ret := _m.Called(key, value, expire)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, time.Duration) error); ok {
		r0 = rf(key, value, expire)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInt provides a mock function with given fields: key
func (_m *mockCacheService) GetInt(key string) (int, error) {
	ret := _m.Called(key)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}