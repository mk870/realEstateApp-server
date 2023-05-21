package repositories

import (
	"errors"

	"realEstateApi/db"
	"realEstateApi/models"

	"gorm.io/gorm"
)

func CreateAgent(user *models.User, agent *models.Agent) bool {
	err := db.DB.Model(user).Association("Agents").Append(agent)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetAgents(id int) []models.Agent {
	var user = models.User{}
	err := db.DB.Preload("Agents").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Agents
}

func GetAgent(userId int, agentId string) models.Agent {
	var user = models.User{}
	err := db.DB.Preload("Agents", "id=?", agentId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Agents[0]
}

func GetUserWithAgentsById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Agents").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateAgent(user *models.User, updateList []models.Agent) bool {
	user.Agents = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteAgentById(user *models.User, agent models.Agent) bool {
	db.DB.Model(&user).Unscoped().Association("Agents").Delete(agent)
	return true
}
