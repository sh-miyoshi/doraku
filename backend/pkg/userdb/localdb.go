package userdb

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"github.com/sh-miyoshi/doraku/pkg/logger"
	"github.com/sh-miyoshi/doraku/pkg/token"
	"os"
	"strconv"
)

type localDBHandler struct {
	userHandler

	fileName string
}

// This func read all csv data at once, so should not use in production
func csvReadAll(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		logger.Error("Failed to open DB file %s in Authenticate: %v", fileName, err)
		return [][]string{}, err
	}

	reader := csv.NewReader(file)
	reader.Comment = '#'

	return reader.ReadAll()
}

func (l *localDBHandler) ConnectDB(connectString string) error {
	// Check DB file exists
	_, err := os.Stat(connectString)
	if err != nil {
		return err
	}

	// TODO check file broken

	l.fileName = connectString
	return nil
}

func (l *localDBHandler) Authenticate(name string, password string) (string, error) {
	data, err := csvReadAll(l.fileName)
	if err != nil {
		return "", err
	}

	for _, line := range data {
		if line[1] == name {
			hashed := base64.StdEncoding.EncodeToString([]byte(password))
			if hashed == line[2] {
				return token.Generate() // Generate JWT Token
			}
			logger.Info("wrong password for user: %s", name)
			return "", ErrAuthFailed
		}
	}

	logger.Info("no such user %s", name)
	return "", ErrAuthFailed
}

func (l *localDBHandler) GetUserByName(name string) (UserData, error) {
	data, err := csvReadAll(l.fileName)
	if err != nil {
		return UserData{}, err
	}

	for _, line := range data {
		if line[1] == name {
			id, _ := strconv.Atoi(line[0])
			res := UserData{
				ID:   id,
				Name: name,
			}
			return res, nil
		}
	}

	logger.Info("no such user %s", name)
	return UserData{}, fmt.Errorf("no such user %s", name)
}
