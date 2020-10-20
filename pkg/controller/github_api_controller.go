package controller

import (
	"github.com/gin-gonic/gin"
	appErr "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/models"
	"github.com/rahul-golang/github_app/pkg/service"
	logUtils "github.com/rahul-golang/github_app/utils/log_utils"
	"net/http"
)

type GithubApiController interface {
	GetRepository(ctx *gin.Context)
	CreatePullRequest(ctx *gin.Context)
}
type githubApiController struct {
	errResponseInterceptor appErr.ErrResponseInterceptor
	githubApiService       service.GithubApiService
}

func (ctrl githubApiController) GetRepository(ctx *gin.Context) {
	gCtx := ctx.Request.Context()
	getRepoRequest := models.GetRepository{}
	if err := ctx.ShouldBindQuery(&getRepoRequest); err != nil {
		logUtils.GetLogger(gCtx).Errorf("GithubOauthController.CreatePullRequest: Error in binding request body. Error: %+v ", err)
		ctrl.errResponseInterceptor.HandleBadRequest(ctx, err)
		return
	}

	repository, serviceError := ctrl.githubApiService.GetRepository(gCtx, getRepoRequest.Repository, getRepoRequest.Owner)
	if serviceError != nil {
		logUtils.GetLogger(gCtx).Errorf("GithubOauthController.GetRepository: Error in git oauth service. Error: %+v", serviceError)
		ctrl.errResponseInterceptor.HandleServiceError(ctx, serviceError)
		return
	}
	ctx.JSON(http.StatusOK, repository)
}

func (ctrl githubApiController) CreatePullRequest(ctx *gin.Context) {
	gCtx := ctx.Request.Context()
	logUtils.GetLogger(gCtx).Infof("GithubOauthController.CreatePullRequest: Create pull request started. ")
	var request models.CreatePullRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logUtils.GetLogger(gCtx).Errorf("GithubOauthController.CreatePullRequest: Error in binding request body. Error: %+v ", err)
		ctrl.errResponseInterceptor.HandleBadRequest(ctx, err)
		return
	}
	repository, serviceError := ctrl.githubApiService.CreatePullRequest(ctx,request)
	if serviceError != nil {
		logUtils.GetLogger(gCtx).Errorf("GithubOauthController.CreatePullRequest: Error in git oauth service. Error: %+v", serviceError)
		ctrl.errResponseInterceptor.HandleServiceError(ctx, serviceError)
		return
	}
	logUtils.GetLogger(gCtx).Infof("GithubOauthController.CreatePullRequest: PR created successfully . ")

	ctx.JSON(http.StatusOK, repository)
}

func NewGithubApiController(errResponseInterceptor appErr.ErrResponseInterceptor, githubApiService service.GithubApiService) GithubApiController {
	return githubApiController{
		errResponseInterceptor: errResponseInterceptor,
		githubApiService:       githubApiService,
	}
}
