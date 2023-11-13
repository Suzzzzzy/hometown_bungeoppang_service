package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

func loginHandler(c *gin.Context) {
	kakaoURL := OauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, kakaoURL)
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
