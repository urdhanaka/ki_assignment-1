package utils

import (
	cryptrand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randSeq(length int) string {
	init := make([]rune, length)

	for i := range init {
		init[i] = charset[rand.Intn(len(charset))]
	}

	return string(init)
}

func GenerateSecretKey() string {
	// Generate Secret Key
	secret := randSeq(32)
	return secret
}

func GenerateSecretKey8Byte() string {
	// Generate Secret Key
	secret := randSeq(8)
	return secret
}

func GenerateIV() string {
	// Generate IV
	iv := randSeq(16)
	return iv
}

func Generate8Byte() string {
	// Generate 8 Byte IV
	iv := randSeq(8)
	return iv
}

func GenerateAsymmetricKeys(id uuid.UUID) error {
	privateKey, err := rsa.GenerateKey(cryptrand.Reader, 2048)
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

func GenerateCertificates(id uuid.UUID) error {
	privateKey, err := GetPrivateKey(id)
	if err != nil {
		return err
	}

	publicKey, err := GetRSAPublicKey(id)
	if err != nil {
		return err
	}

	certificateFolder := "keys/certificate"
	if _, err = os.Stat(certificateFolder); os.IsNotExist(err) {
		if err = os.MkdirAll(certificateFolder, 0o777); err != nil {
			return err
		}
	}

	x509RootCertificate := &x509.Certificate{
		SerialNumber:          big.NewInt(2023),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(5, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	rootCertificateBytes, err := x509.CreateCertificate(cryptrand.Reader, x509RootCertificate, x509RootCertificate, publicKey, privateKey)
	if err != nil {
		return err
	}

	certificatePath := fmt.Sprintf("keys/certificate/%s.pem", id)
	certificatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: rootCertificateBytes,
	})

	err = os.WriteFile(certificatePath, certificatePEM, 0o644)
	if err != nil {
		return nil
	}

	return nil
}
