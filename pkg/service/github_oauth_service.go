package service

import (
	"context"
	"github.com/rahul-golang/github_app/configurations"
	error2 "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/models"
	logUtils "github.com/rahul-golang/github_app/utils/log_utils"
)

type GithubOauthService interface {
	Redirect(ctx context.Context, code models.RedirectResponse) (models.TokenExchangeResp, *error2.APPError)
}

type githubOauthService struct {
	gitConfig        configurations.GithubConfigs
	githubDownStream GithubOAuthDownStream
}

func NewGithubOauthService(gitConfig configurations.GithubConfigs, githubDownStream GithubOAuthDownStream) GithubOauthService {
	return githubOauthService{
		gitConfig:        gitConfig,
		githubDownStream: githubDownStream,
	}
}

func (githubOauthService githubOauthService) Redirect(ctx context.Context, redirectResponse models.RedirectResponse) (models.TokenExchangeResp, *error2.APPError) {
	logUtils.GetLogger(ctx).Infof("GithubOauthService.Redirect")
	exchangeTokenUrl := githubOauthService.gitConfig.GetExchangeTokenURL()
	resp, downstreamError := githubOauthService.githubDownStream.Redirect(ctx,models.NewTokenExchangeReq(exchangeTokenUrl, redirectResponse.Code))
	if downstreamError != nil {
		logUtils.GetLogger(ctx).Errorf("GithubOauthService.Redirect: Error by oauth downstream. Error: %+v", downstreamError)
		return models.TokenExchangeResp{}, downstreamError
	}
	logUtils.GetLogger(ctx).Infof("GithubOauthService.Redirect: Token Exchange success.")
	return resp, nil
}
