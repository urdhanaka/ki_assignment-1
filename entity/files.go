package entity

import (
	"gorm.io/gorm"
)

type Files struct {
	gorm.Model
	Files []byte `json:"files" binding:"required"`
	// user has many files
	UserID uint64 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" binding:"required" json:"user_id"`
}