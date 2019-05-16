package userdb

import (
	"errors"
)

var (
	ErrAuthFailed = errors.New("Failed to authenticate")
)
