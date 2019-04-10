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
}
