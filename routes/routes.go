package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	// EVENT'S ROUTES
	server.GET("/health-check", healthCheck)
	server.GET("/get-all-events", getEvents)
	server.POST("/add-event", addEvent)
	server.GET("/event/:id", getEventById)
	server.PUT("/event/:id", updateEvent)
	server.DELETE("/event/:id", deleteEvent)

	// USER ROUTES
	server.POST("/add-user", registerUser)
	server.GET("/get-all-users", getAllUserDetails)
	server.GET("/get-user/:id", getUserById)
	server.PUT("/user/:id", updateUser)
	server.DELETE("/user/:id", deleteUser)

}
