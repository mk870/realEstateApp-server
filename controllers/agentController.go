package controllers

import (
	"realEstateApi/middleware"
	"realEstateApi/services"

	"github.com/gin-gonic/gin"
)

func CreateAgent(router *gin.Engine) {
	router.POST("/agent", middleware.AuthValidator, services.CreateAgent)
}

func GetAgents(router *gin.Engine) {
	router.GET("/agents", middleware.AuthValidator, services.GetAgents)
}

func DeleteAgent(router *gin.Engine) {
	router.DELETE("/agent/:id", middleware.AuthValidator, services.DeleteAgent)
}
