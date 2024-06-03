package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(PasswordHash), err
}

func VerifyPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}