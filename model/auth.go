package model

import (
	"github.com/dgrijalva/jwt-go"
)

//AuthToken struct
type AuthToken struct {
	TokenType string `json:"tokenType" graphql:"tokenType"`
	Token     string `json:"accessToken" graphql:"token"`
	ExpiresIn int64  `json:"expiresIn" graphql:"expiresIn"`
}

//AuthTokenClaim struct
type AuthTokenClaim struct {
	jwt.StandardClaims
	User
}
