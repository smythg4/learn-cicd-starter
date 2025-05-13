package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectErr   bool
	}{
		{
			name: "Valid Bearer Token",
			headers: http.Header{
				"Authorization": []string{"ApiKey mysampleapikey"},
			},
			expectedKey: "mysampleapikey",
			expectErr:   false,
		},
		{
			name:        "Missing Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectErr:   true,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)

			// Check error expectation
			if tc.expectErr && err == nil {
				t.Error("Expected an error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Did not expect an error but got: %v", err)
			}

			// Check returned key
			if got != tc.expectedKey {
				t.Errorf("Expected key %q, got %q", tc.expectedKey, got)
			}
		})
	}
}
