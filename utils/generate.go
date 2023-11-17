package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSecretKey() string {
	// Generate Secret Key
	secret := make([]byte, 32)
	rand.Read(secret)
	return base64.StdEncoding.EncodeToString(secret)
}

func Generate8ByteSecretKey() string {
	// Generate Secret Key
	secret := make([]byte, 8)
	rand.Read(secret)
	return base64.StdEncoding.EncodeToString(secret)
}

func GenerateIV() string {
	// Generate IV
	iv := make([]byte, 16)
	rand.Read(iv)
	return base64.StdEncoding.EncodeToString(iv)
}

func Generate8Byte() string {
	// Generate 8 Byte IV
	iv := make([]byte, 8)
	rand.Read(iv)
	return base64.StdEncoding.EncodeToString(iv)
}