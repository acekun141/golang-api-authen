package util

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(password string, hashedPassword string) bool {
	bytePassword := []byte(password)
	byteHash := []byte(hashedPassword)
	if err := bcrypt.CompareHashAndPassword(byteHash, bytePassword); err != nil {
		return false
	}
	return true
}
