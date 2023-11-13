package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"os"
)

var (
	ClientID     = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
	RedirectURL  = os.Getenv("REDIRECT_URL")
)

var OauthConfig = &oauth2.Config{
	ClientID:     ClientID,
	ClientSecret: ClientSecret,
	RedirectURL:  RedirectURL,
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://kauth.kakao.com/oauth/authorize",
		TokenURL: "https://kauth.kakao.com/oauth/token",
	},
}

func main() {
	r := gin.Default()

	app := r.Group("/app")

	app.GET("/login", loginHandler)
	app.GET("/callback", callbackHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
