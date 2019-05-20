package userapi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sh-miyoshi/doraku/pkg/logger"
	"github.com/sh-miyoshi/doraku/pkg/token"
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

	token, err := userdb.GetInst().Authenticate(req.Name, req.Password)

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

	res := LoginResponse{
		Token: token,
	}

	resRaw, err := json.Marshal(res)
	if err != nil {
		logger.Error("Failed to marshal hobby %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resRaw)
	logger.Info("Successfully finished LoginHandler")
}

// GetUserHandler validates user id and password, and return JWT token
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetUserHandler method")

	// Validate Token in Header
	reqTokenStr := r.Header.Get("Authorization")
	tokenStr, err := token.ParseHTTPHeaderToken(reqTokenStr)
	if err != nil {
		logger.Info("Failed to get JWT token %v", err)
		http.Error(w, "Cannot find JWT Token in Header", http.StatusBadRequest)
		return
	}
	claims, err := token.Validate(tokenStr)
	if err != nil {
		logger.Info("Failed to auth token %v", err)
		http.Error(w, "Failed to auth token", http.StatusBadRequest)
		return
	}
	logger.Debug("claims in token: %v", claims)
	// TODO validate claims(e.g. expired time, ...)

	vars := mux.Vars(r)
	user, err := userdb.GetInst().GetUserByName(vars["username"])
	if err != nil {
		// TODO: check error
		logger.Info("Failed to get user %v", err)
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}

	res := User{
		ID:    user.ID,
		Name:  user.Name,
		EMail: user.EMail,
	}

	resRaw, err := json.Marshal(res)
	if err != nil {
		logger.Error("Failed to marshal hobby %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resRaw)
	logger.Info("Successfully finished GetUserHandler")
}
