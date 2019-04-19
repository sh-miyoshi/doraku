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

func TestGetHobbyByID(t *testing.T) {
	// Initalize DB
	hobbyFilePath := "../../database/hobby.csv"
	descFilePath := "../../database/description.csv"
	GetInst().Initialize(hobbyFilePath, descFilePath)

	// Test Cases
	tt := []struct {
		id         int
		expectPass bool
	}{
		{0, true},
		{21, true},
		{22, false},
		{-1, false},
	}

	for _, tc := range tt {
		_, err := GetInst().GetHobbyByID(tc.id)
		if tc.expectPass && err != nil {
			t.Errorf("handler should pass with id %d, but got error %v", tc.id, err)
		}
		if !tc.expectPass && err == nil {
			t.Errorf("handler should not pass with id %d, but error is nil", tc.id)
		}
	}
}
