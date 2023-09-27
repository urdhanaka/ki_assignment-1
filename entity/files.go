package entity

type Files struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	Files []byte `json:"files" binding:"required"`

	UserID uint64 `gorm:"foreignKey" json:"user_id"`
	User   *User  `gorm:"constraint:OnUpdate:CASCADE:OnDelete:CASCADE;" json:"user,omitempty"`
}
