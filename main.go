package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/health-check", healthCheck)
	server.Run(":8080")
}
func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is Working Properly"})
}
