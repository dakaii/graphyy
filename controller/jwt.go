package controller

import (
	"fmt"
	"graphyy/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TODO move this file to another package. (make a new package)
func generateJWT(user model.User) model.AuthToken {
	secret, exists := os.LookupEnv("AUTH_SECRET")
	if !exists {
		secret = "secret_key"
	}

	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &model.AuthTokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User: user,
	}

	tokenString, error := token.SignedString([]byte(secret))
	if error != nil {
		fmt.Println(error)
	}
	return model.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	}
}
