package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

// EncryptSymmetricKey encrypts a symmetric key using RSA public key.
func EncryptSymmetricKey(symmetricKey string, publicKeyString *rsa.PublicKey) (string, error) {
	keyBytes := []byte(symmetricKey)
	oaepLabel := []byte("")
	oaepDigest := sha256.New()

	encryptedData, err := rsa.EncryptOAEP(oaepDigest, rand.Reader, publicKeyString, keyBytes, oaepLabel)
	if err != nil {
		return "", err
	}

	// Convert the encrypted data to a string, for example, using base64 encoding
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)

	return encryptedString, nil
}

// DecryptSymmetricKey decrypts a symmetric key using RSA private key.
func DecryptSymmetricKey(encryptedSymmetricKey string, privateKeyString *rsa.PrivateKey) (string, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedSymmetricKey)
	if err != nil {
		return "", err
	}

	oaepLabel := []byte("")
	oaepDigest := sha256.New()

	decryptedData, err := rsa.DecryptOAEP(oaepDigest, rand.Reader, privateKeyString, encryptedData, oaepLabel)
	if err != nil {
		return "", err
	}

	// Convert the decrypted data to a string.
	decryptedString := string(decryptedData)

	return decryptedString, nil
}