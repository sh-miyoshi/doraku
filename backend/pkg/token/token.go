package token

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const testSecretKey = "ghoajg34qyiwgv3y4tgvobyqgqigkhiuqegwehrewhv3qha1254"

func Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "test",
	})

	return token.SignedString([]byte(testSecretKey))
}
