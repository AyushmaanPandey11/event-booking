package main

import (
	"eventBooking.com/m/db"
	"eventBooking.com/m/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
