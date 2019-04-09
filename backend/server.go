package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	apiv1 "github.com/sh-miyoshi/doraku/pkg/hobbyapi/v1"
	"github.com/sh-miyoshi/doraku/pkg/hobbydb"
	"github.com/sh-miyoshi/doraku/pkg/logger"
)

func main() {
	const DefaultPort = 8080
	const DefaultBindAddr = "0.0.0.0"

	var port int
	var bindAddr string

	flag.IntVar(&port, "port", DefaultPort, "set port number for server")
	flag.StringVar(&bindAddr, "bind", DefaultBindAddr, "set bind address for server")
	flag.Parse()

	// If you run doraku-server as debug mode, uncommentout following line
	//logger.InitLogger(true)

	r := mux.NewRouter()

	basePath := "/api/v1"
	r.HandleFunc(basePath+"/hobby/all", apiv1.GetAllHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/today", apiv1.GetTodayHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/recommended", apiv1.GetRecommendedHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/details/{id}", apiv1.GetHobbyDetailsHandler).Methods("GET")

	const filePath = "database/hobby.csv"

	if err := hobbydb.GetInst().Initialize(filePath); err != nil {
		logger.Error("Failed to initialize DB: %v", err)
		os.Exit(1)
	}

	addr := fmt.Sprintf("%s:%d", bindAddr, port)
	logger.Info("start server with %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Error("http ListenAndServe Error: %v", err)
		os.Exit(1)
	}
}
