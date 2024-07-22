package helper

import (
	"strings"
)

func ValidateInput(firstName string, lastName string, email string) (bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	return isValidName, isValidEmail
}
