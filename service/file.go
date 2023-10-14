package service

import (
	"context"
	"errors"
	"ki_assignment-1/dto"
	"ki_assignment-1/entity"
	"ki_assignment-1/repository"

	"github.com/google/uuid"
)

type FileService interface {
	UploadFile(ctx context.Context, fileDTO dto.FileCreateDto) (entity.Files, error)
	GetAllFiles(ctx context.Context) ([]entity.Files, error)
}

type fileService struct {
	FileRpository repository.FileRepository
}

func NewFileService(fileRepo repository.FileRepository) FileService {
	return &fileService{
		FileRpository: fileRepo,
	}
}

func (f *fileService) UploadFile(ctx context.Context, fileDTO dto.FileCreateDto) (entity.Files, error) {
	var file entity.Files

	file.ID = uuid.New()
	file.Name = fileDTO.Name
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

	// Save the files to the

	result, err := f.FileRpository.UploadFile(ctx, file)
	if err != nil {
		return entity.Files{}, err
	}

	return result, nil
}

func (f *fileService) GetAllFiles(ctx context.Context) ([]entity.Files, error) {
	result, err := f.FileRpository.GetAllFiles(ctx)
	if err != nil {
		return []entity.Files{}, err
	}

	return result, nil
}
