package utils

import "testing"

func TestValidatePhoneNumber(t *testing.T) {

	tests := []struct {
		name     string
		phone    string
		expected bool
	}{
		{
			name:     "valid fr phone number 06",
			phone:    "0601020304",
			expected: true,
		},
		{
			name:     "valid fr phone number with +33",
			phone:    "+33601020304",
			expected: true,
		},
		{
			name:     "valid fr phone number with 0033",
			phone:    "0033601020304",
			expected: true,
		},
		{
			name:     "valid fr phone number with 07",
			phone:    "0701020304",
			expected: true,
		},
		{
			name:     "valid fr phone number with 01",
			phone:    "0101020304",
			expected: true,
		},
		{
			name:     "not valid fr phone number",
			phone:    "1234567890",
			expected: false,
		},
		{
			name:     "not valid fr phone number with +",
			phone:    "+1234567890",
			expected: false,
		},
		{
			name:     "not valid fr phone number with 00",
			phone:    "001234567890",
			expected: false,
		},
		{
			name:     "not valid fr phone with americain number",
			phone:    "+11234567890",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidatePhoneNumber(tt.phone)
			if result != tt.expected {
				t.Errorf("Expected %t, but got %t", tt.expected, result)
			}
		})
	}
}
