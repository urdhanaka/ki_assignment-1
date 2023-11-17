package entity

import (
	"github.com/google/uuid"
)

type Files struct {
	ID        uuid.UUID `gorm:"primary_key;not_null;type:char(36)" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Files_AES string    `json:"files_aes" binding:"required"`
	Files_RC4 string    `json:"files_rc4" binding:"required"`
	Files_DEC string    `json:"files_dec" binding:"required"`

	// take the user_id from users table
	UserID uuid.UUID `gorm:"foreignKey;type:char(36)" json:"user_id"`
	Timestamp
}