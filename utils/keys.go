package utils

import (
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func GetPublicKey(id uuid.UUID) (string, error) {
	userPublicKeyFilename := fmt.Sprintf("keys/public-keys/%s.pem", id)

	publicKeyPem, err := os.ReadFile(userPublicKeyFilename)
	if err != nil {
		return "", nil
	}

	publickKeyBlock, _ := pem.Decode(publicKeyPem)

	res := base64.StdEncoding.EncodeToString(publickKeyBlock.Bytes)

	return res, nil
}

func GetPrivateKey(id uuid.UUID) (string, error) {
	userPrivateKeyFilename := fmt.Sprintf("keys/private-keys/%s.pem", id)

	privateKeyPem, err := os.ReadFile(userPrivateKeyFilename)
	if err != nil {
		return "", nil
	}

	privateKeyBlock, _ := pem.Decode(privateKeyPem)

	res := base64.StdEncoding.EncodeToString(privateKeyBlock.Bytes)

	return res, nil
}
