package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahul-golang/github_app/configurations"
	error2 "github.com/rahul-golang/github_app/error"
	"github.com/rahul-golang/github_app/pkg/controller"
	"github.com/rahul-golang/github_app/pkg/service"
	"net/http"
)

//const clientID = ""
//const clientSecret = ""

func main() {
	config := configurations.NewConfiguration("./configurations/config")
	configurations.NewGithubConfigs(config.GithubConfig)
	errInterceptor := error2.NewErrResponseInterceptor()
	githubDownStreamService := service.NewGithubDownStream()
	oauthService := service.NewGithubOauthService(config.GithubConfig, githubDownStreamService)
	authController := controller.NewOAuthController(errInterceptor, oauthService)

	stream := service.NewGithubApiDownStream()
	apiService := service.NewGithubApiService(stream)
	apiController := controller.NewGithubApiController(errInterceptor, apiService)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//fs := http.FileServer(http.Dir("public"))
	//http.Handle("/", fs)

	r.StaticFS("/log", http.Dir("public"))
	r.GET("/oauth/redirect", authController.Redirect)
	r.GET("/git/get-repo", apiController.GetRepository)
	r.POST("/git/pull-request", apiController.CreatePullRequest)

	//os.Setenv("GITHUB_CLIENT_ID", clientID)
	//os.Setenv("GITHUB_CLIENT_SECRET", clientSecret)

	//service.NewGithubApiDownStream().AccessRepositories()
	r.Run()
}
