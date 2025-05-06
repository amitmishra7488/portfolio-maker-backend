package utils

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// CreateFullName ensures a properly formatted full name
func CreateFullName(firstName, lastName string) (string, error) {
	// Trim leading & trailing spaces
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	// Validate names (optional: only allow alphabets)
	if !isValidName(firstName) {
		firstName = ""
	}
	if !isValidName(lastName) {
		lastName = ""
	}

	// Handle different cases
	if firstName == "" && lastName == "" {
		return "", errors.New("Invalid Name")
	} else if firstName == "" {
		return "", errors.New("Invalid Name")
	} else if lastName == "" {
		return firstName, nil
	}

	return firstName + " " + lastName, nil
}

// isValidName checks if the name contains only alphabets (optional)
func isValidName(name string) bool {
	regex := regexp.MustCompile(`^[A-Za-z]+$`)
	return regex.MatchString(name) || name == ""
}

// HashPassword hashes a plain-text password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash verifies a plain-text password against a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // Returns true if password is correct
}
