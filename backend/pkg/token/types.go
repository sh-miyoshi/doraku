package token

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// CreateUserClaims is claim for create user
type CreateUserClaims struct {
	Name           string `json:"name"`
	HashedPassword string `json:"hashedPassword"`

	jwt.StandardClaims
}

// TODO(Error defines)