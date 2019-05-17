package token

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// TODO(use secure key)
const testSecretKey = "ghoajg34qyiwgv3y4tgvobyqgqigkhiuqegwehrewhv3qha1254"

func Generate() (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    "doraku",
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // Expired at 2 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(testSecretKey))
}

func Validate(tokenString string) (jwt.StandardClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(testSecretKey), nil
	})

	if err != nil {
		return jwt.StandardClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.StandardClaims{}, fmt.Errorf("Failed to validate token")
	}
}
