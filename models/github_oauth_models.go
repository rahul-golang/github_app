package models

import (
	"fmt"
	"os"
)

type RedirectResponse struct {
	Code string `form:"code" binding:"required"`
}

type TokenExchangeResp struct {
	AccessToken string `json:"access_token"`
}

type TokenExchangeReq struct {
	Url          string
	ClientID     string
	ClientSecret string
	Code         string
}

func NewTokenExchangeReq(url, code string) TokenExchangeReq {
	return TokenExchangeReq{
		Url:          url,
		Code:         code,
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	}
}

func (tokenExchangeReq TokenExchangeReq) GetExchangeTokenUrl() string {
	return fmt.Sprintf(tokenExchangeReq.Url+"?client_id=%s&client_secret=%s&code=%s", tokenExchangeReq.ClientID, tokenExchangeReq.ClientSecret, tokenExchangeReq.Code)
}
