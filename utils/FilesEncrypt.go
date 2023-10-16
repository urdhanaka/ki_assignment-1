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

	err = os.WriteFile(filePath, fileData, 0666)
	if err != nil {
		return err
	}

	return nil
}