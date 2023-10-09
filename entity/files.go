package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Files struct {
	gorm.Model
	Files 	[]byte `json:"files" binding:"required"`

	// take the user_id from users table
	UserID  uuid.UUID `gorm:"foreignKey;type:char(36)" json:"user_id"`
}
