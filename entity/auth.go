package entity

import (
	"github.com/dgrijalva/jwt-go"
)

// AuthToken struct
type AuthToken struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"accessToken"`
	ExpiresIn int64  `json:"expiresIn"`
}

// AuthTokenClaim struct
type AuthTokenClaim struct {
	jwt.StandardClaims
	User
}
