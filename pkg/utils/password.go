package utils

import "regexp"

// regex for password validation
// password must have at least 8 characters, one uppercase letter, one digit, and one special character
var passwordRegex = `^(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`

func ValidatePassword(password string) bool {
	return regexp.MustCompile(passwordRegex).MatchString(password)
}
