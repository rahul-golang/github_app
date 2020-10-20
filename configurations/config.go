package configurations

type config struct {
	GithubConfig      GithubConfig      `json:"github_config"`
}

type GithubConfig struct {
	BaseUrl     string `json:"base_url"`
}
