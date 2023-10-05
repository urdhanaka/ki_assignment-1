package repository

import (
	"context"
	"ki_assignment-1/entity"
	"ki_assignment-1/utils"

	"gorm.io/gorm"
)

type UserConnection struct {
	connection *gorm.DB
}

type UserRepository interface {
	// functional
	InsertUser(ctx context.Context, username string, password string) error
	GetUserByID(ctx context.Context, userID uint64) (entity.User, error)
	// GetUserByUsername(ctx context.Context, username string, method string) (entity.User, error)
}

func (db *UserConnection) InsertUser(ctx context.Context, username string, password string) error {
	var encryptedUser entity.User

	// AES Encryption starts here
	if encryptedAES, err := utils.EncryptAES([]byte(username), []byte(utils.GetEnv("KEY"))); err == nil {
		encryptedUser.Username_AES = string(encryptedAES)
	}

	if encryptedAES, err := utils.EncryptAES([]byte(password), []byte(utils.GetEnv("KEY"))); err == nil {
		encryptedUser.Password_AES = string(encryptedAES)
	}

	// RC4 Encryption starts here
	if encryptedRC4, err := utils.EncryptRC4([]byte(username), []byte(utils.GetEnv("KEY"))); err == nil {
		encryptedUser.Username_AES = string(encryptedRC4)
	}

	if encryptedRC4, err := utils.EncryptRC4([]byte(username), []byte(utils.GetEnv("KEY"))); err == nil {
		encryptedUser.Username_AES = string(encryptedRC4)
	}

	// DES Encryption starts here
	if encryptedDES, err := utils.EncryptRC4([]byte(username), []byte(utils.GetEnv("KEY"))); err == nil {
		encryptedUser.Username_AES = string(encryptedDES)
	}

	if encryptedDES, err := utils.EncryptRC4([]byte(username), []byte(utils.GetEnv("KEY"))); err == nil {
		encryptedUser.Username_AES = string(encryptedDES)
	}

	// Insert to database
	if err := db.connection.Create(&encryptedUser).Error; err != nil {
		return err
	}

	return nil
}

func (db *UserConnection) GetUserByID(ctx context.Context, userID uint64) (entity.User, error) {
	var user entity.User

	if err := db.connection.Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// func (db *UserConnection) GetUserByUsername(ctx context.Context, username string, method string) (entity.User, error) {
// 	var user entity.User
// 	key := []byte(utils.GetEnv("KEY"))

// 	switch method {
// 	case "AES":

// 	}

// }

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserConnection{
		connection: db,
	}
}
