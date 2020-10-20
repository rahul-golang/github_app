package service

import (
	"context"
	"encoding/json"
	error2 "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/models"
	logUtils "github.com/rahul-golang/github_app/utils/log_utils"
	"net/http"
	"os"
)

type GithubOAuthDownStream interface {
	Redirect(ctx context.Context,exchangeReq models.TokenExchangeReq) (models.TokenExchangeResp ,*error2.APPError)
}

type githubOAuthDownStream struct {
	httpClient http.Client
}
func NewGithubDownStream() GithubOAuthDownStream {
	return githubOAuthDownStream{
		httpClient: http.Client{},
	}
}

func (githubDownStream githubOAuthDownStream) Redirect(ctx context.Context,exchangeReq models.TokenExchangeReq) (models.TokenExchangeResp ,*error2.APPError) {
	logUtils.GetLogger(ctx).Infof("GithubOAuthDownStream.Redirect: creating a exchange token request: %+v", exchangeReq)
	req, err := http.NewRequest(http.MethodPost, exchangeReq.GetExchangeTokenUrl(), nil)
	if err != nil {
		logUtils.GetLogger(ctx).Errorf("GithubOAuthDownStream.Redirect: could not create HTTP request: %v", err)
		return models.TokenExchangeResp{},error2.BadRequestErrorFunc(err.Error())
	}

	req.Header.Set("accept", "application/json")
	res, err := githubDownStream.httpClient.Do(req)
	if err != nil {
		logUtils.GetLogger(ctx).Errorf("GithubOAuthDownStream.Redirect: could not send HTTP request: %v", err)
		return models.TokenExchangeResp{}, error2.InternalServerErrorFunc(err.Error())
	}
	defer res.Body.Close()
	logUtils.GetLogger(ctx).Infof("%+v", res.Body)

	var resp models.TokenExchangeResp
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		logUtils.GetLogger(ctx).Errorf("GithubOAuthDownStream.Redirect: could not parse JSON response: %v", err)
		return models.TokenExchangeResp{}, error2.InternalServerErrorFunc(err.Error())
	}
	os.Setenv("ACCESS_TOCKEN",resp.AccessToken)
	logUtils.GetLogger(ctx).Infof("GithubOAuthDownStream.Redirect: Successfully a exchange token request: %+v",resp )
	return resp, nil
}
