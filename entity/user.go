package entity

import (
	"ki_assignment-1/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"primary_key;not_null;type:char(36)" json:"id"`
	Username_AES string    `json:"username_aes" binding:"required"`
	Username_RC4 string    `json:"username_rc4" binding:"required"`
	Username_DEC string    `json:"username_dec" binding:"required"`
	Password_AES string    `json:"password_aes" binding:"required"`
	Password_RC4 string    `json:"password_rc4" binding:"required"`
	Password_DEC string    `json:"password_dec" binding:"required"`

	// user has many files
	Files []Files `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" binding:"required" json:"files"`

	Timestamp
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if enc, err := utils.EncryptAES(u.Username_AES); err == nil {
		u.Username_AES = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.Username_RC4), []byte(utils.GetEnv("KEY"))); err == nil {
		u.Username_RC4 = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Username_DEC), []byte(utils.GetEnv("KEY"))); err == nil {
		u.Username_DEC = string(enc)
	}

	if enc, err := utils.EncryptAES(u.Password_AES); err == nil {
		u.Password_AES = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.Password_RC4), []byte(utils.GetEnv("KEY"))); err == nil {
		u.Password_RC4 = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Password_DEC), []byte(utils.GetEnv("KEY"))); err == nil {
		u.Password_DEC = string(enc)
	}

	return nil
}
