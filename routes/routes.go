package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health-check", healthCheck)
	server.GET("/get-all-events", getEvents)
	server.POST("/add-event", addEvent)
	server.GET("/event/:id", getEventById)
	server.PUT("/event/:id", updateEvent)
	server.DELETE("/event/:id", deleteEvent)
}
