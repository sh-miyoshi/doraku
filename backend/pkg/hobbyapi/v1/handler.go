package hobbyapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sh-miyoshi/doraku/pkg/logger"
)

// GetAllHobbyHandler return lists of hobbies
func GetAllHobbyHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetAllHobbyHandler method")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
	logger.Info("Successfully finished GetAllHobbyHandler")
}

// GetTodayHobbyHandler return a hobby determined by date
func GetTodayHobbyHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetTodayHobbyHandler method")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
	logger.Info("Successfully finished GetTodayHobbyHandler")
}

// GetRecommendedHobbyHandler return a hobby determined by input value
func GetRecommendedHobbyHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetRecommendedHobbyHandler method")
	// TODO: parse r.Body

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
	logger.Info("Successfully finished GetRecommendedHobbyHandler")
}

// GetHobbyDetailsHandler return the details of input hobby
func GetHobbyDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	logger.Info("call GetHobbyDetailsHandler method by name: %s", name)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
	logger.Info("Successfully finished GetHobbyDetailsHandler")
}
