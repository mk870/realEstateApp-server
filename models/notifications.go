package models

type Notification struct {
	MyModel
	Id          int    `json:"id" gorm:"primary_key"`
	UserId      int    `json:"userId"`
	Type        string `json:"type"`
	Action      string `json:"action"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
