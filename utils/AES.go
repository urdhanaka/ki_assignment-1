package utils

import (
	"crypto/aes"

	"golang.org/x/crypto/bcrypt"
)

const key = "vzhLV1DZPKR9ttZa2NBQN1DD1Y7W9NXk"

func Encrypt(String string) (string, error) {
	c, err := aes.NewCipher([]byte(key))

	return string(bytes), err
}

func Decrypt(String string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(String), 4)

	return string(bytes), err
}
