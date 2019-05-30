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

	userReq := userdb.UserRequest{
		Name:     req.Name,
		Password: req.Password,
	}

	if err := userReq.Validate(); err != nil {
		logger.Info("User Request iis not valid: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	token, err := userdb.GetInst().Authenticate(userReq)

	if err != nil {
		if err == userdb.ErrAuthFailed {
			logger.Info("Failed to login: %v", err)
			http.Error(w, "Invalid Name or Password", http.StatusForbidden)
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

// CreateUserHandler creates new user with name and password
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call CreateUserHandler method")

	var req UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Info("Failed to decode Create User params: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	userReq := userdb.UserRequest{
		Name:     req.Name,
		Password: req.Password,
	}

	if err := userReq.Validate(); err != nil {
		logger.Info("User Request iis not valid: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if err := userdb.GetInst().Create(userReq); err != nil {
		logger.Info("Failed to create new user %s: %v", userReq.Name, err)
		if err == userdb.ErrUserAlreadyExists {
			http.Error(w, "User is already exists", http.StatusConflict)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	logger.Info("Successfully finished CreateUserHandler")
}

// GetUserHandler validates user id and password, and return JWT token
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetUserHandler method")

	// Validate Token in Header
	reqToken := r.Header.Get("Authorization")
	if err := token.Authenticate(reqToken); err != nil {
		logger.Info("Failed to auth token %v", err)
		http.Error(w, "Failed to auth token", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	user, err := userdb.GetInst().GetUserByName(vars["username"])
	if err != nil {
		// TODO: check error
		logger.Info("Failed to get user %v", err)
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}

	res := User{
		Name: user.Name,
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
