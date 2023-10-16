package service

import (
	"context"
	"ki_assignment-1/dto"
	"ki_assignment-1/entity"
	"ki_assignment-1/repository"
	"ki_assignment-1/utils"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDTO dto.UserCreateDto) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, userID string) (entity.User, error)
	UpdateUser(ctx context.Context, userDTO dto.UserUpdateDto) (entity.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetAllUserDecrypted(ctx context.Context) ([]entity.User, error)
	GetUserByIDDecrypted(ctx context.Context, userID string) (entity.User, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		UserRepository: userRepo,
	}
}

func (u *userService) RegisterUser(ctx context.Context, userDTO dto.UserCreateDto) (entity.User, error) {
	var user entity.User

	user.ID = uuid.New()
	user.Username_AES = userDTO.Username
	user.Username_RC4 = userDTO.Username
	user.Username_DEC = userDTO.Username
	user.Password_AES = userDTO.Password
	user.Password_RC4 = userDTO.Password
	user.Password_DEC = userDTO.Password

	result, err := u.UserRepository.RegisterUser(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (u *userService) GetAllUser(ctx context.Context) ([]entity.User, error) {
	result, err := u.UserRepository.GetAllUser(ctx)
	if err != nil {
		return []entity.User{}, err
	}

	return result, nil
}

func (u *userService) DeleteUser(ctx context.Context, userID string) error {
	err := u.UserRepository.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	result, err := u.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (u *userService) UpdateUser(ctx context.Context, userDTO dto.UserUpdateDto) (entity.User, error) {
	var user entity.User

	user.ID = userDTO.ID
	user.Username_AES = userDTO.Username
	user.Username_RC4 = userDTO.Username
	user.Username_DEC = userDTO.Username
	user.Password_AES = userDTO.Password
	user.Password_RC4 = userDTO.Password
	user.Password_DEC = userDTO.Password

	result, err := u.UserRepository.UpdateUser(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (u *userService) GetAllUserDecrypted(ctx context.Context) ([]entity.User, error) {
    users, err := u.UserRepository.GetAllUser(ctx)
    if err != nil {
        return nil, err
    }

    for i := range users {
        decryptedUsernameAES, err := utils.DecryptAES(users[i].Username_AES)
        if err == nil {
            users[i].Username_AES = decryptedUsernameAES
        }

        decryptedPasswordAES, err := utils.DecryptAES(users[i].Password_AES)
        if err == nil {
            users[i].Password_AES = decryptedPasswordAES
        }

        decryptedUsernameDES, err := utils.DecryptDES(users[i].Username_DEC)
        if err == nil {
            users[i].Username_DEC = decryptedUsernameDES
        }

        decryptedPasswordDES, err := utils.DecryptDES(users[i].Password_DEC)
        if err == nil {
            users[i].Password_DEC = decryptedPasswordDES
        }

        decryptedUsernameRC4, err := utils.DecryptRC4(users[i].Username_RC4)
        if err == nil {
            users[i].Username_RC4 = decryptedUsernameRC4
        }

        decryptedPasswordRC4, err := utils.DecryptRC4(users[i].Password_RC4)
        if err == nil {
            users[i].Password_RC4 = decryptedPasswordRC4
        }
    }

    return users, nil
}

func (u *userService) GetUserByIDDecrypted(ctx context.Context, userID string) (entity.User, error) {
    user, err := u.GetUserByID(ctx, userID)
    if err != nil {
        return entity.User{}, err
    }

    decryptedUsernameAES, err := utils.DecryptAES(user.Username_AES)
        if err == nil {
            user.Username_AES = decryptedUsernameAES
        }

        decryptedPasswordAES, err := utils.DecryptAES(user.Password_AES)
        if err == nil {
            user.Password_AES = decryptedPasswordAES
        }

        decryptedUsernameDES, err := utils.DecryptDES(user.Username_DEC)
        if err == nil {
            user.Username_DEC = decryptedUsernameDES
        }

        decryptedPasswordDES, err := utils.DecryptDES(user.Password_DEC)
        if err == nil {
            user.Password_DEC = decryptedPasswordDES
        }

        decryptedUsernameRC4, err := utils.DecryptRC4(user.Username_RC4)
        if err == nil {
            user.Username_RC4 = decryptedUsernameRC4
        }

        decryptedPasswordRC4, err := utils.DecryptRC4(user.Password_RC4)
        if err == nil {
            user.Password_RC4 = decryptedPasswordRC4
        }

    return user, nil
}