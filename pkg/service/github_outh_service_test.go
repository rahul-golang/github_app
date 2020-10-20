package service

import (
	"context"
	"github.com/golang/mock/gomock"
	appErr "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/mocks"
	"github.com/rahul-golang/github_app/models"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type GithubOauthServiceTestSuite struct {
	suite.Suite
	recorder              *httptest.ResponseRecorder
	mockCtrl              *gomock.Controller
	context               context.Context
	githubOauthDownstream *mocks.MockGithubOAuthDownStream
	githubConfig          *mocks.MockGithubConfigs
	githubOauthService    GithubOauthService
}

func TestGithubOauthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GithubOauthServiceTestSuite))
}

func (suite *GithubOauthServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context = context.Background()
	suite.githubOauthDownstream = mocks.NewMockGithubOAuthDownStream(suite.mockCtrl)
	suite.githubConfig = mocks.NewMockGithubConfigs(suite.mockCtrl)
	suite.githubOauthService = NewGithubOauthService(suite.githubConfig, suite.githubOauthDownstream)

}

func (suite GithubOauthServiceTestSuite) TestRedirect_ShouldReturnAnErrorIfDownstreamServiceReturnsAnError() {
	suite.githubConfig.EXPECT().GetExchangeTokenURL().Return("test://test.com")
	expectedErr := appErr.InternalServerErrorFunc("something-went-wrong")
	suite.githubOauthDownstream.EXPECT().Redirect(suite.context, gomock.Any()).Return(models.TokenExchangeResp{}, expectedErr)
	resp, actualErr := suite.githubOauthService.Redirect(suite.context, models.RedirectResponse{Code: "test"})
	suite.Equal(models.TokenExchangeResp{}, resp)
	suite.Equal(expectedErr, actualErr)
}

func (suite GithubOauthServiceTestSuite) TestRedirect_ShouldSucess() {
	suite.githubConfig.EXPECT().GetExchangeTokenURL().Return("test://test.com")
	suite.githubOauthDownstream.EXPECT().Redirect(suite.context, gomock.Any()).Return(models.TokenExchangeResp{AccessToken: "test"}, nil)
	resp, actualErr := suite.githubOauthService.Redirect(suite.context, models.RedirectResponse{Code: "test"})
	suite.Equal(models.TokenExchangeResp{AccessToken: "test"}, resp)
	suite.Nil(actualErr)
}
