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
