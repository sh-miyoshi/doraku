package userdb

import (
	"encoding/base64"
	"encoding/csv"
	"github.com/sh-miyoshi/doraku/pkg/logger"
	"github.com/sh-miyoshi/doraku/pkg/token"
	"io"
	"os"
	"path/filepath"
)

const (
	userDatabaseFile        = "user_data.csv"
	temporaryCreateUserFile = "temp_user.csv"
)

type localDBHandler struct {
	UserHandler

	userFileName string
	tempFileName string
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

// ConnectDB check file exists(connectString is a directory of files)
func (l *localDBHandler) ConnectDB(connectString string) error {
	l.userFileName = filepath.Join(connectString, "/", userDatabaseFile)
	l.tempFileName = filepath.Join(connectString, "/", temporaryCreateUserFile)

	// Check DB file exists
	if _, err := os.Stat(l.userFileName); err != nil {
		return err
	}
	if _, err := os.Stat(l.tempFileName); err != nil {
		return err
	}

	// TODO check file broken

	return nil
}

func (l *localDBHandler) Authenticate(req UserRequest) (string, error) {
	data, err := csvReadAll(l.userFileName)
	if err != nil {
		return "", err
	}

	for _, line := range data {
		if line[0] == req.Name {
			hashed := base64.StdEncoding.EncodeToString([]byte(req.Password))
			if hashed == line[1] {
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
	data, err := csvReadAll(l.userFileName)
	if err != nil {
		return UserData{}, err
	}

	for _, line := range data {
		if line[0] == name {
			res := UserData{
				Name: name,
			}
			return res, nil
		}
	}

	logger.Info("no such user %s", name)
	return UserData{}, ErrNoSuchUser
}

func (l *localDBHandler) CreateUserValidation(newUser UserRequest) error {
	// User is already exists?
	_, err := l.GetUserByName(newUser.Name)
	if err == nil {
		return ErrUserAlreadyExists
	}
	// err is unexpected error
	if err != ErrNoSuchUser {
		return err
	}

	// TODO(check temp file, if ok, add to tempfile)

	//// add new user
	//file, err := os.OpenFile(l.fileName, os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	logger.Error("Failed to open file %s for append new user", l.fileName)
	//}
	//defer file.Close()
	//hashedPassword := base64.StdEncoding.EncodeToString([]byte(newUser.Password))
	//fmt.Fprintf(file, "%s,%s", newUser.Name, hashedPassword)

	logger.Info("User %s is successfully created", newUser.Name)
	return nil
}

func (l *localDBHandler) Delete(userName string) error {
	var data [][]string

	file, err := os.OpenFile(l.userFileName, os.O_RDWR, 0644)
	if err != nil {
		logger.Error("Failed to open DB file %s in Authenticate: %v", l.userFileName, err)
		return err
	}
	defer file.Close()

	isDeleted := false
	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Error("Failed to read data: %v", err)
			return err
		}
		if line[0] == userName {
			isDeleted = true
		} else {
			data = append(data, line)
		}
	}

	// Remove All data at first
	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	writer.WriteAll(data)

	if !isDeleted {
		logger.Info("no such user %s", userName)
		return ErrNoSuchUser
	}

	logger.Info("User %s is successfully delete", userName)
	return nil
}
