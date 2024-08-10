package utils

import (
	"regexp"
	"golang.org/x/crypto/bcrypt"
)



func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	return re.MatchString(email)
}

// password must contain at least 1 uppercase, 1 lowercase letter, 1 digit, 1 special character
// password minimum length is 8
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	uppercase := regexp.MustCompile(`[A-Z]`)
	if !uppercase.MatchString(password) {
		return false
	}

	lowercase := regexp.MustCompile(`[a-z]`)
	if !lowercase.MatchString(password) {
		return false
	}

	digit := regexp.MustCompile(`[0-9]`)
	if !digit.MatchString(password) {
		return false
	}

	special := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)

	return special.MatchString(password)
}

func ComparePasswordHashes(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}