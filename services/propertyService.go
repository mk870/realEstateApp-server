package services

import (
	"net/http"
	"time"

	"realEstateApi/models"
	"realEstateApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateProperty(c *gin.Context) {
	var property models.Property
	err := c.BindJSON(&property)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to bind request body",
		})
		return
	}
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}
	isPropertyCreated := repositories.CreateProperty(user, &property)
	if !isPropertyCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not save property",
		})
		return
	}
	notificationData := models.Notification{
		Type:        "Property",
		Action:      "Post",
		Description: "added property tmx to your rentals watchlist",
		Date:        time.Now().Format("dd-mm-yy"),
	}
	isNotificationCreated := CreateNotification(notificationData, user)
	if isNotificationCreated != "Notification saved" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "property saved but notification not created",
		})
		return
	}
	c.String(http.StatusOK, "Property saved")
}

func GetProperties(c *gin.Context) {
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}
	propertyList := repositories.GetProperties(user.Id)
	c.JSON(http.StatusOK, propertyList)
}

func DeleteProperty(c *gin.Context) {
	property_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}
	property := repositories.GetProperty(user.Id, property_id)
	isDeleted := repositories.DeletePropertyById(user, property)

	if isDeleted {
		c.JSON(http.StatusOK, "Delete successful")
		return
	}
}
