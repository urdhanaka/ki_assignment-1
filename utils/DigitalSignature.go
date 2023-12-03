package utils

import (
	"crypto"
	"crypto/dsa"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenerateSignature(msg string, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) (string, error) {
	msgHash := sha256.Sum256([]byte(msg))
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, msgHash[:])
	if err != nil {
		fmt.Println(err)
	}

	return base64.StdEncoding.EncodeToString([]byte(signature)), err
}

func VerifySignature(msg string, signature string, publicKey *rsa.PublicKey, params ...dsa.Parameters) bool {
	msgHash := sha256.Sum256([]byte(msg))
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, msgHash[:], signatureBytes)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}