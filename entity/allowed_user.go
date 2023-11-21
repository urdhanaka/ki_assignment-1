package entity

import "github.com/google/uuid"

type AllowedUser struct {
	ID            uuid.UUID `gorm:"primary_key;not_null;type:char(36)" json:"id"`
	AllowedUserID uuid.UUID `gorm:"foreignKey" json:"allowed_user_id"`
	UserID        uuid.UUID `gorm:"foreignKey" json:"user_id"`
}
