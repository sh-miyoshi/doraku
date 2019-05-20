package util

import (
	"fmt"
	"strings"
)

func TokenParse(tokenString string) (string, error) {
	var splitToken []string
	if strings.Contains(tokenString, "bearer") {
		splitToken = strings.Split(tokenString, "bearer")
	} else if strings.Contains(tokenString, "Bearer") {
		splitToken = strings.Split(tokenString, "Bearer")
	} else {
		return "", fmt.Errorf("token format is missing")
	}
	reqToken := strings.TrimSpace(splitToken[0])
	return reqToken, nil
}
