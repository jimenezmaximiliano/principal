// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LoggerForPanic is an autogenerated mock type for the LoggerForPanic type
type LoggerForPanic struct {
	mock.Mock
}

// Panic provides a mock function with given fields: err
func (_m *LoggerForPanic) Panic(err error) {
	_m.Called(err)
}

type mockConstructorTestingTNewLoggerForPanic interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoggerForPanic creates a new instance of LoggerForPanic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoggerForPanic(t mockConstructorTestingTNewLoggerForPanic) *LoggerForPanic {
	mock := &LoggerForPanic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
