package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	appErr "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/mocks"
	"github.com/rahul-golang/github_app/models"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GithubApiControllerSuite struct {
	suite.Suite
	mockCtrl               *gomock.Controller
	recorder               *httptest.ResponseRecorder
	context                *gin.Context
	errResponseInterceptor *mocks.MockErrResponseInterceptor
	apiService             *mocks.MockGithubApiService
	apiController          GithubApiController
}

func TestGithubApiControllerSuite(t *testing.T) {
	suite.Run(t, new(GithubApiControllerSuite))
}

func (suite *GithubApiControllerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.errResponseInterceptor = mocks.NewMockErrResponseInterceptor(suite.mockCtrl)
	suite.apiService = mocks.NewMockGithubApiService(suite.mockCtrl)
	suite.apiController = NewGithubApiController(suite.errResponseInterceptor, suite.apiService)
}

func (suite GithubApiControllerSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite GithubApiControllerSuite) TestGetRepository_shouldReturnBadRequestErrorWhenRequestDataIsInvalid() {
	suite.context.Request, _ = http.NewRequest("POST", "get-city?test=test", nil)
	suite.errResponseInterceptor.EXPECT().HandleBadRequest(gomock.Any(), gomock.Any())
	suite.apiController.GetRepository(suite.context)

}

func (suite GithubApiControllerSuite) TestGetRepository_ShouldReturnSameErrorWhenServiceReturnError() {
	request := models.GetRepository{
		Owner:      "test-owner",
		Repository: "test-repository",
	}
	serviceError := appErr.InternalServerErrorFunc("something went wrong")
	suite.context.Request, _ = http.NewRequest("GET", "/test?owner=test-owner&repository=test-repository", nil)
	suite.apiService.EXPECT().GetRepository(gomock.Any(), request.Repository, request.Owner).Return(models.TokenExchangeResp{}, serviceError)
	suite.errResponseInterceptor.EXPECT().HandleServiceError(suite.context, serviceError)
	suite.apiController.GetRepository(suite.context)

}

func (suite GithubApiControllerSuite) TestGetRepository_ShouldReturnSuccessResponse() {
	request := models.GetRepository{
		Owner:      "test-owner",
		Repository: "test-repository",
	}
	resp := models.TokenExchangeResp{}
	suite.context.Request, _ = http.NewRequest("GET", "/test?owner=test-owner&repository=test-repository", nil)
	suite.apiService.EXPECT().GetRepository(gomock.Any(), request.Repository,request.Owner).Return(resp, nil)
	suite.apiController.GetRepository(suite.context)
	bodyBytes, _ := ioutil.ReadAll(suite.recorder.Body)
	json.Unmarshal(bodyBytes, &resp)
	suite.Equal(models.TokenExchangeResp{}, resp)

}


func (suite GithubApiControllerSuite) TestCreatePullRequest_shouldReturnBadRequestErrorWhenRequestDataIsInvalid() {
	suite.context.Request, _ = http.NewRequest("POST", "create-pull", nil)
	suite.errResponseInterceptor.EXPECT().HandleBadRequest(gomock.Any(), gomock.Any())
	suite.apiController.CreatePullRequest(suite.context)

}

func (suite GithubApiControllerSuite) TestCreatePullRequest_ShouldReturnSameErrorWhenServiceReturnError() {
	request := models.CreatePullRequest{
		Owner:      "test-owner",
		Repository: "test-repository",
		Body: "test-body",
		MaintainerCanModify: true,
		Base: "test",
		Head: "test",
		 Title: "test",
	}
	serviceError := appErr.InternalServerErrorFunc("something went wrong")
	suite.context.Request, _ = http.NewRequest("GET", "/test?owner=test-owner&repository=test-repository", getPullRequest(request))
	suite.apiService.EXPECT().CreatePullRequest(gomock.Any(),request).Return("", serviceError)
	suite.errResponseInterceptor.EXPECT().HandleServiceError(suite.context, serviceError)
	suite.apiController.CreatePullRequest(suite.context)

}

func (suite GithubApiControllerSuite) TestCreatePullRequest_ShouldReturnSuccessResponse() {
	request := models.CreatePullRequest{
		Owner:      "test-owner",
		Repository: "test-repository",
		Body: "test-body",
		MaintainerCanModify: true,
		Base: "test",
		Head: "test",
		Title: "test",
	}
	expectredResp := "test-resp"
	actualresp:=""
	suite.context.Request, _ = http.NewRequest("GET", "/test?owner=test-owner&repository=test-repository", getPullRequest(request))
	suite.apiService.EXPECT().CreatePullRequest(gomock.Any(), request).Return(expectredResp, nil)
	suite.apiController.CreatePullRequest(suite.context)
	bodyBytes, _ := ioutil.ReadAll(suite.recorder.Body)
	json.Unmarshal(bodyBytes, &actualresp)
	suite.Equal(expectredResp, actualresp)

}

func getPullRequest(value interface{}) io.Reader {
	byteData, _ := json.Marshal(value)
	return bytes.NewReader(byteData)
}