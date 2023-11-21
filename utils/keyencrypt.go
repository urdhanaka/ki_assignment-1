package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

// EncryptSymmetricKey encrypts a symmetric key using RSA public key.
func EncryptSymmetricKey(symmetricKey string, publicKeyString string) (string, error) {
	keyBytes := []byte(symmetricKey)

	publicKey, err := x509.ParsePKIXPublicKey([]byte(publicKeyString))
	if err != nil {
		return "", err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("failed to parse DER encoded public key")
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, keyBytes)
	if err != nil {
		return "", err
	}

	// Convert the encrypted data to a string, for example, using base64 encoding
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)

	return encryptedString, nil
}
