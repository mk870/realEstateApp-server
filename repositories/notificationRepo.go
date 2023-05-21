package repositories

import (
	"errors"

	"realEstateApi/db"
	"realEstateApi/models"

	"gorm.io/gorm"
)

func CreateNotification(user *models.User, notification *models.Notification) bool {
	err := db.DB.Model(user).Association("Notifications").Append(notification)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetNotifications(id int) []models.Notification {
	var user = models.User{}
	err := db.DB.Preload("Notifications").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Notifications
}

func GetNotification(userId int, notificationId string) models.Notification {
	var user = models.User{}
	err := db.DB.Preload("Notifications", "id=?", notificationId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Notifications[0]
}

func GetUserWithNotificationsById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Notifications").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateNotification(user *models.User, updateList []models.Notification) bool {
	user.Notifications = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteNotificationById(user *models.User, notification models.Notification) bool {
	db.DB.Model(&user).Unscoped().Association("Notifications").Delete(notification)
	return true
}
