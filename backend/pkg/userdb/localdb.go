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
	defer file.Close()

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

func (l *localDBHandler) Authenticate(req UserRequest) (string, error) {
	data, err := csvReadAll(l.fileName)
	if err != nil {
		return "", err
	}

	for _, line := range data {
		if line[1] == req.Name {
			hashed := base64.StdEncoding.EncodeToString([]byte(req.Password))
			if hashed == line[2] {
				return token.Generate() // Generate JWT Token
			}
			logger.Info("wrong password for user: %s", req.Name)
			return "", ErrAuthFailed
		}
	}

	logger.Info("no such user %s", req.Name)
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
	lastUserID := len(data)
	return UserData{ID: lastUserID}, ErrNoSuchUser
}

func (l *localDBHandler) Create(newUser UserRequest) error {
	// User is already exists?
	user, err := l.GetUserByName(newUser.Name)
	if err == nil {
		return ErrUserAlreadyExists
	}
	// err is unexpected error
	if err != ErrNoSuchUser {
		return err
	}

	// add new user
	id := user.ID // Set New User ID

	file, err := os.OpenFile(l.fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.Error("Failed to open file %s for append new user", l.fileName)
	}
	defer file.Close()
	hashedPassword := base64.StdEncoding.EncodeToString([]byte(newUser.Password))
	fmt.Fprintf(file, "%d,%s,%s", id, newUser.Name, hashedPassword)

	logger.Info("User %s is successfully created", newUser.Name)
	return nil
}

func (l *localDBHandler) Delete(userName string) error {
	// Read All Data
	data, err := csvReadAll(l.fileName)
	if err != nil {
		return err
	}

	// Write Data if data.Name != targetUser
	file, err := os.OpenFile(l.fileName, os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("Failed to open file %s for append new user", l.fileName)
		return err
	}
	defer file.Close()

	isDeleted := false
	for _, line := range data {
		if line[1] == userName {
			// Delete Target
			isDeleted = true
		} else {
			fmt.Fprintln(file, line)
		}
	}
	if !isDeleted {
		logger.Info("no such User %s", userName)
		return ErrNoSuchUser
	}

	logger.Info("User %s is successfully delete", userName)
	return nil
}
