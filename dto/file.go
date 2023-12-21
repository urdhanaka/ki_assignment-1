package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type FileCreateDto struct {
	ID     uuid.UUID             `gorm:"primary_key" json:"id"`
	Files  *multipart.FileHeader `form:"file" binding:"required"`
	Name   string                `form:"name" binding:"required"`
	UserID uuid.UUID             `gorm:"foreignKey" json:"user_id" form:"user_id"`
}

type GetFileDto struct {
	Filename  string `form:"filename" binding:"required"`
	PublicKey string `form:"public_key" binding:"required"`
}

type GetFileByFileID struct {
	FileID    uuid.UUID `form:"file_id" binding:"required"`
	UserID    uuid.UUID `form:"user_id" binding:"required"`
	PublicKey string    `form:"public_key" binding:"required"`
}
