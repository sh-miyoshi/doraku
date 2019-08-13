package hobbyapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sh-miyoshi/doraku/pkg/hobbydb"
	"github.com/sh-miyoshi/doraku/pkg/logger"
)

// GetTodayHobbyHandler return a hobby determined by date
func GetTodayHobbyHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetTodayHobbyHandler method")

	_, month, day := time.Now().Date()
	logger.Debug("Month: %d, Day: %d", month, day)

	num := hobbydb.GetInst().GetHobbyNum()
	if num == 0 {
		logger.Error("Please Initalize DB before call API")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	todayInt := (int(month)*12 + day) % num
	logger.Debug("Today number: %d", todayInt)

	hobby, err := hobbydb.GetInst().GetHobbyByID(todayInt)
	if err != nil {
		logger.Error("Failed to get hobby %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := HobbyInfo{
		ID:   hobby.ID,
		Name: hobby.Name,
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
	logger.Info("Successfully finished GetTodayHobbyHandler")
}

// GetRecommendHobbyHandler return a hobby determined by input value
func GetRecommendHobbyHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetRecommendHobbyHandler method")

	var req RecommendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Info("Failed to decode request %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// GetRecommended Hobby
	input := hobbydb.InputValue{
		Outdoor: req.Outdoor,
		Alone:   req.Alone,
		Active:  req.Active,
	}
	hobby, err := hobbydb.GetInst().GetRecommendHobby(input)
	if err != nil {
		logger.Error("Failed to get recommended hobby %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := HobbyInfo{
		ID:   hobby.ID,
		Name: hobby.Name,
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
	logger.Info("Successfully finished GetRecommendedHobbyHandler")
}
