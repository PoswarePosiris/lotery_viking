package utils

import (
	"regexp"
)

func DecryptCode(secret string, secretLength int, code string) bool {
	regex, err := regexp.Compile(secret)
	if err != nil {
		panic(err)
	}

	// Check if the code respects the secret regex
	if !regex.MatchString(code) {
		return false
	}

	// Check if the code has the correct length
	if len(code) != secretLength {
		return false
	}

	return true
}
