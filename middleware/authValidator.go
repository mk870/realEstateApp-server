package middleware

import (
	"net/http"
	"strings"

	"realEstateApi/repositories"
	"realEstateApi/tokens"

	"github.com/gin-gonic/gin"
)

func AuthValidator(c *gin.Context) {
	clientAuthorizationHeader := c.Request.Header.Get("Authorization")
	clientToken := strings.Split(clientAuthorizationHeader, " ")[1]
	if clientToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, you dont have an access token",
		})
		c.Abort()
		return
	}
	claims, err := tokens.ValidateAccessToken(clientToken)
	if err != "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})
		c.Abort()
		return
	}
	user := repositories.GetUserByEmail(claims.Email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "invalid token",
		})
		c.Abort()
		return
	}
	c.Set("user", user)
	c.Next()
}
