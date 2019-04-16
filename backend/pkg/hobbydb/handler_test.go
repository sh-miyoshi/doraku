package hobbydb

import (
	"testing"
)

func TestInitialize(t *testing.T) {
	// Test correct file
	hobbyFilePath := "../../database/hobby.csv"
	descFilePath := "../../database/description.csv"
	err := GetInst().Initialize(hobbyFilePath, descFilePath)
	if err != nil {
		t.Errorf("cannot read correct DB file. err: %v", err)
	}

	// Test not exist hobby file
	hobbyFilePath = "not_exist.csv"
	descFilePath = "../../database/description.csv"
	err = GetInst().Initialize(hobbyFilePath, descFilePath)
	if err == nil {
		t.Errorf("We expect something error, but returned nil")
	}

	// Test not exist description file
	hobbyFilePath = "../../database/hobby.csv"
	descFilePath = "not_exist.csv"
	err = GetInst().Initialize(hobbyFilePath, descFilePath)
	if err == nil {
		t.Errorf("We expect something error, but returned nil")
	}
}
