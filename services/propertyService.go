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
		Description: "added property to your " + property.Status + " watchlist",
		Date:        time.Now().Format("02-01-2006"),
	}
	isNotificationCreated := CreateNotification(notificationData, user)
	if isNotificationCreated != "Notification saved" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "property saved but notification creation failed",
		})
		return
	}
	c.String(http.StatusOK, "Property saved")
}

func GetProperties(c *gin.Context) {
	property_type := c.Param("type")
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
	requestedPropertyList := []models.Property{}
	for _, property := range propertyList {
		if property.Status == property_type {
			requestedPropertyList = append(requestedPropertyList, property)
		}
	}
	c.JSON(http.StatusOK, requestedPropertyList)
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
	notificationData := models.Notification{
		Type:        "Property",
		Action:      "Delete",
		Description: "deleted property from " + property.Status + " watchlist",
		Date:        time.Now().Format("02-01-2006"),
	}
	isNotificationCreated := CreateNotification(notificationData, user)
	if isNotificationCreated != "Notification saved" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "property deleted but notification creation failed",
		})
		return
	}
	if isDeleted {
		c.JSON(http.StatusOK, "Delete successful")
		return
	}
}
