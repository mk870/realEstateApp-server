package models

import "time"

type VerificationToken struct {
	MyModel
	Id         int       `json:"id" gorm:"primary_key"`
	UserId     int       `json:"userId"`
	Token      string    `json:"token"`
	ExpiryDate time.Time `json:"expirydate"`
}
