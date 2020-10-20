package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	appErr "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/models"
	"github.com/rahul-golang/github_app/pkg/service"
	"github.com/rahul-golang/github_app/utils/log_utils"

	"net/http"
)

type GithubOauthController interface {
	Redirect(ctx *gin.Context)
}
type githubOauthController struct {
	errResponseInterceptor appErr.ErrResponseInterceptor
	githubOauthService     service.GithubOauthService
}

func (ctrl githubOauthController) Redirect(ctx *gin.Context) {
	log_utils.GetLogger(ctx).Infof("githubOauthController.Redirect: Inside a callback request")
	gCtx := ctx.Request.Context()
	var redirectParam models.RedirectResponse
	if err := ctx.ShouldBindQuery(&redirectParam); err != nil {
		log_utils.GetLogger(gCtx).Errorf("GithubOauthController.Redirect: Error in form params. Error: %+v", err)
		ctrl.errResponseInterceptor.HandleBadRequest(ctx, errors.New("invalid request"))
		return
	}
	fmt.Println(redirectParam,gCtx)
	resp, serviceError := ctrl.githubOauthService.Redirect(gCtx, redirectParam)
	if serviceError != nil {
		log_utils.GetLogger(gCtx).Errorf("GithubOauthController.Redirect: Error in git oauth service. Error: %+v", serviceError)
		ctrl.errResponseInterceptor.HandleServiceError(ctx, serviceError)
		return
	}
	log_utils.GetLogger(gCtx).Infof("GithubOauthController.Redirect: Token :%+v", resp)
	ctx.Writer.Header().Set("Location", "/welcome.html?access_token="+resp.AccessToken)
	ctx.AbortWithStatus(http.StatusFound)

}

func NewOAuthController(errInterceptor appErr.ErrResponseInterceptor, githubOauthService service.GithubOauthService) GithubOauthController {
	return &githubOauthController{
		errResponseInterceptor: errInterceptor,
		githubOauthService:     githubOauthService,
	}
}
