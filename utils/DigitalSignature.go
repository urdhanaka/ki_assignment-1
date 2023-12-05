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
	"os"
	"strings"
	"time"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/revocation"
	"github.com/digitorus/pdfsign/sign"
	"github.com/google/uuid"
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

	return publicKeyBytes, err
}

func EmbedDigitalSign(fileID, userID uuid.UUID) error {
	filePath := fmt.Sprintf("uploads/%s/files/%s", userID, fileID)
	certificatePath := fmt.Sprintf("keys/certificate/%s", userID)

	privateKey, err := GetPrivateKey(userID)
	if err != nil {
		return err
	}

	certificateFileBytes, err := os.ReadFile(certificatePath)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(certificateFileBytes)
	if block == nil {
		return errors.New("No certificate found in the PEM file")
	}

	certificateBytes := block.Bytes
	rootCertificate, err := x509.ParseCertificate(certificateBytes)
	if err != nil {
		return nil
	}

	roots := x509.NewCertPool()
	roots.AddCert(rootCertificate)
	certificateChain, err := rootCertificate.Verify(x509.VerifyOptions{
		Roots: roots,
	})
	if err != nil {
		return err
	}

	outputPath := fmt.Sprintf("uploads/%s/files/%s.pdf", userID, fileID)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	fileInfo, err := inputFile.Stat()
	if err != nil {
		return err
	}

	size := fileInfo.Size()
	pdfReader, err := pdf.NewReader(inputFile, size)
	if err != nil {
		return err
	}

	err = sign.Sign(inputFile, outputFile, pdfReader, size, sign.SignData{
		Signature: sign.SignDataSignature{
			Info: sign.SignDataSignatureInfo{
				Name: "name",
				Date: time.Now().Local(),
			},
			CertType:   sign.CertificationSignature,
			DocMDPPerm: sign.AllowFillingExistingFormFieldsAndSignaturesPerms,
		},
		Signer:            privateKey,
		DigestAlgorithm:   crypto.SHA256,
		Certificate:       rootCertificate,
		CertificateChains: certificateChain,
		TSA: sign.TSA{
			URL:      "",
			Username: "",
			Password: "",
		},
		RevocationData:     revocation.InfoArchival{},
		RevocationFunction: sign.DefaultEmbedRevocationStatusFunction,
	})
	if err != nil {
		return err
	}

	return nil
}
