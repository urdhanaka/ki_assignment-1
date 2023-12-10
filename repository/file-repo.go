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
	GetFile(ctx context.Context, fileID string) (entity.Files, error)
	GetFileByUserID(ctx context.Context, userID string) ([]entity.Files, error)
	GetFileID(ctx context.Context, filename string) (string, error)
	GetUserIDfromFilename(ctx context.Context, filename string) (string, error)
	GetFileByName(ctx context.Context, fileID string) (entity.Files, error)
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

// Get File
func (db *FileConnection) GetFile(ctx context.Context, fileID string) (entity.Files, error) {
	var file entity.Files

	if err := db.connection.Where("id = ?", fileID).First(&file).Error; err != nil {
		return entity.Files{}, err
	}

	return file, nil
}

// Get File by user id
func (db *FileConnection) GetFileByUserID(ctx context.Context, userID string) ([]entity.Files, error) {
	var files []entity.Files

	if err := db.connection.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&files).Error; err != nil {
		return []entity.Files{}, err
	}

	return files, nil
}

// Get File ID from Filename
func (db *FileConnection) GetFileID(ctx context.Context, filename string) (string, error) {
	var file entity.Files
	if err := db.connection.Where("name = ?", filename).Select("id").First(&file).Error; err != nil {
		return "", err
	}
	return file.ID.String(), nil
}

// Get User ID from Filename
func (db *FileConnection) GetUserIDfromFilename(ctx context.Context, filename string) (string, error) {
	var file entity.Files
	if err := db.connection.Where("name = ?", filename).Select("user_id").First(&file).Error; err != nil {
		return "", err
	}
	return file.UserID.String(), nil
}

func (db *FileConnection) GetFileByName(ctx context.Context, filename string) (entity.Files, error) {
	var file entity.Files

	if err := db.connection.Where("name = ?", filename).First(&file).Error; err != nil {
		return entity.Files{}, err
	}

	return file, nil
}