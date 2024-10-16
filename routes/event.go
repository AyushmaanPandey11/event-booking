package routes

import (
	"net/http"
	"strconv"

	"eventBooking.com/m/models"
	"github.com/gin-gonic/gin"
)

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is Working Properly"})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error occured in fetching events from DB"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func addEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data sent!"})
	}
	userId := context.GetInt64("userId")
	event.User_id = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error occured in adding event to DB"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Event created!", "event": event})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find the event"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Event Found Succesfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event"})
		return
	}
	userId := context.GetInt64("userId")
	if userId != event.User_id {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event"})
		return
	}
	updatedEvent.Id = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in updating event"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Updated Succesfully", "Event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event"})
		return
	}
	userId := context.GetInt64("userId")
	if userId != event.User_id {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error in deleting the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully", "event": event})
}

func userRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event"})
		return
	}
	err = event.UserRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in registering user for the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user registered successfully", "event": event})

}

func userCancellation(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}
	var event models.Event
	event.Id = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in cancelling user registration for the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user registration cancelled successfully", "event": event})

}
