package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sh-miyoshi/doraku/pkg/hobby"
	"github.com/sh-miyoshi/doraku/pkg/logger"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/hobby", hobby.GetHobbyHandler).Methods("GET")

	logger.Info("start server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("http ListenAndServe Error: %v", err)
		os.Exit(1)
	}
}
