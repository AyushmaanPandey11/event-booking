package middlewares

import (
	"net/http"

	"eventBooking.com/m/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
