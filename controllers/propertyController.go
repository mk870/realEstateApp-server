package controllers

import (
	"realEstateApi/middleware"
	"realEstateApi/services"

	"github.com/gin-gonic/gin"
)

func CreateProperty(router *gin.Engine) {
	router.POST("/property", middleware.AuthValidator, services.CreateProperty)
}

func GetProperties(router *gin.Engine) {
	router.GET("/properties/:type", middleware.AuthValidator, services.GetProperties)
}

func DeleteProperty(router *gin.Engine) {
	router.DELETE("/property/:id", middleware.AuthValidator, services.DeleteProperty)
}
