package services

import (
	"net/http"

	"realEstateApi/models"
	"realEstateApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateNotification(notificationData models.Notification, user *models.User) string {
	newNotification := models.Notification{
		Type:        notificationData.Type,
		Action:      notificationData.Action,
		Description: notificationData.Description,
		Date:        notificationData.Date,
		UserId:      user.Id,
	}
	isNotificationCreated := repositories.CreateNotification(user, &newNotification)
	if !isNotificationCreated {
		return "notification could not be saved"
	}
	return "Notification saved"
}

func GetNotifications(c *gin.Context) {
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}
	notificationList := repositories.GetNotifications(user.Id)
	c.JSON(http.StatusOK, notificationList)
}

func DeleteNotification(c *gin.Context) {
	notification_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}
	notification := repositories.GetNotification(user.Id, notification_id)
	isDeleted := repositories.DeleteNotificationById(user, notification)
	if isDeleted {
		c.JSON(http.StatusOK, "Delete successful")
		return
	}
}
