package utils

import (
	"crypto/rc4"
	"encoding/hex"
)

// RC4 is cryptographically broken and should not be used for secure
// applications.

func EncryptRC4(key []byte, data []byte) (string, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return "", err
	}
	dst := make([]byte, len(data))
	c.XORKeyStream(dst, data)

	return hex.EncodeToString(dst), nil
}

func DecryptRC4(decrypted []byte, key []byte, encrypted string) error {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return err
	}
	encryptedBytes, err := hex.DecodeString(string(encrypted))
	if err != nil {
		return err
	}
	cipher.XORKeyStream(decrypted, encryptedBytes)

	return nil
}