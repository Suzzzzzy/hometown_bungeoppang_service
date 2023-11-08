package handler

import (
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
	RedirectURL = os.Getenv("REDIRECT_URL")
)

func generateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(1 * 24 * time.Hour)

	b := make([byte], 16)
	rand.Read(b)

}
