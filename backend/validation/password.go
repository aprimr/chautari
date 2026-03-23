package validation

import (
	"regexp"
)

var (
	hasLetter  = regexp.MustCompile(`[A-Za-z]`)
	hasNumber  = regexp.MustCompile(`\d`)
	hasSpecial = regexp.MustCompile(`[^A-Za-z\d]`)
	hasSpace   = regexp.MustCompile(`\s`)
)

// Must be greater than or equal to 8 characters
// Must contain a letter
// Must contain a number
// Must contain a special character
// Must not contain any spaces
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	if hasSpace.MatchString(password) {
		return false
	}
	if !hasLetter.MatchString(password) {
		return false
	}
	if !hasNumber.MatchString(password) {
		return false
	}
	if !hasSpecial.MatchString(password) {
		return false
	}
	return true
}
