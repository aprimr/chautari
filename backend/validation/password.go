package validation

import (
	"regexp"
)

var (
	hasLetter  = regexp.MustCompile(`[A-Za-z]`)
	hasNumber  = regexp.MustCompile(`\d`)
	hasSpecial = regexp.MustCompile(`[^A-Za-z\d]`)
)

func IsValidPassword(password string) bool {
	if len(password) < 8 {
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
