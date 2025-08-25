package validate

import (
	"errors"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"regexp"
	"strings"
	"unicode"
)

func ValidatingUser(user User.User) error {
	if !IsValidUsername(user.Name) {
		return errors.New("username must contain at least 3;20 characters")
	}
	if user.Email != nil && !IsValidEmail(user.Email) {
		return errors.New("email address is invalid")
	}
	if !IsValidPassword(user.Password) {
		return errors.New("password is invalid")
	}

	return nil
}
func IsValidUsername(username string) bool {
	if len(username) <= 3 || len(username) > 20 {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username)
}
func IsValidEmail(email *string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(*email)
}
func IsValidPassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	hasLetter := false
	hasNumber := false
	hasSpecial := false
	// Определяем специальные символы
	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~"
	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
		}
		if unicode.IsDigit(char) {
			hasNumber = true
		}
		if strings.ContainsRune(specialChars, char) {
			hasSpecial = true
		}
	}
	return hasLetter && hasNumber && hasSpecial
}
