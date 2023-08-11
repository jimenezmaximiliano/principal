// Code generated by mockery v2.32.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// LoggerForPanic is an autogenerated mock type for the LoggerForPanic type
type LoggerForPanic struct {
	mock.Mock
}

// Panic provides a mock function with given fields: ctx, err
func (_m *LoggerForPanic) Panic(ctx context.Context, err error) {
	_m.Called(ctx, err)
}

// NewLoggerForPanic creates a new instance of LoggerForPanic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoggerForPanic(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoggerForPanic {
	mock := &LoggerForPanic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
