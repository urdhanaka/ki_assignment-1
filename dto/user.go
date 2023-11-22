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

type UserRequestDataDTO struct {
	Username          string `json:"username"`
	EncyptedSecretKey string `json:"encrypted_secret_key"`
	EncryptedIvKey    string `json:"encrypted_iv_key"`
}

type RequestedUserSymmetricKeysDTO struct {
	SecretKey      string `json:"secret_key"`
	IV             string `json:"iv"`
	SecretKey8Byte string `json:"secret_key_8_byte"`
	IV8Byte        string `json:"iv_8_byte"`
}

type EncryptedRequestedUserSymmetricKeysDTO struct {
	EncryptedSecretKey string `json:"encrypted_secret_key"`
	EncryptedIVKey     string `json:"encrypted_iv"`
}

type DecryptedRequestedUserSymmetricKeysDTO struct {
	DecryptedSecretKey string `json:"decrypted_secret_key"`
	DecryptedIVKey     string `json:"decrypted_iv"`
}
