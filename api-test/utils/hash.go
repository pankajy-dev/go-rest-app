package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}
