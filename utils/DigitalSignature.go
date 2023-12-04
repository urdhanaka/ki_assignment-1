package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func GenerateSignature(msg []byte, privateKey *rsa.PrivateKey) (string, error) {
	msgHash := sha256.Sum256(msg)
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, msgHash[:])
	if err != nil {
		fmt.Println(err)
	}

	return base64.StdEncoding.EncodeToString([]byte(signature)), err
}

func VerifySignature(msg []byte, signature string, publicKey *rsa.PublicKey) bool {
	msgHash := sha256.Sum256(msg)
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

func ParsePublicKeyFromPEM(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
			return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
			return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
			return pub, nil
	default:
			return nil, errors.New("public key is not of type RSA")
	}
}