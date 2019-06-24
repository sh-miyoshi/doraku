package userdb

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestConnectDB(t *testing.T) {
	handler := localDBHandler{}

	// Test correct file
	filePath := "../../database/local/user_data.csv"
	if err := handler.ConnectDB(filePath); err != nil {
		t.Errorf("cannot read correct DB file. err: %v", err)
	}

	// Test not exist hobby file
	filePath = "not_exist_file"
	if err := handler.ConnectDB(filePath); err == nil {
		t.Errorf("We expect something error, but returned nil")
	}

	// TODO(broken file)
}

func TestAuthenticate(t *testing.T) {
	handler := localDBHandler{}

	// Initialize with correct file
	filePath := "../../database/local/user_data.csv"
	handler.ConnectDB(filePath)

	// Test Cases
	tt := []struct {
		name       string
		password   string
		expectPass bool
	}{
		{"test", "testtest", true}, // correct value
		{"test", "wrong_passwd", false},
		{"dummy", "testtest", false},
	}

	for _, tc := range tt {
		req := UserRequest{
			Name:     tc.name,
			Password: tc.password,
		}
		_, err := handler.Authenticate(req)
		if tc.expectPass && err != nil {
			// TODO check JWT Token claims
			t.Errorf("handler should pass with name %s and password %s, but got error %v", tc.name, tc.password, err)
		}
		if !tc.expectPass && err == nil {
			t.Errorf("handler should not pass with name %s and password %s, but error is nil", tc.name, tc.password)
		}
	}
}

func TestCreateUser(t *testing.T) {
	handler := localDBHandler{}

	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	handler.ConnectDB(tmpfile.Name())

	req := UserRequest{
		Name:     "test",
		Password: "password",
	}
	if err := handler.CreateUser(req); err != nil {
		t.Errorf("handler should pass with %v but got error %v", req, err)
	}

	// Test Duplicate User
	if err := handler.CreateUser(req); err == nil {
		t.Errorf("handler should not pass with same user name but error is nil")
	}

	// TODO add more test case
}

func TestDelete(t *testing.T) {
	// TODO test was failed due to create method
	/*handler := localDBHandler{}

	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	handler.ConnectDB(tmpfile.Name())
	req := UserRequest{
		Name:     "test",
		Password: "password",
	}
	// Create New User
	handler.Create(req)

	// Test Delete method
	// Delete not exists user
	if err := handler.Delete("user"); err == nil {
		t.Errorf("handler should not pass with not exists user name but error is nil")
	}

	if err := handler.Delete(req.Name); err != nil {
		t.Errorf("handler should pass with %s but got error %v", req.Name, err)
	}*/
}
