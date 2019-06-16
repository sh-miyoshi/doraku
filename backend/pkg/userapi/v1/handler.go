package userapi

import (
	"encoding/base64"
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

// CreateUserRequestHandler creates new user with name and password
func CreateUserRequestHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call CreateUserRequestHandler method")

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
		logger.Info("User Request is not valid: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if _, err := userdb.GetInst().GetUserByName(req.Name); err == nil {
		logger.Info("User %s is already exists", req.Name)
		http.Error(w, "User is already exists", http.StatusConflict)
		return
	}

	hashedPassword := base64.StdEncoding.EncodeToString([]byte(req.Password))
	resToken, err := token.GenerateCreateUserToken(req.Name, hashedPassword)
	if err != nil {
		logger.Error("Failed to generate token %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := UserCreateResponse{
		Token: resToken,
	}
	resRaw, err := json.Marshal(res)
	if err != nil {
		logger.Error("Failed to marshal response %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resRaw)

	logger.Info("Successfully operation accepted")
}

// RegisterUserHandler registers user with validation of token
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call RegisterUserHandler method")

	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Info("Failed to decode Register User params: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	logger.Info("token: %s", req.Token)

	// Validate token
	if err := token.Authenticate(req.Token); err != nil {
		// TODO(check error e.g. token expired or invalid token format)
		logger.Info("Failed to validate token: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	userReq, _ := token.GetUserInfo(req.Token)
	logger.Debug("User Request: %v", userReq)

	registerData := userdb.UserRequest{
		Name:     userReq.Name,
		Password: userReq.HashedPassword,
	}
	if err := registerData.Validate(); err != nil {
		logger.Info("Failed to validate request data: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	if err := userdb.GetInst().CreateUser(registerData); err != nil {
		logger.Info("Failed to register data: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully operation accepted")
}

// GetUserHandler validates user id and password, and return JWT token
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetUserHandler method")

	// Validate Token in Header
	headerToken := r.Header.Get("Authorization")
	reqToken, err := token.ParseHTTPHeaderToken(headerToken)
	if err != nil {
		logger.Info("Failed to get token %v", err)
		http.Error(w, "Failed to get token", http.StatusUnauthorized)
		return
	}
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

// DeleteUserHandler delete a user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call DeleteUserHandler method")

	// Validate Token in Header
	headerToken := r.Header.Get("Authorization")
	reqToken, err := token.ParseHTTPHeaderToken(headerToken)
	if err != nil {
		logger.Info("Failed to get token %v", err)
		http.Error(w, "Failed to get token", http.StatusUnauthorized)
		return
	}

	if err := token.Authenticate(reqToken); err != nil {
		logger.Info("Failed to auth token %v", err)
		http.Error(w, "Failed to auth token", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	if err := userdb.GetInst().Delete(vars["username"]); err != nil {
		logger.Info("Failed to delete user %v", err)
		http.Error(w, "failed to delete user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	logger.Info("Successfully finished DeleteUserHandler")
}
