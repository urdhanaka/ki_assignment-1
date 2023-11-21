package entity

import (
	"ki_assignment-1/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Identity struct {
		Name_AES    string    `json:"name_aes" binding:"required"`
		Name_RC4    string    `json:"name_rc4" binding:"required"`
		Name_DEC    string    `json:"name_dec" binding:"required"`
		Number_AES  string    `json:"number_aes" binding:"required"`
		Number_RC4  string    `json:"number_rc4" binding:"required"`
		Number_DEC  string    `json:"number_dec" binding:"required"`
		CV_ID       uuid.UUID `json:"cv_id"`
		CV_AES      string    `json:"cv_aes" binding:"required"`
		CV_RC4      string    `json:"cv_rc4" binding:"required"`
		CV_DEC      string    `json:"cv_dec" binding:"required"`
		ID_Card_ID  uuid.UUID `json:"id_card_id"`
		ID_Card_AES string    `json:"id_card_aes" binding:"required"`
		ID_Card_RC4 string    `json:"id_card_rc4" binding:"required"`
		ID_Card_DEC string    `json:"id_card_dec" binding:"required"`
	}

	Credential struct {
		Username     string `json:"username" binding:"required"`
		Username_AES string `json:"username_aes" binding:"required"`
		Username_RC4 string `json:"username_rc4" binding:"required"`
		Username_DEC string `json:"username_dec" binding:"required"`
		Password_AES string `json:"password_aes" binding:"required"`
		Password_RC4 string `json:"password_rc4" binding:"required"`
		Password_DEC string `json:"password_dec" binding:"required"`
	}

	Key struct {
		SecretKey      string `json:"secret" binding:"required"`
		IV             string `json:"iv" binding:"required"`
		SecretKey8Byte string `json:"secret_key_8_byte" binding:"required"`
		IV8Byte        string `json:"iv_8_byte" binding:"required"`
	}
)

type User struct {
	ID uuid.UUID `gorm:"primary_key;not_null;type:char(36)" json:"id"`
	Identity
	Credential
	Key

	// user has many files
	Files []Files `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" binding:"required" json:"files"`

	// user has many allowed users
	AllowedUsers []AllowedUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" binding:"required" json:"allowed_users"`

	Timestamp
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Done
	if enc, err := utils.EncryptAES([]byte(u.Username_AES), u.SecretKey, u.IV); err == nil {
		u.Username_AES = string(enc)
	}

	// Done
	if enc, err := utils.EncryptRC4([]byte(u.Username_RC4), u.SecretKey); err == nil {
		u.Username_RC4 = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Username_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.Username_DEC = string(enc)
	}

	// Done
	if enc, err := utils.EncryptAES([]byte(u.Password_AES), u.SecretKey, u.IV); err == nil {
		u.Password_AES = string(enc)
	}

	// Done
	if enc, err := utils.EncryptRC4([]byte(u.Password_RC4), u.SecretKey); err == nil {
		u.Password_RC4 = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Password_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.Password_DEC = string(enc)
	}

	// Identity
	// Name
	if enc, err := utils.EncryptAES([]byte(u.Name_AES), u.SecretKey, u.IV); err == nil {
		u.Name_AES = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Name_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.Name_DEC = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.Name_RC4), u.SecretKey); err == nil {
		u.Name_RC4 = string(enc)
	}

	// Number
	if enc, err := utils.EncryptAES([]byte(u.Number_AES), u.SecretKey, u.IV); err == nil {
		u.Number_AES = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Number_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.Number_DEC = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.Number_RC4), u.SecretKey); err == nil {
		u.Number_RC4 = string(enc)
	}

	// CV
	if enc, err := utils.EncryptAES([]byte(u.CV_AES), u.SecretKey, u.IV); err == nil {
		u.CV_AES = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.CV_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.CV_DEC = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.CV_RC4), u.SecretKey); err == nil {
		u.CV_RC4 = string(enc)
	}

	// ID_Card
	if enc, err := utils.EncryptAES([]byte(u.ID_Card_AES), u.SecretKey, u.IV); err == nil {
		u.ID_Card_AES = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.ID_Card_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.ID_Card_DEC = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.ID_Card_RC4), u.SecretKey); err == nil {
		u.ID_Card_RC4 = string(enc)
	}

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	// Done
	if enc, err := utils.EncryptAES([]byte(u.Username_AES), u.SecretKey, u.IV); err == nil {
		u.Username_AES = string(enc)
	}

	// Done
	if enc, err := utils.EncryptRC4([]byte(u.Username_RC4), u.SecretKey); err == nil {
		u.Username_RC4 = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Username_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.Username_DEC = string(enc)
	}

	// Done
	if enc, err := utils.EncryptAES([]byte(u.Password_AES), u.SecretKey, u.IV); err == nil {
		u.Password_AES = string(enc)
	}

	// Done
	if enc, err := utils.EncryptRC4([]byte(u.Password_RC4), u.SecretKey); err == nil {
		u.Password_RC4 = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Password_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.Password_DEC = string(enc)
	}

	// Identity
	// Name
	if enc, err := utils.EncryptAES([]byte(u.Name_AES), u.SecretKey, u.IV); err == nil {
		u.Name_AES = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Name_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.Name_DEC = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.Name_RC4), u.SecretKey); err == nil {
		u.Name_RC4 = string(enc)
	}

	// Number
	if enc, err := utils.EncryptAES([]byte(u.Number_AES), u.SecretKey, u.IV); err == nil {
		u.Number_AES = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.Number_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.Number_DEC = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.Number_RC4), u.SecretKey); err == nil {
		u.Number_RC4 = string(enc)
	}

	// CV
	if enc, err := utils.EncryptAES([]byte(u.CV_AES), u.SecretKey, u.IV); err == nil {
		u.CV_AES = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.CV_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.CV_DEC = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.CV_RC4), u.SecretKey); err == nil {
		u.CV_RC4 = string(enc)
	}

	// ID_Card
	if enc, err := utils.EncryptAES([]byte(u.ID_Card_AES), u.SecretKey, u.IV); err == nil {
		u.ID_Card_AES = string(enc)
	}

	if enc, err := utils.EncryptDES([]byte(u.ID_Card_DEC), u.SecretKey8Byte, u.IV8Byte); err == nil {
		u.ID_Card_DEC = string(enc)
	}

	if enc, err := utils.EncryptRC4([]byte(u.ID_Card_RC4), u.SecretKey); err == nil {
		u.ID_Card_RC4 = string(enc)
	}

	return nil
}
