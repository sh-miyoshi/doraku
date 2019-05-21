package token

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sh-miyoshi/doraku/pkg/logger"
	"strings"
	"time"
)

// TODO(use secure key)
const testSecretKey = "ghoajg34qyiwgv3y4tgvobyqgqigkhiuqegwehrewhv3qha1254"
const dorakuIssuer = "doraku"

func Generate() (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    dorakuIssuer,
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // Expired at 2 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(testSecretKey))
}

func validate(tokenString string) (jwt.StandardClaims, error) {
	claims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(testSecretKey), nil
	})

	if err != nil {
		return jwt.StandardClaims{}, err
	}

	logger.Debug("Claims: %v\n", claims)

	if token.Valid {
		return claims, nil
	} else {
		return jwt.StandardClaims{}, fmt.Errorf("Failed to validate token")
	}
}

func parseHTTPHeaderToken(tokenString string) (string, error) {
	var splitToken []string
	if strings.Contains(tokenString, "bearer") {
		splitToken = strings.Split(tokenString, "bearer")
	} else if strings.Contains(tokenString, "Bearer") {
		splitToken = strings.Split(tokenString, "Bearer")
	} else {
		return "", fmt.Errorf("token format is missing")
	}
	reqToken := strings.TrimSpace(splitToken[1])
	return reqToken, nil
}

func Authenticate(reqToken string) error {
	tokenStr, err := parseHTTPHeaderToken(reqToken)
	if err != nil {
		logger.Info("Failed to get JWT token %v", err)
		return err
	}
	logger.Debug("Token String: %s", tokenStr)

	claims, err := validate(tokenStr)
	if err != nil {
		logger.Info("Failed to auth token %v", err)
		return err
	}
	logger.Debug("claims in token: %v", claims)

	// Validate claims
	if claims.Issuer != dorakuIssuer {
		logger.Info("Issuer want %s, but got %s", dorakuIssuer, claims.Issuer)
		return fmt.Errorf("Issuer want %s, but got %s", dorakuIssuer, claims.Issuer)
	}

	now := time.Now().Unix()
	if now > claims.ExpiresAt {
		logger.Info("Token is expired at %d. now: %d", claims.ExpiresAt, now)
		return fmt.Errorf("Token is expired at %d. now: %d", claims.ExpiresAt, now)
	}

	return nil
}
