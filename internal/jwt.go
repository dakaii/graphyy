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
		return entity.User{}, errors.New("invalid token")
	}

	decoded := make(map[string]interface{})
	for key, val := range claims {
		decoded[key] = val
	}
	var username string
	if keyExists(decoded, "username") {
		username = decoded["username"].(string)
	}

	// var createdAt time.Time
	// if keyExists(decoded, "createdAt") {
	// 	createdAt, err = parseTimeString(decoded["createdAt"])
	// 	if err != nil {
	// 		return entity.User{}, fmt.Errorf("could not parse createdAt time: %w", err)
	// 	}
	// }
	return entity.User{Username: username}, nil
}

func keyExists(decoded map[string]interface{}, key string) bool {
	val, ok := decoded[key]
	return ok && val != nil
}

// func parseTimeString(param interface{}) (time.Time, error) {
// 	timeStringTypes := []string{time.RFC3339, time.RFC3339Nano, "2006-01-02T15:04:05.999999999Z07:00", "2024-04-27T02:56:51.45722294Z"}
// 	var t time.Time
// 	var err error

// 	switch v := param.(type) {
// 	case string:
// 		for _, timeStringType := range timeStringTypes {
// 			t, err = time.Parse(timeStringType, v)
// 		}
// 		if t.IsZero() {
// 			return time.Time{}, fmt.Errorf("could not parse string to time: %w", err)
// 		}
// 	case time.Time:
// 		t = v
// 	default:
// 		return time.Time{}, errors.New("unsupported type for time parsing")
// 	}

// 	return t, nil
// }
