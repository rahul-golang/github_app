// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/github_oauth_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	error "github.com/rahul-golang/github_app/error"
	models "github.com/rahul-golang/github_app/models"
	reflect "reflect"
)

// MockGithubOauthService is a mock of GithubOauthService interface
type MockGithubOauthService struct {
	ctrl     *gomock.Controller
	recorder *MockGithubOauthServiceMockRecorder
}

// MockGithubOauthServiceMockRecorder is the mock recorder for MockGithubOauthService
type MockGithubOauthServiceMockRecorder struct {
	mock *MockGithubOauthService
}

// NewMockGithubOauthService creates a new mock instance
func NewMockGithubOauthService(ctrl *gomock.Controller) *MockGithubOauthService {
	mock := &MockGithubOauthService{ctrl: ctrl}
	mock.recorder = &MockGithubOauthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGithubOauthService) EXPECT() *MockGithubOauthServiceMockRecorder {
	return m.recorder
}

// Redirect mocks base method
func (m *MockGithubOauthService) Redirect(ctx context.Context, code models.RedirectResponse) (models.TokenExchangeResp, *error.APPError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Redirect", ctx, code)
	ret0, _ := ret[0].(models.TokenExchangeResp)
	ret1, _ := ret[1].(*error.APPError)
	return ret0, ret1
}

// Redirect indicates an expected call of Redirect
func (mr *MockGithubOauthServiceMockRecorder) Redirect(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redirect", reflect.TypeOf((*MockGithubOauthService)(nil).Redirect), ctx, code)
}
