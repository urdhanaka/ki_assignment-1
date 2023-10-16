package entity

import (
	"ki_assignment-1/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Files struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Files_AES string    `json:"files_aes" binding:"required"`
	Files_RC4 string    `json:"files_rc4" binding:"required"`
	Files_DEC string    `json:"files_dec" binding:"required"`

	// take the user_id from users table
	UserID uuid.UUID `gorm:"foreignKey;type:char(36)" json:"user_id"`
	Timestamp
}

func (Files) TableName() string {
	return "files"
}

func (f *Files) BeforeCreate(tx *gorm.DB) error {
	// Done
	if enc, err := utils.EncryptAES(f.Files_AES); err == nil {
		f.Files_AES = string(enc)
	}

	// Done
	if enc, err := utils.EncryptRC4(f.Files_RC4); err == nil {
		f.Files_RC4 = string(enc)
	}

	if enc, err := utils.EncryptDES(f.Files_DEC); err == nil {
		f.Files_DEC = string(enc)
	}

	return nil
}
