package services

import (
	"net/http"

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

	isVerificationEmailSent := SendVerificationEmail(user.Email, user.FirstName, "https://movie-plus-frontend.vercel.app/verification/"+verificationToken.Token)
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
	id := c.Param("id")
	user := repositories.GetUserById(id)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var update models.User
	c.BindJSON(&update)
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	oldData := repositories.GetUserByEmail(email)
	if oldData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if update.FirstName != "" {
		oldData.FirstName = update.FirstName
	}
	if update.LastName != "" {
		oldData.LastName = update.LastName
	}
	if update.Bio != "" {
		oldData.Bio = update.Bio
	}
	if update.DateOfBirth != "" {
		oldData.DateOfBirth = update.DateOfBirth
	}
	if update.StreetName != "" {
		oldData.StreetName = update.StreetName
	}
	if update.StreetNumber != "" {
		oldData.StreetNumber = update.StreetNumber
	}
	if update.City != "" {
		oldData.City = update.City
	}
	if update.Country != "" {
		oldData.Country = update.Country
	}
	if update.State != "" {
		oldData.State = update.State
	}
	if update.Phone != "" {
		oldData.Phone = update.Phone
	}
	if update.Photo != "" {
		oldData.Photo = update.Photo
	}
	updateResult := repositories.SaveUserUpdate(oldData)
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
	if updateResult {
		c.String(http.StatusOK, "update successful")
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
