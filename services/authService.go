package services

import (
	"net/http"

	"realEstateApi/models"
	"realEstateApi/repositories"
	"realEstateApi/tokens"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RefreshAccessToken(c *gin.Context) {
	expiredAccessToken := c.Request.Header.Get("token")
	if expiredAccessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no access token found",
		})
		return
	}

	accessTokenClaims, _ := tokens.ValidateAccessToken(expiredAccessToken)
	user := repositories.GetUserByEmail(accessTokenClaims.Email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get user information",
		})
		return
	}
	refreshToken := user.RefreshToken
	if refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no refresh token found",
		})
		return
	}
	_, err := tokens.ValidateRefreshToken(refreshToken)

	var newRefreshToken string
	if err == "token has expired" {
		newRefreshToken = tokens.GenerateRefreshToken(user.FirstName, user.LastName, user.Email)
		if newRefreshToken == "failed" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not create a new refresh token",
			})
			return
		}

		user.RefreshToken = newRefreshToken
		isStudentUpdated := repositories.SaveUserUpdate(user)
		if isStudentUpdated {
			newAccessToken := tokens.GenerateAccessToken(user.FirstName, user.LastName, user.Email, user.Id)
			if newAccessToken == "failed" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "could not create a new access token",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"accessToken": newAccessToken,
			})
			return
		}
	}
	newAccessToken := tokens.GenerateAccessToken(user.FirstName, user.LastName, user.Email, user.Id)
	if newAccessToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create a new access token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessToken,
	})
}

func Logout(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	loggedInUser := repositories.GetUserByEmail(user.Email)
	if loggedInUser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user information",
		})
		return
	}
	loggedInUser.RefreshToken = ""
	isRefreshTokenDeleted := repositories.SaveUserUpdate(loggedInUser)
	if isRefreshTokenDeleted {
		c.JSON(http.StatusOK, "you have logged out successfully")
		return
	}
}

func Login(c *gin.Context) {
	type Login struct {
		Email    string
		Password string
	}
	var loginData = Login{}
	c.BindJSON(&loginData)
	if loginData.Email == "" || loginData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please enter both email and password"})
		return
	}
	user := repositories.GetUserByEmail(loginData.Email)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this account does not exist"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password or email"})
		return
	}

	refreshToken := tokens.GenerateRefreshToken(user.FirstName, user.LastName, user.Email)
	if refreshToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create refresh token",
		})
		return
	}
	user.RefreshToken = refreshToken
	repositories.SaveUserUpdate(user)
	accessToken := tokens.GenerateAccessToken(user.FirstName, user.LastName, user.Email, user.Id)
	if accessToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create access token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
		"response":    "logged in successfully",
	})
}
