package userdb

import (
	"errors"
)

var (
	ErrAuthFailed = errors.New("Failed to authenticate")
)

type UserData struct {
	ID   int
	Name string
}

type UserRequest struct {
	Name     string
	Password string
}
