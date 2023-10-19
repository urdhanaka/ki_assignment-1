package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func UploadFileUtility(file *multipart.FileHeader, path string) error {
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

	encryptedFileData, err := EncryptAESFile(fileData)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, encryptedFileData, 0666)
	if err != nil {
		return err
	}

	return nil
}

func GetFileUtility(path string) ([]byte, error) {
	// Get it from /uploads/userid/files/[fileid]
	// parts := strings.Split(path, "/")

	// fileId := parts[2]
	fileId := path
	userID := "96052b2b-02a8-4747-8210-6d4820804dd5"

	filePath := fmt.Sprintf("uploads/%s/files/%s", userID, fileId)

	file, err := os.Open(filePath)
	if err != nil {
		return []byte{}, err
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		return []byte{}, err
	}

	return fileData, nil
}