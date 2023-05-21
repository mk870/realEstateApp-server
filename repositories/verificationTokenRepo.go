package repositories

import (
	"errors"

	"realEstateApi/db"
	"realEstateApi/models"

	"gorm.io/gorm"
)

func CreateVerificationToken(user *models.User, verificationToken models.VerificationToken) bool {
	db.DB.Model(user).Association("RegistrationToken").Replace(verificationToken)
	return true
}

func GetVerificationTokenById(userId string) models.VerificationToken {
	var user models.User
	err := db.DB.Preload("RegistrationToken").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.RegistrationToken
}

func GetVerificationTokenByToken(token string) *models.VerificationToken {
	var verificationToken = models.VerificationToken{}
	result := db.DB.Preload("RegistrationToken").Where("token =?", token).First(&verificationToken)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &verificationToken
}
