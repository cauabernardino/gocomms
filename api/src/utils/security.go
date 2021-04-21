package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHash transforms the given string into hash
func GenerateHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword will compare a password with a hash an returns if they are equivalent
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
