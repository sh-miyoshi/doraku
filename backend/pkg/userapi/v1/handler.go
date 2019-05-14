package userapi

import (
	"encoding/json"
	"net/http"

	"github.com/sh-miyoshi/doraku/pkg/logger"
)

// LoginHandler validates user id and password, and return JWT token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call LoginHandler method with Body: %v", r.Body)

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Info("Failed to decode Login params: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// TODO access to userdb and check id and password
	// TODO generate JWT token(https://github.com/dgrijalva/jwt-go)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//w.Write(resRaw)
	logger.Info("Successfully finished LoginHandler")
}
