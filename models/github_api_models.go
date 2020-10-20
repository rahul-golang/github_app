package models

import "github.com/google/go-github/github"

type GetRepository struct {
	Owner      string `form:"owner" binding:"required"`
	Repository string `form:"repository" binding:"required"`
}

type CreatePullRequest struct {
	Repository          string `json:"repository" binding:"required"`
	Owner               string `json:"owner" binding:"required"`
	Title               string `json:"title,omitempty" binding:"required"`
	Head                string `json:"head,omitempty" binding:"required"`
	Base                string `json:"base,omitempty" binding:"required"`
	Body                string `json:"body,omitempty" binding:"required"`
	Issue               int    `json:"issue,omitempty" `
	MaintainerCanModify bool   `json:"maintainer_can_modify,omitempty" binding:"required"`
	Draft               bool   `json:"draft,omitempty"`
}

func (pr CreatePullRequest) MapToGithubPullRequest() *github.NewPullRequest {
	return &github.NewPullRequest{
		Title:               github.String(pr.Title),
		Head:                github.String(pr.Head),
		Base:                github.String(pr.Base),
		Body:                github.String(pr.Body),
		MaintainerCanModify: github.Bool(pr.MaintainerCanModify),
		Issue:               github.Int(pr.Issue),
		//Draft:               github.Bool(pr.Draft),
	}
}
