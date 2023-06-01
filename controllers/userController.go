package controllers

import (
	"realEstateApi/middleware"
	"realEstateApi/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(router *gin.Engine) {
	router.POST("/user", services.CreateUser)
}

func GetUsers(router *gin.Engine) {
	router.GET("/users", middleware.AuthValidator, services.GetUsers)
}

func UpdateUser(router *gin.Engine) {
	router.PUT("/user", middleware.AuthValidator, services.UpdateUser)
}

func UpdatePassword(router *gin.Engine) {
	router.PUT("/user/password", middleware.AuthValidator, services.UpdatePassword)
}

func GetUser(router *gin.Engine) {
	router.GET("/user", middleware.AuthValidator, services.GetUser)
}

func DeleteUser(router *gin.Engine) {
	router.DELETE("/user/:id", middleware.AuthValidator, services.DeleteUser)
}
