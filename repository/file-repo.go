package repository

import (
	"context"
	"ki_assignment-1/entity"

	"gorm.io/gorm"
)

type FileConnection struct {
	connection *gorm.DB
}

type FileRepository interface {
	UploadFile(ctx context.Context, file entity.Files) (entity.Files, error)
	GetAllFiles(ctx context.Context) ([]entity.Files, error)
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &FileConnection{
		connection: db,
	}
}

// User can upload file PDF/DOC/XLS/PHOTO/Videos to the database
func (db *FileConnection) UploadFile(ctx context.Context, file entity.Files) (entity.Files, error) {
	if err := db.connection.Create(&file).Error; err != nil {
		return entity.Files{}, err
	}

	return file, nil
}

// Create function for get all files users
func (db *FileConnection) GetAllFiles(ctx context.Context) ([]entity.Files, error) {
	var files []entity.Files

	if err := db.connection.Find(&files).Error; err != nil {
		return []entity.Files{}, err
	}

	return files, nil
}