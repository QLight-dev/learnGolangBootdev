package main

import "unicode"

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	hasDigit := false
	hasCapital := false

	for _, r := range password {
		if unicode.IsDigit(r) {
			hasDigit = true
		}
		if unicode.IsUpper(r) {
			hasCapital = true
		}
	}

	return hasDigit && hasCapital
}
