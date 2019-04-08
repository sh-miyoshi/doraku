package hobbyapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

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

	_, month, day := time.Now().Date()
	num := hobbydb.GetInst().GetHobbyNum()
	if num == 0 {
		logger.Error("Please Initalize DB before call API")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	todayInt := (int(month)*12 + day) % num

	hobby, err := hobbydb.GetInst().GetHobbyByID(todayInt)
	if err != nil {
		logger.Error("Failed to get hobby %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := HobbyKey{
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

// GetRecommendedHobbyHandler return a hobby determined by input value
func GetRecommendedHobbyHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("call GetRecommendedHobbyHandler method")

	if r.Body == nil {
		logger.Info("Request Body is nil")
        http.Error(w, "This APi requests Body parameters", http.StatusBadRequest)
        return
	}

	var userInput SelectValue
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		logger.Info("Failed to decode user request Body: %v", err)
		http.Error(w, "Request Body maybe broken", http.StatusBadRequest)
		return
	}
	defer func() {
		// Drain and close the body to let the Transport reuse the connection
		io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()
	}()

	// GetRecommended Hobby
	input := hobbydb.InputValue{
		Outdoor: userInput.Outdoor,
		Alone:   userInput.Alone,
		Active:  userInput.Active,
	}
	hobby, err := hobbydb.GetInst().GetRecommendedHobby(input)
	if err != nil {
		logger.Error("Failed to get recommended hobby %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := HobbyKey{
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

// GetHobbyDetailsHandler return the details of input hobby
func GetHobbyDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Info("Failed to get hobby %v", err)
		http.Error(w, "No such hobby", http.StatusNotFound)
		return
	}

	logger.Info("call GetHobbyDetailsHandler method by id: %d", id)

	hobby, err := hobbydb.GetInst().GetHobbyByID(id)
	if err != nil {
		logger.Info("Failed to get hobby %v", err)
		http.Error(w, "No such hobby", http.StatusNotFound)
		return
	}

	res := Hobby{
		ID:     hobby.ID,
		Name:   hobby.Name,
		NameEN: hobby.NameEN,
	}

	// TODO: set description, image and groupInfo

	resRaw, err := json.Marshal(res)
	if err != nil {
		logger.Error("Failed to marshal hobby %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resRaw)
	logger.Info("Successfully finished GetHobbyDetailsHandler")
}
