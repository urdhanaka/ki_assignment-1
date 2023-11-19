package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func GenerateSecretKey() []byte {
	// Generate Secret Key
	secret := make([]byte, 32)
	rand.Read(secret)
	return secret
}

func GenerateSecretKey8Byte() []byte {
	// Generate Secret Key
	secret := make([]byte, 8)
	rand.Read(secret)
	return secret
}

func GenerateIV() []byte {
	// Generate IV
	iv := make([]byte, 16)
	rand.Read(iv)
	return iv
}

func Generate8Byte() []byte {
	// Generate 8 Byte IV
	iv := make([]byte, 8)
	rand.Read(iv)
	return iv
}

func GenerateAsymmetricKeys(id uuid.UUID) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	publicKey := &privateKey.PublicKey

	// Create keys directory
	privateKeyDirectory := "keys/private-keys"
	publicKeyDirectory := "keys/public-keys"

	if _, err = os.Stat(privateKeyDirectory); os.IsNotExist(err) {
		if err = os.MkdirAll(privateKeyDirectory, 0o777); err != nil {
			return err
		}
	}

	if _, err = os.Stat(publicKeyDirectory); os.IsNotExist(err) {
		if err = os.MkdirAll(publicKeyDirectory, 0o777); err != nil {
			return err
		}
	}

	privateKeyFilename := fmt.Sprintf("keys/private-keys/%s.pem", id)
	publicKeyFilename := fmt.Sprintf("keys/public-keys/%s.pem", id)

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	err = os.WriteFile(privateKeyFilename, privateKeyPEM, 0o644)
	if err != nil {
		return nil
	}

	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	err = os.WriteFile(publicKeyFilename, publicKeyPEM, 0o644)
	if err != nil {
		return nil
	}

	return nil
}
