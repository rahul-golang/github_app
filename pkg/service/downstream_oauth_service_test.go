package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type GithubOAuthDownStreamTestSuite struct {
	suite.Suite
	recorder         *httptest.ResponseRecorder
	context          context.Context
	githubDownStream GithubOAuthDownStream
}

func TestGithubOAuthDownStreamTestSuite(t *testing.T) {
	suite.Run(t, new(GithubOAuthDownStreamTestSuite))
}

func (suite *GithubOAuthDownStreamTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.context = context.Background()
	suite.githubDownStream = NewGithubDownStream()
}