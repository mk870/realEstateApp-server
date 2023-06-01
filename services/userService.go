package services

import (
	"net/http"
	"time"

	"realEstateApi/models"
	"realEstateApi/repositories"
	"realEstateApi/tokens"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user models.User
	validateModelFields := validator.New()
	c.BindJSON(&user)

	modelFieldsValidationError := validateModelFields.Struct(user)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	refreshTokenResult := tokens.GenerateRefreshToken(user.FirstName, user.LastName, user.Email)
	if refreshTokenResult == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create refresh token",
		})
		return
	}

	verificationToken := CreateVerificationToken()
	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  string(hashedPassword),
		Email:     user.Email,
		RegistrationToken: models.VerificationToken{
			Token:      verificationToken.Token,
			ExpiryDate: verificationToken.ExpiryDate,
		},
	}

	isUserCreated := repositories.CreateUser(&newUser)
	if !isUserCreated {
		c.JSON(http.StatusForbidden, gin.H{"error": "this email already exists"})
		return
	}

	isVerificationEmailSent := SendVerificationEmail(user.Email, user.FirstName, "https://r-estates.vercel.app/verification/"+verificationToken.Token)
	if !isVerificationEmailSent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email"})
		return
	}
	c.String(http.StatusOK, "please check your email for verification")

}

func GetUsers(c *gin.Context) {
	usersList := repositories.GetUsers()
	c.JSON(http.StatusOK, usersList)
}

func GetUser(c *gin.Context) {
	type UserProfile struct {
		FirstName    string
		LastName     string
		Email        string
		Bio          string
		Photo        string
		DateOfBirth  string
		Phone        string
		City         string
		StreetName   string
		StreetNumber string
		Country      string
		State        string
		Id           int
	}
	var userProfile UserProfile
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	userProfile.Bio = user.Bio
	userProfile.City = user.City
	userProfile.Country = user.Country
	userProfile.State = user.State
	userProfile.DateOfBirth = user.DateOfBirth
	userProfile.StreetName = user.StreetName
	userProfile.StreetNumber = user.StreetNumber
	userProfile.Photo = user.Photo
	userProfile.Phone = user.Phone
	userProfile.Email = user.Email
	userProfile.FirstName = user.FirstName
	userProfile.LastName = user.LastName
	userProfile.Id = user.Id
	c.JSON(http.StatusOK, userProfile)
}

func UpdateUser(c *gin.Context) {
	var update models.User
	c.BindJSON(&update)
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if update.FirstName != "" {
		user.FirstName = update.FirstName
	}
	if update.LastName != "" {
		user.LastName = update.LastName
	}
	if update.Bio != "" {
		user.Bio = update.Bio
	}
	if update.DateOfBirth != "" {
		user.DateOfBirth = update.DateOfBirth
	}
	if update.StreetName != "" {
		user.StreetName = update.StreetName
	}
	if update.StreetNumber != "" {
		user.StreetNumber = update.StreetNumber
	}
	if update.City != "" {
		user.City = update.City
	}
	if update.Country != "" {
		user.Country = update.Country
	}
	if update.State != "" {
		user.State = update.State
	}
	if update.Phone != "" {
		user.Phone = update.Phone
	}
	if update.Photo != "" {
		user.Photo = update.Photo
	}
	updateResult := repositories.SaveUserUpdate(user)
	notificationData := models.Notification{
		Type:        "Profile",
		Action:      "Update",
		Description: "profile updated",
		Date:        time.Now().Format("02-01-2006"),
	}
	isNotificationCreated := CreateNotification(notificationData, user)
	if isNotificationCreated != "Notification saved" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "profile updated but notification creation failed",
		})
		return
	}
	if updateResult {
		c.String(http.StatusOK, "update successful")
	}
}

func UpdatePassword(c *gin.Context) {
	type PasswordBody struct {
		OldPassword     string
		NewPassword     string
		ConfirmPassword string
	}
	var body PasswordBody
	c.BindJSON(&body)
	if body.ConfirmPassword == "" || body.OldPassword == "" || body.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please enter all required fields"})
		return
	}
	if body.ConfirmPassword != body.NewPassword {
		c.JSON(http.StatusForbidden, gin.H{"error": "new password is not identical to the confirmed"})
		return
	}
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong old password"})
		return
	}
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user.Password = string(hashedNewPassword)
	updateResult := repositories.SaveUserUpdate(user)
	notificationData := models.Notification{
		Type:        "Password",
		Action:      "Update",
		Description: "password change",
		Date:        time.Now().Format("02-01-2006"),
	}
	isNotificationCreated := CreateNotification(notificationData, user)
	if isNotificationCreated != "Notification saved" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "password changed but notification creation failed",
		})
		return
	}
	if updateResult {
		c.String(http.StatusOK, "password change successful")
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	result := repositories.DeleteUserById(id)
	if result {
		c.String(http.StatusOK, "user successfully deleted")
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
	}
}
