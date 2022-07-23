package security

import (
	"brodo-demo/service/security"

	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHash struct{}

func NewBcryptPasswordHash() security.PasswordHash {
	return &BcryptPasswordHash{}
}

func (bc *BcryptPasswordHash) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (bc *BcryptPasswordHash) IsMatch(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
