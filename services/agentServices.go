package services

import (
	"net/http"
	"time"

	"realEstateApi/models"
	"realEstateApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateAgent(c *gin.Context) {
	var agent models.Agent
	err := c.BindJSON(&agent)
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
	isAgentCreated := repositories.CreateAgent(user, &agent)
	if !isAgentCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not save agent",
		})
		return
	}
	notificationData := models.Notification{
		Type:        "Agent",
		Action:      "Post",
		Description: "added " + agent.Name + " to your agent watchlist",
		Date:        time.Now().Format("02-01-2006"),
	}
	isNotificationCreated := CreateNotification(notificationData, user)
	if isNotificationCreated != "Notification saved" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "agent saved but notification creation failed",
		})
		return
	}
	c.String(http.StatusOK, "Agent saved")
}

func GetAgents(c *gin.Context) {
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}
	agentList := repositories.GetAgents(user.Id)
	c.JSON(http.StatusOK, agentList)
}

func DeleteAgent(c *gin.Context) {
	agent_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}
	agent := repositories.GetAgent(user.Id, agent_id)
	isDeleted := repositories.DeleteAgentById(user, agent)
	notificationData := models.Notification{
		Type:        "Agent",
		Action:      "Delete",
		Description: "deleted " + agent.Name + " from your agent watchlist",
		Date:        time.Now().Format("02-01-2006"),
	}
	isNotificationCreated := CreateNotification(notificationData, user)
	if isNotificationCreated != "Notification saved" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "agent deleted but notification creation failed",
		})
		return
	}
	if isDeleted {
		c.JSON(http.StatusOK, "Delete successful")
		return
	}
}
