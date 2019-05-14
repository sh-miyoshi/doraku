package userdb

import (
	"fmt"
	"github.com/sh-miyoshi/doraku/pkg/logger"
)

type DBType int

const (
	DBRemote DBType = iota
	DBLocal
)

type userHandler interface {
	ConnectDB(connectString string) error
	// TODO Create, Delete, Authenticate
}

var instance userHandler

func InitUserHandler(dbType DBType) error {
	switch dbType {
	case DBRemote:
		logger.Info("Run User DB as Remote Mode")
		return fmt.Errorf("Sorry, not implemented yet")
	case DBLocal:
		logger.Info("Run User DB as Local Mode")
		instance = &localDBHandler{}
		return nil
	}
	return fmt.Errorf("No such database type")
}

func GetInst() userHandler {
	return instance
}
