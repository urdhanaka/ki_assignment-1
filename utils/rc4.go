package utils

import (
	"crypto/rc4"
	"encoding/base64"

	"time"
)

// RC4 is cryptographically broken and should not be used for secure
// applications.

func EncryptRC4(data []byte) (string, error) {

	elapsedTime := timer("RC4-Encrypt")
	defer elapsedTime()
	time.Sleep(1 * time.Second)
	
	key := []byte(GetEnv("KEY"))

	c, err := rc4.NewCipher(key)
	if err != nil {
		return "", err
	}
	dst := make([]byte, len(data))
	c.XORKeyStream(dst, data)

	res := base64.StdEncoding.EncodeToString(dst)

	return res, nil
}

func DecryptRC4(encrypted string) (string, error) {
	
	elapsedTime := timer("RC4-Decrypt")
	defer elapsedTime()
	time.Sleep(1 * time.Second)
	
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