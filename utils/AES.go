package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(src []byte) []byte {
	length := len(src)

	if length == 0 {
		return nil
	}
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

func EncryptAES(plaintext []byte, secretKeyParam string, ivParam string) (string, error) {
	// Retrieve the key and iv
	key := []byte(secretKeyParam)
	iv := []byte(ivParam)

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

	res := base64.StdEncoding.EncodeToString(ciphertext)

	return res, nil
}

func DecryptAES(ciphertext string, secretKeyParam string, ivParam string) (string, error) {
	// Retrieve the key and iv
	key := []byte(secretKeyParam)
	iv := []byte(ivParam)

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

func EncryptAESFile(plainfile []byte, secretKeyParam string, ivParam string) ([]byte, error) {
	// Retrieve the key and iv
	key := []byte(secretKeyParam)
	iv := []byte(ivParam)

	// Use padding function
	bplainfile := PKCS5Padding(plainfile, aes.BlockSize)

	// Make new cipher key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Make a variable to hold the ciphertext
	cipherfile := make([]byte, len(bplainfile))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherfile, bplainfile)

	return cipherfile, nil
}

func DecryptAESFile(cipherfile []byte, secretKeyParam string, ivParam string) ([]byte, error) {
	// Retrieve the key and iv
	key := []byte(secretKeyParam)
	iv := []byte(ivParam)

	// Decode the base64 cipherfile

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("AES cipher creation error: %v", err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherfile, cipherfile)

	// Unpad the byte
	cipherfile = PKCS5Unpadding(cipherfile)

	return cipherfile, nil
}
