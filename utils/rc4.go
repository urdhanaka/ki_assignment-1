package utils

import (
	"crypto/rc4"
	"encoding/base64"
)

// RC4 is cryptographically broken and should not be used for secure
// applications.

func EncryptRC4(data string) (string, error) {
	key := []byte(GetEnv("KEY"))

	dataBytes := []byte(data)

	c, err := rc4.NewCipher(key)
	if err != nil {
		return "", err
	}
	dst := make([]byte, len(dataBytes))
	c.XORKeyStream(dst, dataBytes)

	return base64.StdEncoding.EncodeToString(dst), nil
}

func DecryptRC4(encrypted string) (string, error) {
	key := []byte(GetEnv("KEY"))

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return "", err
	}
	encryptedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(encryptedBytes))
	cipher.XORKeyStream(decrypted, encryptedBytes)

	return string(decrypted), err
}