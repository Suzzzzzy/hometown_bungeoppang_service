package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"log"
	"os"
)

var OauthConfig *oauth2.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	OauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://kauth.kakao.com/oauth/authorize",
			TokenURL: "https://kauth.kakao.com/oauth/token",
		},
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	app := r.Group("/app")

	app.GET("/auth/login", loginHandler)
	app.GET("/auth/callback", callbackHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
