package hobbyapi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sh-miyoshi/doraku/pkg/hobbydb"
	"github.com/sh-miyoshi/doraku/pkg/logger"
)

// GetAllHobbyHandler return lists of hobbies
func GetAllHobbyHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetAllHobbyHandler method")

	var res []HobbyKey

	for _, h := range hobbydb.GetInst().GetAllHobby() {
		tmp := HobbyKey{
			ID:   h.ID,
			Name: h.Name,
		}
		res = append(res, tmp)
	}

	resRaw, err := json.Marshal(res)
	if err != nil {
		logger.Error("Failed to marshal hobby lists %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resRaw)
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
