package transaction_service

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/app/auth/login", kakaoLoginHandler)
	r.GET("/app/auth/callback", kakaoAuthCallback)

	r.Run(":8080")
}
