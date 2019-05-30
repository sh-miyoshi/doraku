package userdb

import (
	"errors"
)

const (
	// NameLengthMin is minimum length of name
	NameLengthMin = 4
	// NameLengthMax is maximum length of name
	NameLengthMax = 32
	// PasswordLengthMin is minimum length of password
	PasswordLengthMin = 8
	// PasswordLengthMax is maximum length of password
	PasswordLengthMax = 128
)

var (
	// ErrAuthFailed is an error for authentication failed
	ErrAuthFailed = errors.New("Failed to authenticate")
	// ErrUserAlreadyExists is an error for user is already
	ErrUserAlreadyExists = errors.New("User is already exists")
	// ErrNoSuchUser is an error for no such user
	ErrNoSuchUser = errors.New("No such user")
)

// UserData is data for user
type UserData struct {
	Name string
	// TODO(myHobbyList, ...)
}

//UserRequest is a request param for user method
type UserRequest struct {
	Name     string
	Password string
}

// Validate method validates UserRequest
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
