package userdb

import (
	"encoding/base64"
	"encoding/csv"
	"github.com/sh-miyoshi/doraku/pkg/logger"
	"os"
)

type localDBHandler struct {
	userHandler

	fileName string
}

func (l *localDBHandler) ConnectDB(connectString string) error {
	// Check DB file exists
	_, err := os.Stat(connectString)
	if err != nil {
		return err
	}
	l.fileName = connectString
	return nil
}

func (l *localDBHandler) Authenticate(id string, password string) error {
	file, err := os.Open(l.fileName)
	if err != nil {
		logger.Error("Failed to open DB file %s in Authenticate: %v", l.fileName, err)
		return err
	}

	reader := csv.NewReader(file)
	reader.Comment = '#'

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		if line[0] == id {
			hashed := base64.StdEncoding.EncodeToString([]byte(password))
			if hashed == line[1] {
				return nil
			}
			logger.Info("wrong password for id: %s", id)
			return ErrAuthFailed
		}
	}

	logger.Info("no such id %s", id)
	return ErrAuthFailed
}
