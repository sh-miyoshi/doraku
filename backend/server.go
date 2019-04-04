package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sh-miyoshi/doraku/pkg/hobbyapi"
	"github.com/sh-miyoshi/doraku/pkg/hobbydb"
	"github.com/sh-miyoshi/doraku/pkg/logger"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/hobby", hobbyapi.GetHobbyHandler).Methods("GET")
	r.HandleFunc("/api/v1/hobby/group/{groupNo}", hobbyapi.GetHobbyByGroupNoHandler).Methods("GET")
	r.HandleFunc("/api/v1/hobby/id/{id}", hobbyapi.GetHobbyByIDHandler).Methods("GET")

	if err := hobbydb.GetInst().Initialize("mongodb://localhost/"); err != nil {
		logger.Error("Failed to connect Mongo DB: %v", err)
        os.Exit(1)
	}

	logger.Info("start server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("http ListenAndServe Error: %v", err)
		os.Exit(1)
	}
}
