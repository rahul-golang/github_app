package configurations

type GithubConfigs interface {
	GetBaseURL() string
	GetExchangeTokenURL() string
}

func NewGithubConfigs(githubConfig GithubConfig) GithubConfigs {
	return githubConfig
}

func (config GithubConfig) GetBaseURL() string {
	return config.BaseUrl
}
func (config GithubConfig) GetExchangeTokenURL() string {
	return config.GetBaseURL() + "/login/oauth/access_token"
}
