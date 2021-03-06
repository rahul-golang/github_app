// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/github_api_downstream.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	error "github.com/rahul-golang/github_app/error"
	models "github.com/rahul-golang/github_app/models"
	reflect "reflect"
)

// MockGithubApiDownstream is a mock of GithubApiDownstream interface
type MockGithubApiDownstream struct {
	ctrl     *gomock.Controller
	recorder *MockGithubApiDownstreamMockRecorder
}

// MockGithubApiDownstreamMockRecorder is the mock recorder for MockGithubApiDownstream
type MockGithubApiDownstreamMockRecorder struct {
	mock *MockGithubApiDownstream
}

// NewMockGithubApiDownstream creates a new mock instance
func NewMockGithubApiDownstream(ctrl *gomock.Controller) *MockGithubApiDownstream {
	mock := &MockGithubApiDownstream{ctrl: ctrl}
	mock.recorder = &MockGithubApiDownstreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGithubApiDownstream) EXPECT() *MockGithubApiDownstreamMockRecorder {
	return m.recorder
}

// AccessGitRepository mocks base method
func (m *MockGithubApiDownstream) AccessGitRepository(ctx context.Context, repository, owner string) (interface{}, *error.APPError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessGitRepository", ctx, repository, owner)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*error.APPError)
	return ret0, ret1
}

// AccessGitRepository indicates an expected call of AccessGitRepository
func (mr *MockGithubApiDownstreamMockRecorder) AccessGitRepository(ctx, repository, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessGitRepository", reflect.TypeOf((*MockGithubApiDownstream)(nil).AccessGitRepository), ctx, repository, owner)
}

// CreateBranch mocks base method
func (m *MockGithubApiDownstream) CreateBranch(ctx context.Context) (interface{}, *error.APPError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBranch", ctx)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*error.APPError)
	return ret0, ret1
}

// CreateBranch indicates an expected call of CreateBranch
func (mr *MockGithubApiDownstreamMockRecorder) CreateBranch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBranch", reflect.TypeOf((*MockGithubApiDownstream)(nil).CreateBranch), ctx)
}

// CreatePullRequest mocks base method
func (m *MockGithubApiDownstream) CreatePullRequest(ctx context.Context, request models.CreatePullRequest) (string, *error.APPError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePullRequest", ctx, request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*error.APPError)
	return ret0, ret1
}

// CreatePullRequest indicates an expected call of CreatePullRequest
func (mr *MockGithubApiDownstreamMockRecorder) CreatePullRequest(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePullRequest", reflect.TypeOf((*MockGithubApiDownstream)(nil).CreatePullRequest), ctx, request)
}
