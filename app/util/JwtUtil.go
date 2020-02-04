package util

import "golang.org/x/crypto/bcrypt"

func Encode(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
