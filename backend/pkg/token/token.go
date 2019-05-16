package token

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

const testSecretKey = "ghoajg34qyiwgv3y4tgvobyqgqigkhiuqegwehrewhv3qha1254"

func Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "test",
	})

	return token.SignedString([]byte(testSecretKey))
}

func Validate(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(testSecretKey), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, fmt.Errorf("Failed to validate token")
	}
}
