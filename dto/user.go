package dto

import "github.com/google/uuid"

type UserCreateDto struct {
	ID 		   uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type UserUpdateDto struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}