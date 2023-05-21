package repositories

import (
	"errors"

	"realEstateApi/db"
	"realEstateApi/models"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) bool {
	var existingUser models.User
	email := user.Email
	result := db.DB.Where("email =?", email).First(&existingUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.DB.Create(user)
		return true
	} else {
		return false
	}
}

func GetUsers() []models.User {
	var users = []models.User{}
	err := db.DB.Find(&users)
	if err != nil {
		println(err.Error, err.Name())
	}
	return users
}

func GetUserById(id string) *models.User {
	var user models.User
	result := db.DB.Where("id = ?", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func SaveUserUpdate(update *models.User) bool {
	db.DB.Save(update)
	return true
}

func DeleteUserById(id string) bool {
	user := GetUserById(id)
	if user == nil {
		return false
	}
	db.DB.Unscoped().Delete(&models.User{}, id)
	return true
}

func GetUserByEmail(email string) *models.User {
	var user models.User
	result := db.DB.Where("email =?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	} else {
		return &user
	}
}
