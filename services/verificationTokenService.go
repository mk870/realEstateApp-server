package services

import (
	"net/http"
	"strconv"
	"time"

	"realEstateApi/models"
	"realEstateApi/repositories"
	"realEstateApi/tokens"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateVerificationToken() models.VerificationToken {
	var verificationToken models.VerificationToken
	token := uuid.New().String()
	verificationToken.ExpiryDate = time.Now().Add(time.Minute * 15)
	verificationToken.Token = token
	return verificationToken
}

func VerifyToken(c *gin.Context) {
	userVerificationToken := c.Param("token")
	storedVerificationToken := repositories.GetVerificationTokenByToken(userVerificationToken)
	if storedVerificationToken == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "this verification token does not exist",
		})
		return
	}

	if storedVerificationToken.ExpiryDate.Unix() < time.Now().Local().Unix() {
		isUserDeleted := repositories.DeleteUserById(strconv.Itoa(storedVerificationToken.UserId))
		if !isUserDeleted {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "this user does not exist",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "this verification token has expired please signup again",
			})
			return
		}
	}
	user := repositories.GetUserById(strconv.Itoa(storedVerificationToken.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	refreshToken := tokens.GenerateRefreshToken(user.FirstName, user.LastName, user.Email)
	if refreshToken == "failed" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not generate your refresh token",
		})
		return
	}
	user.RefreshToken = refreshToken
	user.IsActive = true
	isUpdated := repositories.SaveUserUpdate(user)
	if isUpdated {
		accessToken := tokens.GenerateAccessToken(user.FirstName, user.LastName, user.Email, user.Id)
		if accessToken == "failed" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not generate your access token",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"accessToken": accessToken,
		})
		return
	}
}

func GetVerificationToken(c *gin.Context) {
	userId := c.Param("id")
	verificationToken := repositories.GetVerificationTokenById(userId)
	c.JSON(http.StatusOK, gin.H{
		"verificationToken": verificationToken,
	})
}
