package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	apiv1 "github.com/sh-miyoshi/doraku/pkg/hobbyapi/v1"
	"github.com/sh-miyoshi/doraku/pkg/hobbydb"
	"github.com/sh-miyoshi/doraku/pkg/logger"
)

func main() {
	// If you run doraku-server as debug mode, uncommentout following line
	//logger.InitLogger(true)

	r := mux.NewRouter()

	basePath := "/api/v1"
	r.HandleFunc(basePath+"/hobby/all", apiv1.GetAllHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/today", apiv1.GetTodayHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/recommended", apiv1.GetRecommendedHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/details/{name}", apiv1.GetHobbyDetailsHandler).Methods("GET")

	if err := hobbydb.GetInst().Initialize(); err != nil {
		logger.Error("Failed to initialize DB: %v", err)
		os.Exit(1)
	}

	logger.Info("start server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("http ListenAndServe Error: %v", err)
		os.Exit(1)
	}
}
