package repository

import (
	"context"
	"ki_assignment-1/entity"

	"gorm.io/gorm"
)

type UserConnection struct {
	connection *gorm.DB
}

type UserRepository interface {
	RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, userID uint64) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, userID uint64) (error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserConnection{
		connection: db,
	}
}

func (db *UserConnection) RegisterUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := db.connection.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (db *UserConnection) GetUserByID(ctx context.Context, userID uint64) (entity.User, error) {
	var user entity.User

	if err := db.connection.Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (db *UserConnection) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := db.connection.Save(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (db *UserConnection) DeleteUser(ctx context.Context, userID uint64) (error) {
	if err := db.connection.Where("id = ?", userID).Delete(&entity.User{}).Error; err != nil {
		return err
	}

	return nil
}