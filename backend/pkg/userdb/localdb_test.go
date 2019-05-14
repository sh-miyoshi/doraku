package userdb

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	handler := localDBHandler{}

	// Test correct file
	filePath := "../../database/local_debug_user.csv"
	if err := handler.ConnectDB(filePath); err != nil {
		t.Errorf("cannot read correct DB file. err: %v", err)
	}

	// Test not exist hobby file
	filePath = "not_exist.csv"
	if err := handler.ConnectDB(filePath); err == nil {
		t.Errorf("We expect something error, but returned nil")
	}

	// TODO(broken file)
}

func TestAuthenticate(t *testing.T) {
	handler := localDBHandler{}

	// Initialize with correct file
	filePath := "../../database/local_debug_user.csv"
	handler.ConnectDB(filePath)

	// Test with correct value
	if err := handler.Authenticate("test", "testtest"); err != nil {
		t.Errorf("Failed to auth correct data: %v", err)
	}

	// Test with incorrect value
	if err := handler.Authenticate("test", "wrong_passwd"); err == nil {
		t.Errorf("Success to auth incorrect data: %v", err)
	}

	// Test with not exists user
	if err := handler.Authenticate("dummy", "testtest"); err == nil {
		t.Errorf("Success to auth incorrect data: %v", err)
	}
}
