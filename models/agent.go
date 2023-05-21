package models

type Agent struct {
	MyModel
	Id       int    `json:"id" gorm:"primary_key"`
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	Agent_id int    `json:"agent_id"`
	Photo    string `json:"photo" gorm:"nullable"`
}
