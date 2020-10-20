package configurations

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type GithubConfigsSuite struct {
	suite.Suite
	gitConfigs GithubConfigs
}

func TestGithubConfigsSuite(t *testing.T) {
	suite.Run(t, new(GithubConfigsSuite))
}

func (suite *GithubConfigsSuite) SetupTest() {
	gitConf := GithubConfig{
		BaseUrl: "test://test.com",
	}
	suite.gitConfigs = NewGithubConfigs(gitConf)
}

func (suite GithubConfigsSuite) TearDownTest() {
}

func(suite GithubConfigsSuite)TestGetBaseURL_ShouldReturnBaseUrl(){
	expect:="test://test.com"
	actual:=suite.gitConfigs.GetBaseURL()
	suite.Equal(expect,actual)
}


func(suite GithubConfigsSuite)TestGetExchangeTokenURL_ShouldReturnUrl(){
	expect:="test://test.com/login/oauth/access_token"
	actual:=suite.gitConfigs.GetExchangeTokenURL()
	suite.Equal(expect,actual)
}