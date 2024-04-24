package internal

import (
	"errors"
	"fmt"
	"graphyy/entity"
	"graphyy/internal/envvar"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TODO move this file to another package. (make a new package)
func GenerateJWT(user entity.User) entity.AuthToken {
	secret := envvar.AuthSecret()
	expiresAt := time.Now().Add(time.Minute * 15).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &entity.AuthTokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User: user,
	}

	tokenString, error := token.SignedString([]byte(secret))
	if error != nil {
		fmt.Println(error)
	}
	return entity.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	}
}

func VerifyJWT(tknStr string) (entity.User, error) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(envvar.AuthSecret()), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return entity.User{}, errors.New("signature invalid")
		}
		return entity.User{}, errors.New("could not parse the auth token")
	}
	if !token.Valid {
		return entity.User{}, errors.New("Invalid token")
	}
	fmt.Println("TOKEN is:", token.Valid)

	decoded := make(map[string]interface{})
	for key, val := range claims {
		decoded[key] = val
	}
	// decoded := claims["data"].([]interface{})
	var username string
	if keyExists(decoded, "username") {
		username = decoded["username"].(string)
	}

	var createdAt time.Time
	if keyExists(decoded, "createdAt") {
		createdAt = decoded["createdAt"].(time.Time)
	}
	return entity.User{Username: username, CreatedAt: createdAt}, nil
}

func keyExists(decoded map[string]interface{}, key string) bool {
	val, ok := decoded[key]
	return ok && val != nil
}
