package utils

import (
	"crypto/rand"
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
