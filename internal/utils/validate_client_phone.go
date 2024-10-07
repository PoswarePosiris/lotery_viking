package utils

import "regexp"

// ValidatePhoneNumber checks if the given phone number is valid based on French and international formats.
// Valid formats:
// - Starts with +XX or 00XX followed by 9 to 13 digits (for international numbers)
// - Starts with 06 or 07 followed by 8 digits (for French local numbers)
func ValidatePhoneNumber(clientPhone string) bool {
	// Pattern for French and international phone numbers
	pattern := `^(?:0[67]\d{8})$|^(?:(?:\+|00)[1-9]\d{1,3}\d{6,12})$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(clientPhone)
}
