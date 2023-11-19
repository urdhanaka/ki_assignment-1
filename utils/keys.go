package utils

import (
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func GetPublicKey(id string) (string, error) {
	userPublicKeyFilename := fmt.Sprintf("keys/public-keys/%s.pem", id)

	publicKeyPem, err := os.ReadFile(userPublicKeyFilename)
	if err != nil {
		return "", nil
	}

	publickKeyBlock, _ := pem.Decode(publicKeyPem)

	res := base64.StdEncoding.EncodeToString(publickKeyBlock.Bytes)

	return res, nil
}
