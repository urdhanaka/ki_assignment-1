package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
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

func EncryptAES(plaintext []byte) (string, error) {
	// Retrieve the key and iv
	key := []byte(GetEnv("KEY"))
	iv := []byte(GetEnv("IV"))

	// Use padding function
	bPlaintext := PKCS5Padding(plaintext, aes.BlockSize)

	// Make new cipher key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Make a variable to hold the ciphertext
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES(ciphertext string) (string, error) {
	// Retrieve the key and iv
	key := []byte(GetEnv("KEY"))
	iv := []byte(GetEnv("IV"))

	// Decode the base64 ciphertext
	ciphertextDecoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 decoding error: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("AES cipher creation error: %v", err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks([]byte(ciphertextDecoded), []byte(ciphertextDecoded))

	// Unpad the byte
	ciphertextDecoded = PKCS5Unpadding(ciphertextDecoded)

	return string(ciphertextDecoded), nil
}
