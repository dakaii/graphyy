package controllers

import (
	"fmt"
	"graphyy/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TODO move this file to another package. (make a new package)
func generateJWT(user models.User) models.AuthToken {
	secret, exists := os.LookupEnv("AUTH_SECRET")
	if !exists {
		secret = "secret_key"
	}

	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &models.AuthTokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User: models.User{Username: user.Username, Password: user.Password},
	}

	tokenString, error := token.SignedString([]byte(secret))
	if error != nil {
		fmt.Println(error)
	}
	return models.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	}
}
