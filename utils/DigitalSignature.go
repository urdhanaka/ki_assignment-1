package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"strings"
)

func GenerateSignature(msg []byte, privateKey *rsa.PrivateKey) (string, error) {
	msgHash := sha256.Sum256(msg)
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, msgHash[:])
	if err != nil {
		fmt.Println(err)
	}

	return base64.StdEncoding.EncodeToString([]byte(signature)), err
}

func VerifySignature(msg []byte, signature string, publickey []byte) bool {
	msgHash := sha256.Sum256(msg)
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println(err)
		return false
	}

	publicKey, err := x509.ParsePKCS1PublicKey(publickey)
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

func ParsePublicKeyFromString(publicKey string) ([]byte, error) {
	publicKey = strings.ReplaceAll(publicKey, " ", "+")

	padding := len(publicKey) % 4
	if padding > 0 {
		publicKey += strings.Repeat("=", 4-padding)
	}

	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		fmt.Printf("Error decoding base64: %v\n", err)
		fmt.Println("Base64 String:", publicKey)
		return nil, err
	}

	return publicKeyBytes, nil

}