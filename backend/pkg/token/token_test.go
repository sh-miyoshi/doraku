package token

import (
	"testing"
)

// TODO Testing private method
func TestValidate(t *testing.T) {
	token, _ := Generate()
	_, err := validate(token)
	if err != nil {
		t.Errorf("handler should pass with token %s, but got error %v", token, err)
	}
}

// TODO Testing private method
func TestTokenParse(t *testing.T) {
	// Test Cases
	tt := []struct {
		token      string
		expectPass bool
	}{
		{"Bearer aaaa", true},
		{"bearer aaaa", true},
		{"", false},
		{"aaaa", false},
		{"bbbb aaaa", false},
	}

	for _, tc := range tt {
		_, err := parseHTTPHeaderToken(tc.token)
		if tc.expectPass && err != nil {
			t.Errorf("handler should pass with token %s, but got error %v", tc.token, err)
		}
		if !tc.expectPass && err == nil {
			t.Errorf("handler should not pass with token %s, but error is nil", tc.token)
		}
	}
}
