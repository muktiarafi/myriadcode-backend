package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
	"os"
	"time"
)

var jwtKey = os.Getenv("JWT_KEY")

type Claims struct {
	User *models.UserPayload
	jwt.StandardClaims
}

func CreateToken(user *models.UserPayload) (string, error) {
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(336 * time.Hour).Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string) (*jwt.Token, *models.UserPayload, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	return token, claims.User, err
}
