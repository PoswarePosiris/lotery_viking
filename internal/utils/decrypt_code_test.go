package utils

import "testing"

func TestDecryptCode(t *testing.T) {

	tests := []struct {
		name         string
		secret       string
		secretLength int
		code         string
		expected     bool
	}{
		{
			name:         "valid code",
			secret:       "^[0-9]+$",
			secretLength: 6,
			code:         "123456",
			expected:     true,
		},
		{
			name:         "invalid code (does not match regex)",
			secret:       "^[0-9]{6}$",
			secretLength: 6,
			code:         "abcdef",
			expected:     false,
		},
		{
			name:         "invalid code (wrong length)",
			secret:       "^[0-9]$",
			secretLength: 6,
			code:         "12345",
			expected:     false,
		},
		{
			name:         "invalid code (empty)",
			secret:       "^[0-9]$",
			secretLength: 6,
			code:         "",
			expected:     false,
		},
		{
			name:         "valid code 2 letters and 6 numbers",
			secret:       "^[a-zA-Z]{2}[0-9]{6}$",
			secretLength: 8,
			code:         "ab123456",
			expected:     true,
		},
		{
			name:         "valid code 6 numbers and 2 letters",
			secret:       "^[0-9]{6}[a-zA-Z]{2}$",
			secretLength: 8,
			code:         "123456ab",
			expected:     true,
		},
		{
			name:         "invalid code 6 numbers and 2 letters",
			secret:       "^[0-9]{6}[a-zA-Z]{2}$",
			secretLength: 8,
			code:         "ab123456",
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DecryptCode(tt.secret, tt.secretLength, tt.code)
			if result != tt.expected {
				t.Errorf("Expected %t, but got %t", tt.expected, result)
			}
		})
	}
}
