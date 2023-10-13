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
	GetAllUser(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, userID uint64) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, userID uint64) (error)

	CalculateAESAlgorithmTime(start int64, end int64) uint64
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

func (db *UserConnection) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	if err := db.connection.Find(&users).Error; err != nil {
		return []entity.User{}, err
	}

	return users, nil
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

func (db *UserConnection) CalculateAESAlgorithmTime(start int64, end int64) uint64 {
	var timeElapsed uint64 = uint64(end - start)
	return timeElapsed
}