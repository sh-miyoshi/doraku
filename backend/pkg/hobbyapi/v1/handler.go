package hobbyapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	// if query == yes then, set var = true
	outdoor := (r.FormValue("outdoor") == "yes")
	alone := (r.FormValue("alone") == "yes")
	active := (r.FormValue("active") == "yes")

	// GetRecommended Hobby
	input := hobbydb.InputValue{
		Outdoor: outdoor,
		Alone:   alone,
		Active:  active,
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

	desc := ""
	// Set description if hobby.Description includes http(it mean http[s]://...)
	if strings.Contains(hobby.Description, "http") {
		desc = hobby.Description
	}

	res := Hobby{
		ID:          hobby.ID,
		Name:        hobby.Name,
		NameEN:      hobby.NameEN,
		Description: desc,
	}

	// TODO: set groupInfo

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

// GetImageHandler return the image binary of hobby
func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Info("Failed to get image %v", err)
		http.Error(w, "No such hobby", http.StatusNotFound)
		return
	}

	logger.Info("call GetImageHandler method by id: %d", id)

	filename := fmt.Sprintf("database/images/%d.png", id)

	// Check file exists
	_, err = os.Stat(filename)
	if err != nil {
		logger.Info("Failed to get image file by id(%d) %v", id, err)
		http.Error(w, "No such hobby", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filename)

	logger.Info("Successfully finished GetImageHandler")
}
