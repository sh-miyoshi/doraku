package userapi

import (
	"encoding/json"
	"net/http"

	"github.com/sh-miyoshi/doraku/pkg/logger"
	"github.com/sh-miyoshi/doraku/pkg/userdb"
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

	token, err := userdb.GetInst().Authenticate(req.ID, req.Password)

	if err != nil {
		if err == userdb.ErrAuthFailed {
			logger.Info("Failed to login: %v", err)
			http.Error(w, "Invalid ID or Password", http.StatusForbidden)
		} else {
			logger.Error("Unexpected error is occured in login: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	logger.Debug("Generated token: %s", token)

	// todo set response

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//w.Write(resRaw)
	logger.Info("Successfully finished LoginHandler")
}
