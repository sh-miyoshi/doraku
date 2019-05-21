package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	hobbyapiv1 "github.com/sh-miyoshi/doraku/pkg/hobbyapi/v1"
	"github.com/sh-miyoshi/doraku/pkg/hobbydb"
	"github.com/sh-miyoshi/doraku/pkg/logger"
	userapiv1 "github.com/sh-miyoshi/doraku/pkg/userapi/v1"
	"github.com/sh-miyoshi/doraku/pkg/userdb"
)

func main() {
	const DefaultPort = 8080
	const DefaultBindAddr = "0.0.0.0"

	var port int
	var bindAddr string
	var logFile string
	var debug bool

	flag.IntVar(&port, "port", DefaultPort, "set port number for server")
	flag.StringVar(&bindAddr, "bind", DefaultBindAddr, "set bind address for server")
	flag.StringVar(&logFile, "logfile", "", "write log to file, output os.Stdout when do not set this")
	flag.BoolVar(&debug, "debug", false, "if true, run server as debug mode")
	flag.Parse()

	logger.InitLogger(debug, logFile)

	const hobbyFilePath = "database/hobby.csv"
	const descFilePath = "database/description.csv"

	if err := hobbydb.GetInst().Initialize(hobbyFilePath, descFilePath); err != nil {
		logger.Error("Failed to initialize Hobby DB: %v", err)
		os.Exit(1)
	}

	// TODO run db in local
	if err := userdb.InitUserHandler(userdb.DBLocal); err != nil {
		logger.Error("Failed to initialize User DB: %v", err)
		os.Exit(1)
	}

	if err := userdb.GetInst().ConnectDB("database/local_debug_user.csv"); err != nil {
		logger.Error("Failed to initialize User DB: %v", err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	// Health Check
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}).Methods("GET")

	basePath := "/api/v1"

	// Hobby API
	r.HandleFunc(basePath+"/hobby/all", hobbyapiv1.GetAllHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/today", hobbyapiv1.GetTodayHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/recommended", hobbyapiv1.GetRecommendedHobbyHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/details/{id}", hobbyapiv1.GetHobbyDetailsHandler).Methods("GET")
	r.HandleFunc(basePath+"/hobby/image/{id}", hobbyapiv1.GetImageHandler).Methods("GET")

	// User API
	r.HandleFunc(basePath+"/login", userapiv1.LoginHandler).Methods("POST")
	r.HandleFunc(basePath+"/user/{username}", userapiv1.GetUserHandler).Methods("GET")

	corsObj := handlers.AllowedOrigins([]string{"*"})

	addr := fmt.Sprintf("%s:%d", bindAddr, port)
	logger.Info("start server with %s", addr)
	if err := http.ListenAndServe(addr, handlers.CORS(corsObj)(r)); err != nil {
		logger.Error("http ListenAndServe Error: %v", err)
		os.Exit(1)
	}
}
