package utils

import (
	"crypto/aes"
)

func EncryptAES(key []byte, data []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(data))
	c.Encrypt(dst, data)
	return dst, nil
}

func DecryptAES(dst []byte, key []byte, data []byte) error {
	c, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	c.Decrypt(dst, data)
	return nil
}
