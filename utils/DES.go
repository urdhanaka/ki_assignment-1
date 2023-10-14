package utils

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
)

func EncryptDES(plaintext string) (string, error) {
	key := []byte(GetEnv("KEY8"))
	iv := []byte(GetEnv("IV8"))

	// Make new DES cipher key
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Use padding to plaintext
	bPlaintext := PKCS5Padding([]byte(plaintext), block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(bPlaintext))
	mode.CryptBlocks(ciphertext, bPlaintext)

	return hex.EncodeToString(ciphertext), nil
}

func DecryptDES(ciphertext string) (string, error) {
	key := []byte(GetEnv("KEY8"))
	iv := []byte(GetEnv("IV8"))

	ciphertextDecoded, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	originalData := make([]byte, len(ciphertextDecoded))
	mode.CryptBlocks(originalData, []byte(ciphertextDecoded))
	originalData = PKCS5Unpadding(originalData)

	return hex.EncodeToString(originalData), nil
}
