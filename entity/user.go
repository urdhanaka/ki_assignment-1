package entity

import (
	// "ki_assignment-1/utils"
	// "gopkg.in/mail.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username_AES string `json:"username_aes" binding:"required"`
	Username_RC4 string `json:"username_rc4" binding:"required"`
	Username_DEC string `json:"username_dec" binding:"required"`
	Password_AES string `json:"password_aes" binding:"required"`
	Password_RC4 string `json:"password_rc4" binding:"required"`
	Password_DEC string `json:"password_dec" binding:"required"`
	Files 			 []Files `gorm:"foreignKey:UserID" json:"files" binding:"required"`
}

func (User) TableName() string {
	return "users"
}