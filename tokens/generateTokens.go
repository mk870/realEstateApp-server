package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type RefreshTokenClaims struct {
	FirstName string
	LastName  string
	Email     string
	jwt.StandardClaims
}

type AccessTokenClaims struct {
	FirstName string
	LastName  string
	Email     string
	Id        string
	Photo     string
	jwt.StandardClaims
}

func GenerateAccessToken(firstName string, lastName string, email string, id int, photo string) string {
	claims := &AccessTokenClaims{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Id:        string(rune(id)),
		Photo:     photo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(60)).Unix(),
			Id:        string(rune(id)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(GetAccessTokenSecret()))
	if err != nil {
		println("error", err.Error())
		return "failed"
	} else {
		return tokenString
	}

}

func GenerateRefreshToken(firstName string, lastName string, email string) string {
	claims := &RefreshTokenClaims{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(GetRefreshTokenSecret()))
	if err != nil {
		return "failed"
	} else {
		return tokenString
	}

}
