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

	// TODO(broken file)
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

func TestGetRecommendedHobby(t *testing.T) {
	// Initalize DB
	hobbyFilePath := "../../database/hobby.csv"
	descFilePath := "../../database/description.csv"
	GetInst().Initialize(hobbyFilePath, descFilePath)

	// Test Cases
	tt := []struct {
		input         InputValue
		expectPass    bool
		expectGroupNo int64
	}{
		{InputValue{}, true, 0},
		{InputValue{Outdoor: false, Alone: false, Active: true}, true, 1},
		{InputValue{Outdoor: false, Alone: true, Active: false}, true, 2},
		{InputValue{Outdoor: false, Alone: true, Active: true}, true, 3},
		{InputValue{Outdoor: true, Alone: false, Active: false}, true, 4},
		{InputValue{Outdoor: true, Alone: false, Active: true}, true, 5},
		{InputValue{Outdoor: true, Alone: true, Active: false}, true, 6},
		{InputValue{Outdoor: true, Alone: true, Active: true}, true, 7},
		// Failed test when candidates is nothing
	}

	for _, tc := range tt {
		hobby, err := GetInst().GetRecommendedHobby(tc.input)
		if tc.expectPass {
			if err != nil {
				t.Errorf("handler should pass, but got error %v", err)
				continue
			}
			// Check GroupNo
			if hobby.GroupNo != tc.expectGroupNo {
				t.Errorf("expect groupNo: %d, but got %d", tc.expectGroupNo, hobby.GroupNo)
			}
		}
		if !tc.expectPass && err == nil {
			t.Errorf("handler should not pass")
		}
	}
}
