package utils

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

func EncryptDES(plaintext []byte, secretKeyParam string, ivParam string) (string, error) {
	key := []byte(secretKeyParam)
	iv := []byte(ivParam)

	// Make new DES cipher key
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Use padding to plaintext
	bPlaintext := PKCS5Padding(plaintext, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(bPlaintext))
	mode.CryptBlocks(ciphertext, bPlaintext)

	res := base64.StdEncoding.EncodeToString(ciphertext)

	return res, nil
}

func DecryptDES(ciphertext string, secretKeyParam string, ivParam string) (string, error) {
	key := []byte(secretKeyParam)
	iv := []byte(ivParam)

	ciphertextDecoded, err := base64.StdEncoding.DecodeString(ciphertext)
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

	return string(originalData), nil
}
