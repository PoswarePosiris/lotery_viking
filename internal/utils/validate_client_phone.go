package utils

import "regexp"

// Check if the phone number is valid or not
func ValidatePhoneNumber(clientPhone string) bool {
	// French phone number pattern: starts with +33, followed by 9 digits, or starts with 06 or 07, followed by 8 digits
	pattern := `^(?:(?:\+|00)33|0)\s*[1-9](?:[\s.-]*\d{2}){4}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(clientPhone)
}
