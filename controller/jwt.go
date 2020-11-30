package controller

import (
	"fmt"
	"graphyy/envvar"
	"graphyy/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TODO move this file to another package. (make a new package)
func generateJWT(user model.User) model.AuthToken {
	secret := envvar.AuthSecret()
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
