package service

import (
	"context"
	"github.com/google/go-github/github"
	appErr "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/models"
	logUtils "github.com/rahul-golang/github_app/utils/log_utils"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type GithubApiDownstream interface {
	AccessGitRepository(ctx context.Context, repository, owner string) (interface{}, *appErr.APPError)
	CreateBranch(ctx context.Context) (interface{}, *appErr.APPError)
	CreatePullRequest(ctx context.Context, request models.CreatePullRequest) (string, *appErr.APPError)
}

type githubApiDownStream struct {
}

func (githubApiDownStream githubApiDownStream) CreatePullRequest(ctx context.Context, request models.CreatePullRequest) (string, *appErr.APPError) {
	logUtils.GetLogger(ctx).Infof("GithubApiDownStream.CreatePullRequest: Creating api call create pull request: %+v", request)
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOCKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	pr, _, err := client.PullRequests.Create(context.Background(), request.Owner, request.Repository, request.MapToGithubPullRequest())
	if err != nil {
		logUtils.GetLogger(ctx).Errorf("GithubApiDownStream.CreatePullRequest: Error in creating pull request. Error: %+v", err)

		if strings.Contains(err.Error(), "401 Bad credentials") {
			return "", appErr.ErrNotAuthorized
		} else {
			return "", appErr.InternalServerErrorFunc(err.Error())
		}
	}

	logUtils.GetLogger(ctx).Infof("GithubApiDownStream.CreatePullRequest: PR created successfully: %s\n", pr.GetHTMLURL())
	return pr.GetHTMLURL(), nil
}

func NewGithubApiDownStream() GithubApiDownstream {
	return githubApiDownStream{}
}

func (githubApiDownStream githubApiDownStream) AccessGitRepository(ctx context.Context, repository, owner string) (interface{}, *appErr.APPError) {
	logUtils.GetLogger(ctx).Info("GithubApiDownStream.AccessGitRepository: Accessing github repository.")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOCKEN")},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repos, _, err := client.Repositories.Get(ctx, owner, repository)
	if err != nil {
		logUtils.GetLogger(ctx).Errorf("GithubApiDownStream.AccessGitRepository: Error in accessing github repository. Error: %+v", err)
		if strings.Contains(err.Error(), "401 Bad credentials") {
			return "", appErr.ErrNotAuthorized
		} else {
			return "", appErr.InternalServerErrorFunc(err.Error())
		}
	}
	logUtils.GetLogger(ctx).Infof("GithubApiDownStream.AccessGitRepository: Successfully accessed repository. Error: %+v", repos)
	return repos, nil
}

func (githubApiDownStream githubApiDownStream) CreateBranch(ctx context.Context) (interface{}, *appErr.APPError) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "81bf44f20dfd11a94429dc348cf65264983bb191"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.Get(ctx, "users", "ecommerce")
	if err != nil {
		logUtils.GetLogger(context.Background()).Errorf("%+v", err)
		return nil, appErr.InternalServerErrorFunc(err.Error())
	}
	logUtils.GetLogger(context.Background()).Infof("%+v", repos)
	return repos, nil
}
