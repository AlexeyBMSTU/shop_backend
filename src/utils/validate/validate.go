package validate

import (
	"errors"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"regexp"
	"strings"
)

func ValidatingUser(user User.User) error {
	if !IsValidUsername(user.Username) {
		return errors.New("username must contain at least 3;6 characters")
	}
	if !IsValidEmail(user.Email) {
		return errors.New("email address is invalid")
	}
	if !IsValidPassword(user.Password) {
		return errors.New("password must contain at least 8 characters and at least one letter and one number")
	}

	return nil
}
func IsValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username)
}
func IsValidEmail(email string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasLetter := false
	hasNumber := false
	for _, char := range password {
		if strings.ContainsAny(string(char), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			hasLetter = true
		}
		if strings.ContainsAny(string(char), "0123456789") {
			hasNumber = true
		}
	}
	return hasLetter && hasNumber
}
