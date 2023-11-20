package service

import (
	"context"
	"errors"
	"fmt"

	"ki_assignment-1/dto"
	"ki_assignment-1/entity"
	"ki_assignment-1/repository"
	"ki_assignment-1/utils"

	"github.com/google/uuid"
)

type FileService interface {
	UploadFile(ctx context.Context, fileDTO dto.FileCreateDto) (entity.Files, error)
	GetAllFiles(ctx context.Context) ([]entity.Files, error)
	GetFilePath(ctx context.Context, filename string) (string, error)
	GetFile(ctx context.Context, filePath string, username string) (string, error)
	GetFileByUserID(ctx context.Context, userID string) ([]entity.Files, error)
}

type fileService struct {
	FileRepository repository.FileRepository
	UserRepository repository.UserRepository
}

func NewFileService(fileRepo repository.FileRepository, userRepo repository.UserRepository) FileService {
	return &fileService{
		FileRepository: fileRepo,
		UserRepository: userRepo,
	}
}

func (f *fileService) UploadFile(ctx context.Context, fileDTO dto.FileCreateDto) (entity.Files, error) {
	var file entity.Files

	user, err := f.UserRepository.GetUserByID(ctx, fileDTO.UserID)
	if err != nil {
		return entity.Files{}, err
	}

	Files_AES, err := utils.EncryptAES([]byte(fileDTO.Files.Filename), user.SecretKey, user.IV)
	if err != nil {
		return entity.Files{}, err
	}
	Files_RC4, err := utils.EncryptAES([]byte(fileDTO.Files.Filename), user.SecretKey, user.IV)
	if err != nil {
		return entity.Files{}, err
	}
	Files_DES, err := utils.EncryptAES([]byte(fileDTO.Files.Filename), user.SecretKey, user.IV)
	if err != nil {
		return entity.Files{}, err
	}

	file.ID = uuid.New()
	file.Name = fileDTO.Name
	file.Files_AES = Files_AES
	file.Files_RC4 = Files_RC4
	file.Files_DEC = Files_DES
	file.UserID, _ = uuid.Parse(fileDTO.UserID)

	// Check file type
	if fileDTO.Files.Header.Get("Content-Type") != "application/pdf" && fileDTO.Files.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" && fileDTO.Files.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" && fileDTO.Files.Header.Get("Content-Type") != "image/jpeg" && fileDTO.Files.Header.Get("Content-Type") != "image/png" && fileDTO.Files.Header.Get("Content-Type") != "video/mp4" {
		return entity.Files{}, errors.New("file type is not supported")
	}

	// Check file size
	if fileDTO.Files.Size > 1000000 {
		return entity.Files{}, errors.New("file size is too large")
	}

	// Check file name
	if fileDTO.Files.Filename == "" {
		return entity.Files{}, errors.New("file name is not valid")
	}

	// Save the files to the uploads folder
	fileName := fmt.Sprintf("%s/files/%s", file.UserID, file.ID)
	if err := utils.UploadFileUtility(fileDTO.Files, fileName, user.SecretKey, user.IV); err != nil {
		return entity.Files{}, err
	}

	result, err := f.FileRepository.UploadFile(ctx, file)
	if err != nil {
		return entity.Files{}, err
	}

	return result, nil
}

func (f *fileService) GetAllFiles(ctx context.Context) ([]entity.Files, error) {
	result, err := f.FileRepository.GetAllFiles(ctx)
	if err != nil {
		return []entity.Files{}, err
	}

	return result, nil
}

func (f *fileService) GetFilePath(ctx context.Context, filename string) (string, error) {
	userID, err := f.FileRepository.GetUserIDfromFilename(ctx, filename)
	if err != nil {
		return filename, err
	}

	user, err := f.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return "", err
	}

	encryptedFilenameAES, err := utils.EncryptAES([]byte(filename), user.SecretKey, user.IV)
	if err == nil {
		filename = string(encryptedFilenameAES)
	}

	fileID, err := f.FileRepository.GetFileID(ctx, filename)
	if err != nil {
		return filename, err
	}

	result := fmt.Sprintf("uploads/%s/files/%s", userID, fileID)

	return result, nil
}

// Get File from repository
func (f *fileService) GetFile(ctx context.Context, filePath string, username string) (string, error) {
	user, err := f.UserRepository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	res, err := utils.GetFileUtility(filePath, user.SecretKey, user.IV)
	if err != nil {
		return "", err
	}

	return res, nil
}

// Get File by User id
func (f *fileService) GetFileByUserID(ctx context.Context, userID string) ([]entity.Files, error) {
	result, err := f.FileRepository.GetFileByUserID(ctx, userID)
	if err != nil {
		return []entity.Files{}, err
	}

	return result, nil
}
