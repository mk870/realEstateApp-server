package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ValidateAccessToken(clientToken string) (claims *AccessTokenClaims, msg string) {
	token, err := jwt.ParseWithClaims(clientToken, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(GetAccessTokenSecret()), nil
	})
	if err != nil {
		msg = err.Error()
	}
	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		msg = "token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token has expired"
		return
	}
	return claims, msg
}

func ValidateRefreshToken(refreshToken string) (claims *RefreshTokenClaims, msg string) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(GetRefreshTokenSecret()), nil
	})
	if err != nil {
		msg = err.Error()

	}
	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok {
		msg = "token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token has expired"
		return
	}
	return claims, msg
}
