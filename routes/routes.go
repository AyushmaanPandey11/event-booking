package routes

import (
	"eventBooking.com/m/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// EVENT'S ROUTES
	server.GET("/health-check", healthCheck)
	server.GET("/get-all-events", getEvents)
	eventAuthenticated := server.Group("/event")
	eventAuthenticated.Use(middlewares.Authenticate)
	eventAuthenticated.POST("/add-event", addEvent)
	eventAuthenticated.GET("/:id", getEventById)
	eventAuthenticated.PUT("/:id", updateEvent)
	eventAuthenticated.DELETE("/:id", deleteEvent)

	// USER ROUTES
	server.POST("/add-user", registerUser)
	server.POST("/login-user", loginUser)
	userAuthenticated := server.Group("/user")
	userAuthenticated.Use(middlewares.Authenticate)
	userAuthenticated.GET("/get-all-users", getAllUserDetails)
	userAuthenticated.GET("/get-user/:id", getUserById)
	userAuthenticated.PUT("/:id", updateUser)
	userAuthenticated.DELETE("/:id", deleteUser)

	// user event registration
	userAuthenticated.POST("/:id/register", userRegistration)
	userAuthenticated.DELETE("/:id/register", userCancellation)
}
