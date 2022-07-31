package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPasswrod, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswrod), []byte(password))

	return err == nil
}
