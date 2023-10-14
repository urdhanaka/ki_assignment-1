package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

func EncryptAES(plaintext string) (string, error) {
	// Retrieve the key and iv
	key := []byte(GetEnv("KEY"))
	iv := []byte(GetEnv("IV"))

	// Use padding function
	bPlaintext := PKCS5Padding([]byte(plaintext), aes.BlockSize)

	// Make new cipher key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Make a variable to hold the ciphertext
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return hex.EncodeToString(ciphertext), nil
}

func DecryptAES(ciphertext string) (string, error) {
	// Retrieve the key and iv
	key := []byte(GetEnv("KEY"))
	iv := []byte(GetEnv("IV"))

	ciphertextDecoded, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks([]byte(ciphertextDecoded), []byte(ciphertextDecoded))

	return string(ciphertextDecoded), nil
}
