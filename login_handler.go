package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

func loginHandler(c *gin.Context) {
	kakaoURL := OauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, kakaoURL)

	fmt.Println("CLIENT_ID:", os.Getenv("CLIENT_ID"))
	fmt.Println("CLIENT_ID: ", OauthConfig.ClientID)

}

func callbackHandler(c *gin.Context) {
	code := c.Query("code")
	token, err := OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "[Error] Failed to handle kakao callback")
		return
	}

	c.String(200, "Token: "+token.AccessToken)
}
