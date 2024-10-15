package routes

import (
	"net/http"
	"strconv"

	"eventBooking.com/m/models"
	"eventBooking.com/m/utils"
	"github.com/gin-gonic/gin"
)

func registerUser(context *gin.Context) {
	var newUser models.User
	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unavailable to create account!"})
		return
	}
	newUser.Id = 101
	err = newUser.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error inserting user to database"})
		return
	}
	token, err := utils.GenerateToken(newUser.Email, newUser.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error Creating token"})
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{"message": "User Created Successfully", "user": newUser, "token": token})
}

func loginUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid credentials"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unavailable to login"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error Creating token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "user": user, "token": token})
}

func getAllUserDetails(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error occured in fetching users from DB"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func getUserById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user, User Doesn't Exists"})
		return
	}
	user, err := models.GetUserById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error occured in fetching user from DB"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func updateUser(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user, User Doesn't Exists"})
		return
	}
	user, err := models.GetUserById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error occured in fetching user from DB"})
		return
	}
	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	updatedUser.Id = user.Id
	err = updatedUser.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "User Created successfully", "createdUser": updatedUser})
}

func deleteUser(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user, User Doesn't Exists"})
		return
	}
	user, err := models.GetUserById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error occured in fetching user from DB"})
		return
	}
	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error in deleting user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "deletedUser": user})
}
