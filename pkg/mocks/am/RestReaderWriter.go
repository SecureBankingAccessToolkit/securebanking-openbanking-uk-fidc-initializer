// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RestReaderWriter is an autogenerated mock type for the RestReaderWriter type
type RestReaderWriter struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *RestReaderWriter) Get(_a0 string, _a1 map[string]string) ([]byte, int) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, map[string]string) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0, 200
}

func (_m *RestReaderWriter) GetRS(_a0 string, _a1 map[string]string) ([]byte, int) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, map[string]string) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0, 200
}

// Patch provides a mock function with given fields: _a0, _a1, _a2
func (_m *RestReaderWriter) Patch(_a0 string, _a1 interface{}, _a2 map[string]string) int {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, interface{}, map[string]string) int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Post provides a mock function with given fields: _a0, _a1, _a2
func (_m *RestReaderWriter) Post(_a0 string, _a1 interface{}, _a2 map[string]string) ([]byte, int) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, interface{}, map[string]string) []byte); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).([]byte)
	}

	return r0, 200
}

func (_m *RestReaderWriter) PostRS(_a0 string, _a2 map[string]string) int {
	ret := _m.Called(_a0, _a2)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, map[string]string) int); ok {
		r0 = rf(_a0, _a2)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Put provides a mock function with given fields: _a0, _a1, _a2
func (_m *RestReaderWriter) Put(_a0 string, _a1 interface{}, _a2 map[string]string) int {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, interface{}, map[string]string) int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
