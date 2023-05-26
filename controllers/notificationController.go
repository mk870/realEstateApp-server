package controllers

import (
	"realEstateApi/middleware"
	"realEstateApi/services"

	"github.com/gin-gonic/gin"
)

func GetNotifications(router *gin.Engine) {
	router.GET("/notifications", middleware.AuthValidator, services.GetNotifications)
}

func DeleteNotification(router *gin.Engine) {
	router.DELETE("/notification/:id", middleware.AuthValidator, services.DeleteNotification)
}
