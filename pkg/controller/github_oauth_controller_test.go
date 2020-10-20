package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	appErr "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/mocks"
	"github.com/rahul-golang/github_app/models"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GithubOauthControllerSuite struct {
	suite.Suite
	mockCtrl               *gomock.Controller
	recorder               *httptest.ResponseRecorder
	context                *gin.Context
	errResponseInterceptor *mocks.MockErrResponseInterceptor
	oauthService           *mocks.MockGithubOauthService
	oauthController        GithubOauthController
}

func TestGithubOauthControllerSuite(t *testing.T) {
	suite.Run(t, new(GithubOauthControllerSuite))
}

func (suite *GithubOauthControllerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.errResponseInterceptor = mocks.NewMockErrResponseInterceptor(suite.mockCtrl)
	suite.oauthService = mocks.NewMockGithubOauthService(suite.mockCtrl)
	suite.oauthController = NewOAuthController(suite.errResponseInterceptor, suite.oauthService)
}

func (suite GithubOauthControllerSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite GithubOauthControllerSuite) TestGetCity_shouldReturnBadRequestErrorWhenRequestDataIsInvalid() {
	suite.context.Request, _ = http.NewRequest("GET", "get-city?test=test", nil)
	suite.errResponseInterceptor.EXPECT().HandleBadRequest(suite.context, errors.New("invalid request"))
	suite.oauthController.Redirect(suite.context)

}

func (suite GithubOauthControllerSuite) TestGetCity_ShouldReturnSameErrorWhenServiceReturnError() {
	request := models.RedirectResponse{
		Code: "code-test",
	}
	serviceError := appErr.InternalServerErrorFunc("something went wrong")
	//data := url.Values{}
	//data.Set("code", "code-test")
	suite.context.Request, _ = http.NewRequest("POST", "/oauth/redirect?code=code-test", nil)
	suite.oauthService.EXPECT().Redirect(gomock.Any(), request).Return(models.TokenExchangeResp{}, serviceError)
	suite.errResponseInterceptor.EXPECT().HandleServiceError(suite.context, serviceError)
	suite.oauthController.Redirect(suite.context)

}

func (suite GithubOauthControllerSuite) TestGetCity_ShouldReturnSuccessResponse() {
	request := models.RedirectResponse{Code: "code-test"}
	resp := models.TokenExchangeResp{}
	suite.context.Request, _ = http.NewRequest("", "/oauth/redirect?code=code-test", nil)
	suite.oauthService.EXPECT().Redirect(gomock.Any(), request).Return(resp, nil)
	suite.oauthController.Redirect(suite.context)
	bodyBytes, _ := ioutil.ReadAll(suite.recorder.Body)
	json.Unmarshal(bodyBytes, &resp)
	suite.Equal(models.TokenExchangeResp{}, resp)

}
