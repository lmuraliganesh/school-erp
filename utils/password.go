package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plain text password and returns its bcrypt hash
func HashPassword(password string) (string, error) {
	// 1. Call bcrypt.GenerateFromPassword()
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	// 2. If err != nil, return "", err
	if err != nil {
		return "", err
	}
	
	// 3. Otherwise, return string(hashedBytes), nil
	return string(hashedBytes), nil
}

// CheckPasswordHash compares a plain text password against a stored hash
func CheckPasswordHash(password, hash string) bool {
	// 1. Call bcrypt.CompareHashAndPassword()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	
	// 2. Return true if err is nil, false otherwise
	return err == nil
}