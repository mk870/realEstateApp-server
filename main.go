package main

import (
	"realEstateApi/controllers"
	"realEstateApi/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "token", "User-Agent", "Accept")
	router.Use(cors.New(config))
	db.Connect()
	controllers.GetVerificationToken(router)
	controllers.VerificationTokenValidation(router)
	controllers.GetNewAccessToken(router)
	controllers.CreateUser(router)
	controllers.GetUsers(router)
	controllers.UpdateUser(router)
	controllers.UpdatePassword(router)
	controllers.GetUser(router)
	controllers.DeleteUser(router)
	controllers.Login(router)
	controllers.LoginOut(router)
	controllers.CreateAgent(router)
	controllers.GetAgents(router)
	controllers.DeleteAgent(router)
	controllers.CreateProperty(router)
	controllers.GetProperties(router)
	controllers.DeleteProperty(router)
	controllers.GetNotifications(router)
	controllers.DeleteNotification(router)
	router.Run()
}
