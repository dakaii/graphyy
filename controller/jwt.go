package controller

import (
	"errors"
	"fmt"
	"graphyy/envvar"
	"graphyy/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TODO move this file to another package. (make a new package)
func generateJWT(user model.User) model.AuthToken {
	secret := envvar.AuthSecret()
	expiresAt := time.Now().Add(time.Minute * 15).Unix()

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

func verifyJWT(tknStr string) (model.User, error) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(envvar.AuthSecret()), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return model.User{}, errors.New("signature invalid")
		}
		return model.User{}, errors.New("could not parse the auth token")
	}
	if !token.Valid {
		return model.User{}, errors.New("Invalid token")
	}
	fmt.Println("TOKEN is:", token.Valid)

	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	data := claims["data"].(map[string]interface{})
	username := data["username"].(string)
	createdAt := data["createdAt"].(time.Time)
	return model.User{Username: username, CreatedAt: createdAt}, nil
}
