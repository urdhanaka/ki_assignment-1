package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func UploadFileUtility(file *multipart.FileHeader, path string, secretKeyParam string, ivParam string) error {
	parts := strings.Split(path, "/")

	fileId := parts[2]
	directoryPath := fmt.Sprintf("uploads/%s/%s", parts[0], parts[1])

	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		if err := os.MkdirAll(directoryPath, 0777); err != nil {
			return err
		}
	}

	filePath := fmt.Sprintf("%s/%s", directoryPath, fileId)

	uploadedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadedFile.Close()

	fileData, err := io.ReadAll(uploadedFile)
	if err != nil {
		return err
	}

	encryptedFileData, err := EncryptAESFile(fileData, secretKeyParam, ivParam)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, encryptedFileData, 0666)
	if err != nil {
		return err
	}

	return nil
}

func GetFileUtility(path string, secretKeyParam []byte, ivParam []byte) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "path", err
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		return "path", err
	}

	decryptedFileData, err := DecryptAESFile(fileData, string(secretKeyParam), string(ivParam))
	if err != nil {
		return "", err
	}

	pathSplit := strings.Split(path, "/")

	tempPath := fmt.Sprintf("%s/%s/temp", pathSplit[0], pathSplit[1])

	err = os.WriteFile(tempPath, decryptedFileData, 0666)
	if err != nil {
		return "", err
	}

	return tempPath, nil
}
