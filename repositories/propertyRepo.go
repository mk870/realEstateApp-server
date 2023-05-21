package repositories

import (
	"errors"

	"realEstateApi/db"
	"realEstateApi/models"

	"gorm.io/gorm"
)

func CreateProperty(user *models.User, property *models.Property) bool {
	err := db.DB.Model(user).Association("Properties").Append(property)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetProperties(id int) []models.Property {
	var user = models.User{}
	err := db.DB.Preload("Properties").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Properties
}

func GetProperty(userId int, propertyId string) models.Property {
	var user = models.User{}
	err := db.DB.Preload("Properties", "id=?", propertyId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Properties[0]
}

func GetUserWithPropertiesById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Properties").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateProperty(user *models.User, updateList []models.Property) bool {
	user.Properties = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeletePropertyById(user *models.User, property models.Property) bool {
	db.DB.Model(&user).Unscoped().Association("Properties").Delete(property)
	return true
}
