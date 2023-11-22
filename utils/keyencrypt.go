package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"ki_assignment-1/dto"
)

// EncryptSymmetricKey encrypts a symmetric key using RSA public key.
func EncryptSymmetricKey(request dto.RequestedUserSymmetricKeysDTO, publicKeyString *rsa.PublicKey) (dto.EncryptedRequestedUserSymmetricKeysDTO, error) {
	var encryptedRequestedUserSymmetricKeysDto dto.EncryptedRequestedUserSymmetricKeysDTO

	secretKeyBytes := []byte(request.SecretKey)
	ivKeyBytes := []byte(request.IV)
	oaepLabel := []byte("")
	oaepDigest := sha256.New()
	random := rand.Reader

	encryptedSecretKey, err := rsa.EncryptOAEP(oaepDigest, random, publicKeyString, secretKeyBytes, oaepLabel)
	if err != nil {
		return dto.EncryptedRequestedUserSymmetricKeysDTO{}, err
	}

	encryptedIvKey, err := rsa.EncryptOAEP(oaepDigest, random, publicKeyString, ivKeyBytes, oaepLabel)
	if err != nil {
		return dto.EncryptedRequestedUserSymmetricKeysDTO{}, err
	}

	// Convert the encrypted data to a string, for example, using base64 encoding
	encryptedRequestedUserSymmetricKeysDto.EncryptedSecretKey = base64.StdEncoding.EncodeToString(encryptedSecretKey)
	encryptedRequestedUserSymmetricKeysDto.EncryptedIVKey = base64.StdEncoding.EncodeToString(encryptedIvKey)

	return encryptedRequestedUserSymmetricKeysDto, nil
}

// DecryptSymmetricKey decrypts a symmetric key using RSA private key.
func DecryptSymmetricKey(request dto.EncryptedRequestedUserSymmetricKeysDTO, privateKeyString *rsa.PrivateKey) (dto.DecryptedRequestedUserSymmetricKeysDTO, error) {
	var decryptedRequestedUserSymmetricKeysDto dto.DecryptedRequestedUserSymmetricKeysDTO

	encryptedSecretKey, err := base64.StdEncoding.DecodeString(request.EncryptedSecretKey)
	if err != nil {
		return dto.DecryptedRequestedUserSymmetricKeysDTO{}, err
	}

	encryptedIvKey, err := base64.StdEncoding.DecodeString(request.EncryptedIVKey)
	if err != nil {
		return dto.DecryptedRequestedUserSymmetricKeysDTO{}, err
	}

	oaepLabel := []byte("")
	oaepDigest := sha256.New()

	decryptedSecretKey, err := rsa.DecryptOAEP(oaepDigest, rand.Reader, privateKeyString, encryptedSecretKey, oaepLabel)
	if err != nil {
		return dto.DecryptedRequestedUserSymmetricKeysDTO{}, err
	}

	decryptedIvKey, err := rsa.DecryptOAEP(oaepDigest, rand.Reader, privateKeyString, encryptedIvKey, oaepLabel)
	if err != nil {
		return dto.DecryptedRequestedUserSymmetricKeysDTO{}, err
	}

	// Convert the decrypted data to a string.
	decryptedRequestedUserSymmetricKeysDto.DecryptedSecretKey = string(decryptedSecretKey)
	decryptedRequestedUserSymmetricKeysDto.DecryptedIVKey = string(decryptedIvKey)

	return decryptedRequestedUserSymmetricKeysDto, nil
}
