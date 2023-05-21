package models

import (
	"time"

	"gorm.io/gorm"
)

type MyModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	MyModel
	Id                int               `json:"id" gorm:"primary_key"`
	FirstName         string            `json:"firstName" validate:"required,min=2,max=50"`
	LastName          string            `json:"lastName" validate:"required,min=2,max=50"`
	Email             string            `json:"email" gorm:"unique" validate:"email,required"`
	Password          string            `json:"password" validate:"required,min=2,max=50"`
	Bio               string            `json:"bio"`
	Photo             string            `json:"photo" gorm:"nullable"`
	DateOfBirth       string            `json:"dateOfBirth"`
	Phone             string            `json:"phone"`
	City              string            `json:"city"`
	StreetName        string            `json:"streetName"`
	StreetNumber      string            `json:"streetNumber"`
	Country           string            `json:"country"`
	State             string            `json:"state"`
	RefreshToken      string            `json:"refreshToken"`
	IsActive          bool              `json:"isActive"`
	Properties        []Property        `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Agents            []Agent           `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Notifications     []Notification    `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RegistrationToken VerificationToken `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
