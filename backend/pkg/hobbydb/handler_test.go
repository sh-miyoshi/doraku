package hobbydb

import (
	"testing"
)

func TestInitialize(t *testing.T) {
	// Test correct file
	filePath := "../../database/hobby.csv"
	err := GetInst().Initialize(filePath)
	if err != nil {
		t.Errorf("cannot read correct DB file. err: %v", err)
	}

	// Test not exist file
	filePath = "not_exist.csv"
	err = GetInst().Initialize(filePath)
	if err == nil {
		t.Errorf("We expect something error, but returned nil")
	}
}
