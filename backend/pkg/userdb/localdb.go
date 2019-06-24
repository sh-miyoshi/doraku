package userdb

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"github.com/sh-miyoshi/doraku/pkg/logger"
	"github.com/sh-miyoshi/doraku/pkg/token"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	myHobbySeparator = ";"
)

type localDBHandler struct {
	UserHandler

	userFileName string
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

// ConnectDB check file exists(connectString is a file path of user data)
func (l *localDBHandler) ConnectDB(connectString string) error {

	// Check DB file exists
	if _, err := os.Stat(connectString); err != nil {
		return err
	}

	// TODO check file broken

	l.userFileName = connectString
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
			strHobbies := []string{}
			if len(line) >= 3 {
				strHobbies = strings.Split(line[2], myHobbySeparator)
			}
			var hobbies []int
			for _, h := range strHobbies {
				hobby, _ := strconv.Atoi(h)
				hobbies = append(hobbies, hobby)
			}
			res := UserData{
				Name:        name,
				MyHobbyList: hobbies,
			}
			return res, nil
		}
	}

	return UserData{}, ErrNoSuchUser
}

func (l *localDBHandler) CreateUser(newUser UserRequest) error {
	// User is already exists?
	_, err := l.GetUserByName(newUser.Name)
	if err == nil {
		return ErrUserAlreadyExists
	}
	// err is unexpected error
	if err != ErrNoSuchUser {
		return err
	}

	// add new user
	file, err := os.OpenFile(l.userFileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.Error("Failed to open file %s for append new user", l.userFileName)
	}
	defer file.Close()
	fmt.Fprintf(file, "%s,%s", newUser.Name, newUser.Password)

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

func (l *localDBHandler) AddMyHobby(userName string, hobbyID int) error {
	var data [][]string

	file, err := os.OpenFile(l.userFileName, os.O_RDWR, 0644)
	if err != nil {
		logger.Error("Failed to open DB file %s in Authenticate: %v", l.userFileName, err)
		return err
	}
	defer file.Close()

	isAdded := false
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
			isAdded = true
			// add hobby data
			strHobby := strconv.Itoa(hobbyID)
			if len(line) <=2 || len(line[2]) == 0 {
				line[2] = strHobby
			} else {
				duplicated := false
				for _, h := range strings.Split(line[2], myHobbySeparator) {
					if h == strHobby {
						duplicated = true
						break
					}
				}
				if !duplicated {
					line[2] += myHobbySeparator + strHobby
				}
			}
		}
		data = append(data, line)
	}

	// Remove All data at first
	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	writer.WriteAll(data)

	if !isAdded {
		logger.Info("no such user %s", userName)
		return ErrNoSuchUser
	}

	return nil
}
