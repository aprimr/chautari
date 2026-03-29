package validation

import "regexp"

var (
	validChar = regexp.MustCompile(`^[a-zA-Z0-9_.]+$`)
)

// Must be greater than or equal to 8 characters
// Must contain a letter
// Must not contain any spaces
// Can have `_`, `.`
func IsValidUsername(username string) (string, bool) {
	if len(username) < 8 {
		return "Username must be at least 8 characters long", false
	}

	if hasSpace.MatchString(username) {
		return "Username cannot contain spaces", false
	}

	if !hasLetter.MatchString(username) {
		return "Username must contain at least one letter", false
	}

	if !validChar.MatchString(username) {
		return "Username can only contain letters, numbers, '_' and '.'", false
	}

	return "", true
}
