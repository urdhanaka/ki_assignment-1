package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func GenerateEncryptedHash(msg []byte, publicKey *rsa.PublicKey) (string, error) {
	msgHash := sha256.Sum256(msg)

	encryptedMsgHash, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, msgHash[:], nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedMsgHash), nil
}

// unused
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

	encryptedMsgHash, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, msgHash[:], nil)
	if err != nil {
		fmt.Println(err)
		return false
	}

	encryptedMsgHashString := base64.StdEncoding.EncodeToString(encryptedMsgHash)

	return encryptedMsgHashString == string(signature)
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

// func EmbedDigitalSign(fileID, userID uuid.UUID, secretKey string, ivKey string) error {
// 	tempFilePath := fmt.Sprintf("uploads/%s/files/tmp.pdf", userID)
// 	filePath := fmt.Sprintf("uploads/%s/files/%s", userID, fileID)
// 	certificatePath := fmt.Sprintf("keys/certificate/%s.pem", userID)
//
// 	privateKey, err := GetPrivateKey(userID)
// 	if err != nil {
// 		return err
// 	}
//
// 	certificateFileBytes, err := os.ReadFile(certificatePath)
// 	if err != nil {
// 		return err
// 	}
//
// 	block, _ := pem.Decode(certificateFileBytes)
// 	if block == nil {
// 		return errors.New("no certificate found in the PEM file")
// 	}
//
// 	certificateBytes := block.Bytes
// 	rootCertificate, err := x509.ParseCertificate(certificateBytes)
// 	if err != nil {
// 		return nil
// 	}
//
// 	roots := x509.NewCertPool()
// 	roots.AddCert(rootCertificate)
// 	certificateChain, err := rootCertificate.Verify(x509.VerifyOptions{
// 		Roots: roots,
// 	})
// 	if err != nil {
// 		return err
// 	}
//
// 	outputPath := fmt.Sprintf("uploads/%s/files/res.pdf", userID)
// 	outputFile, err := os.Create(outputPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer outputFile.Close()
//
// 	inputFile, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer inputFile.Close()
//
// 	fileData, err := io.ReadAll(inputFile)
// 	if err != nil {
// 		return err
// 	}
//
// 	decryptedFileData, err := DecryptAESFile(fileData, secretKey, ivKey)
// 	if err != nil {
// 		return err
// 	}
//
// 	err = os.WriteFile(tempFilePath, decryptedFileData, 0666)
// 	if err != nil {
// 		return err
// 	}
//
// 	decryptedFile, err := os.Open(tempFilePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer decryptedFile.Close()
//
// 	fileInfo, err := decryptedFile.Stat()
// 	if err != nil {
// 		return err
// 	}
//
// 	size := fileInfo.Size()
// 	pdfReader, err := pdf.NewReader(decryptedFile, size)
// 	if err != nil {
// 		return err
// 	}
//
// 	err = sign.Sign(inputFile, outputFile, pdfReader, size, sign.SignData{
// 		Signature: sign.SignDataSignature{
// 			Info: sign.SignDataSignatureInfo{
// 				Name: "name",
// 				Date: time.Now().Local(),
// 			},
// 			CertType:   sign.CertificationSignature,
// 			DocMDPPerm: sign.AllowFillingExistingFormFieldsAndSignaturesPerms,
// 		},
// 		Signer:            privateKey,
// 		DigestAlgorithm:   crypto.SHA256,
// 		Certificate:       rootCertificate,
// 		CertificateChains: certificateChain,
// 		TSA: sign.TSA{
// 			URL:      "",
// 			Username: "",
// 			Password: "",
// 		},
// 		RevocationData:     revocation.InfoArchival{},
// 		RevocationFunction: sign.DefaultEmbedRevocationStatusFunction,
// 	})
// 	if err != nil {
// 		return err
// 	}
//
// 	decryptedFile.Close()
// 	os.Remove(tempFilePath)
//
// 	return nil
// }
