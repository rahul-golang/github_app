// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/downstream_oauth_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	error "github.com/rahul-golang/github_app/error"
	models "github.com/rahul-golang/github_app/models"
	reflect "reflect"
)

// MockGithubOAuthDownStream is a mock of GithubOAuthDownStream interface
type MockGithubOAuthDownStream struct {
	ctrl     *gomock.Controller
	recorder *MockGithubOAuthDownStreamMockRecorder
}

// MockGithubOAuthDownStreamMockRecorder is the mock recorder for MockGithubOAuthDownStream
type MockGithubOAuthDownStreamMockRecorder struct {
	mock *MockGithubOAuthDownStream
}

// NewMockGithubOAuthDownStream creates a new mock instance
func NewMockGithubOAuthDownStream(ctrl *gomock.Controller) *MockGithubOAuthDownStream {
	mock := &MockGithubOAuthDownStream{ctrl: ctrl}
	mock.recorder = &MockGithubOAuthDownStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGithubOAuthDownStream) EXPECT() *MockGithubOAuthDownStreamMockRecorder {
	return m.recorder
}

// Redirect mocks base method
func (m *MockGithubOAuthDownStream) Redirect(ctx context.Context, exchangeReq models.TokenExchangeReq) (models.TokenExchangeResp, *error.APPError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Redirect", ctx, exchangeReq)
	ret0, _ := ret[0].(models.TokenExchangeResp)
	ret1, _ := ret[1].(*error.APPError)
	return ret0, ret1
}

// Redirect indicates an expected call of Redirect
func (mr *MockGithubOAuthDownStreamMockRecorder) Redirect(ctx, exchangeReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redirect", reflect.TypeOf((*MockGithubOAuthDownStream)(nil).Redirect), ctx, exchangeReq)
}