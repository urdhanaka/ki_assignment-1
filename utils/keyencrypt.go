package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"ki_assignment-1/dto"
)

// EncryptSymmetricKey encrypts a symmetric key using RSA public key.
func EncryptSymmetricKey(request dto.RequestedUserSymmetricKeysDTO, publicKeyString *rsa.PublicKey) (dto.EncryptedRequestedUserSymmetricKeysDTO, error) {
	var encryptedRequestedUserSymmetricKeysDto dto.EncryptedRequestedUserSymmetricKeysDTO

	secretKeyBytes := []byte(request.SecretKey)
	ivKeyBytes := []byte(request.IV)

	encryptedSecretKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKeyString, secretKeyBytes, nil)
	if err != nil {
		return dto.EncryptedRequestedUserSymmetricKeysDTO{}, err
	}

	encryptedIvKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKeyString, ivKeyBytes, nil)
	if err != nil {
		return dto.EncryptedRequestedUserSymmetricKeysDTO{}, err
	}

	// Convert the encrypted data to a string, for example, using base64 encoding
	encryptedRequestedUserSymmetricKeysDto.EncryptedSecretKey = base64.RawStdEncoding.EncodeToString(encryptedSecretKey)
	encryptedRequestedUserSymmetricKeysDto.EncryptedIVKey = base64.RawStdEncoding.EncodeToString(encryptedIvKey)

	return encryptedRequestedUserSymmetricKeysDto, nil
}

// DecryptSymmetricKey decrypts a symmetric key using RSA private key.
func DecryptSymmetricKey(request dto.EncryptedRequestedUserSymmetricKeysDTO, privateKeyString *rsa.PrivateKey) (dto.DecryptedRequestedUserSymmetricKeysDTO, error) {
	var decryptedRequestedUserSymmetricKeysDto dto.DecryptedRequestedUserSymmetricKeysDTO

	encryptedSecretKey, err := base64.RawStdEncoding.DecodeString(request.EncryptedSecretKey)
	if err != nil {
		return dto.DecryptedRequestedUserSymmetricKeysDTO{}, err
	}

	encryptedIvKey, err := base64.RawStdEncoding.DecodeString(request.EncryptedIVKey)
	if err != nil {
		return dto.DecryptedRequestedUserSymmetricKeysDTO{}, err
	}

	decryptedSecretKey, err := privateKeyString.Decrypt(nil, encryptedSecretKey, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return dto.DecryptedRequestedUserSymmetricKeysDTO{}, errors.New("1")
	}

	decryptedIvKey, err := privateKeyString.Decrypt(nil, encryptedIvKey, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return dto.DecryptedRequestedUserSymmetricKeysDTO{}, errors.New("2")
	}

	// Convert the decrypted data to a string.
	decryptedRequestedUserSymmetricKeysDto.DecryptedSecretKey = string(decryptedSecretKey)
	decryptedRequestedUserSymmetricKeysDto.DecryptedIVKey = string(decryptedIvKey)

	return decryptedRequestedUserSymmetricKeysDto, nil
}
