package controllers

import (
	"realEstateApi/services"

	"github.com/gin-gonic/gin"
)

func VerificationTokenValidation(router *gin.Engine) {
	router.GET("/verification/:token", services.VerifyToken)
}

func GetVerificationToken(router *gin.Engine) {
	router.GET("/verification-token/:id", services.GetVerificationToken)
}
