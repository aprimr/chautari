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
func IsValidPassword(password string) (string, bool) {
	if len(password) < 8 {
		return "Password must be atleast 8 characters long", false
	}
	if hasSpace.MatchString(password) {
		return "Password cannot contain any spaces", false
	}
	if !hasLetter.MatchString(password) {
		return "Password must contain a letter", false
	}
	if !hasNumber.MatchString(password) {
		return "Password must contain a number", false
	}
	if !hasSpecial.MatchString(password) {
		return "Password must contain a special character", false
	}
	return "", true
}
