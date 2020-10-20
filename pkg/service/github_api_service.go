package service

import (
	"context"
	appErr "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/models"
	logUtils "github.com/rahul-golang/github_app/utils/log_utils"
)

type GithubApiService interface {
	GetRepository(ctx context.Context, repository, owner string) (interface{}, *appErr.APPError)
	CreatePullRequest(ctx context.Context,repository models.CreatePullRequest) (string, *appErr.APPError)
}
type githubApiService struct {
	gitApiDownstream GithubApiDownstream
}

func (g githubApiService) GetRepository(ctx context.Context, repository, owner string) (interface{}, *appErr.APPError) {
	logUtils.GetLogger(ctx).Info("GithubApiService.GetRepository: Accessing github repository")
	repositoryResp, downstreamError := g.gitApiDownstream.AccessGitRepository(ctx, repository, owner)
	if downstreamError != nil {
		logUtils.GetLogger(ctx).Errorf("GithubApiService.GetRepository: Error in downstream. Error: %+v", downstreamError)
		return nil, downstreamError
	}
	logUtils.GetLogger(ctx).Info("GithubApiService.GetRepository: Success.")
	return repositoryResp, nil
}
func (g githubApiService) CreateBranch(ctx context.Context) (interface{}, *appErr.APPError) {
	repository, downstreamError := g.gitApiDownstream.CreateBranch(ctx)
	if downstreamError != nil {
		logUtils.GetLogger(ctx).Errorf("Error: %+v", downstreamError)
		return nil, downstreamError
	}
	return repository, nil
}

func (g githubApiService) CreatePullRequest(ctx context.Context, request models.CreatePullRequest) (string, *appErr.APPError) {
	logUtils.GetLogger(ctx).Infof("githubApiService.CreatePullRequest: Request: %+v", request)
	prHtmlUrl, downstreamError := g.gitApiDownstream.CreatePullRequest(ctx, request)
	if downstreamError != nil {
		logUtils.GetLogger(ctx).Errorf("githubApiService.CreatePullRequest. Error in downstream while creating pull request. Error: %+v", downstreamError)
		return "", downstreamError
	}
	logUtils.GetLogger(ctx).Info("githubApiService.CreatePullRequest: Success pr creation")
	return prHtmlUrl, nil
}

func NewGithubApiService(gitApiDownstream GithubApiDownstream) GithubApiService {
	return githubApiService{
		gitApiDownstream: gitApiDownstream,
	}
}
