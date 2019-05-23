package userdb

import (
	"errors"
)

const (
	NameLengthMin     = 4
	NameLengthMax     = 32
	PasswordLengthMin = 8
	PasswordLengthMax = 128
)

var (
	ErrAuthFailed        = errors.New("Failed to authenticate")
	ErrUserAlreadyExists = errors.New("User is already exists")
	ErrNoSuchUser        = errors.New("No such user")
)

type UserData struct {
	ID   int
	Name string
}

type UserRequest struct {
	Name     string
	Password string
}

func (r *UserRequest) Validate() error {
	if len(r.Name) < NameLengthMin {
		return errors.New("Name Length is too small")
	}
	if len(r.Name) > NameLengthMax {
		return errors.New("Name Length is too long")
	}

	if len(r.Password) < PasswordLengthMin {
		return errors.New("Password Length is too small")
	}
	if len(r.Password) > PasswordLengthMax {
		return errors.New("Password Length is too long")
	}

	return nil
}
