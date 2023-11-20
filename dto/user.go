package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UserCreateDto struct {
	ID       uuid.UUID             `gorm:"primary_key" json:"id"`
	Name     string                `json:"name" binding:"required"`
	Number   string                `json:"number" binding:"required"`
	CV       *multipart.FileHeader `json:"cv" binding:"required"`
	ID_Card  *multipart.FileHeader `json:"id_card" binding:"required"`
	Username string                `json:"username" binding:"required"`
	Password string                `json:"password" binding:"required"`
}

type UserLoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCredentialUpdateDto struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type UserIdentityUpdateDto struct {
	ID      uuid.UUID             `json:"id" binding:"required"`
	Name    string                `json:"name" binding:"required"`
	Number  string                `json:"number" binding:"required"`
	CV      *multipart.FileHeader `json:"cv" binding:"required"`
	ID_Card *multipart.FileHeader `json:"id_card" binding:"required"`
}
