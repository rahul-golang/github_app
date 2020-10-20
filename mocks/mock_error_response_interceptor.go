// Code generated by MockGen. DO NOT EDIT.
// Source: error/error_response_interceptor.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	appError "github.com/rahul-golang/github_app/error"
	reflect "reflect"
)

// MockErrResponseInterceptor is a mock of ErrResponseInterceptor interface
type MockErrResponseInterceptor struct {
	ctrl     *gomock.Controller
	recorder *MockErrResponseInterceptorMockRecorder
}

// MockErrResponseInterceptorMockRecorder is the mock recorder for MockErrResponseInterceptor
type MockErrResponseInterceptorMockRecorder struct {
	mock *MockErrResponseInterceptor
}

// NewMockErrResponseInterceptor creates a new mock instance
func NewMockErrResponseInterceptor(ctrl *gomock.Controller) *MockErrResponseInterceptor {
	mock := &MockErrResponseInterceptor{ctrl: ctrl}
	mock.recorder = &MockErrResponseInterceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockErrResponseInterceptor) EXPECT() *MockErrResponseInterceptorMockRecorder {
	return m.recorder
}

// HandleBadRequest mocks base method
func (m *MockErrResponseInterceptor) HandleBadRequest(ctx *gin.Context, bindErr error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleBadRequest", ctx, bindErr)
}

// HandleBadRequest indicates an expected call of HandleBadRequest
func (mr *MockErrResponseInterceptorMockRecorder) HandleBadRequest(ctx, bindErr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBadRequest", reflect.TypeOf((*MockErrResponseInterceptor)(nil).HandleBadRequest), ctx, bindErr)
}

// HandleServiceError mocks base method
func (m *MockErrResponseInterceptor) HandleServiceError(ctx *gin.Context, serviceErr *appError.APPError) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleServiceError", ctx, serviceErr)
}

// HandleServiceError indicates an expected call of HandleServiceError
func (mr *MockErrResponseInterceptorMockRecorder) HandleServiceError(ctx, serviceErr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleServiceError", reflect.TypeOf((*MockErrResponseInterceptor)(nil).HandleServiceError), ctx, serviceErr)
}
