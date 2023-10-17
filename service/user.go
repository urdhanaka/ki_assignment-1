package service

import (
	"context"
	"errors"
	"fmt"
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
	CredentialUpdateUser(ctx context.Context, userDTO dto.UserCredentialUpdateDto) (entity.User, error)
	IdentityUpdateUser(ctx context.Context, userDTO dto.UserIdentityUpdateDto) (entity.User, error)
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

	// Credential
	user.Username_AES = userDTO.Username
	user.Username_RC4 = userDTO.Username
	user.Username_DEC = userDTO.Username
	user.Password_AES = userDTO.Password
	user.Password_RC4 = userDTO.Password
	user.Password_DEC = userDTO.Password

	// Identity
	user.Name_AES = userDTO.Name
	user.Name_DEC = userDTO.Name
	user.Name_RC4 = userDTO.Name
	user.Number_AES = userDTO.Number
	user.Number_DEC = userDTO.Number
	user.Number_RC4 = userDTO.Number

	// CV
	user.CV_ID = uuid.New()
	user.CV_AES = userDTO.CV.Filename
	user.CV_DEC = userDTO.CV.Filename
	user.CV_RC4 = userDTO.CV.Filename

	// ID Card
	user.ID_Card_ID = uuid.New()
	user.ID_Card_AES = userDTO.ID_Card.Filename
	user.ID_Card_DEC = userDTO.ID_Card.Filename
	user.ID_Card_RC4 = userDTO.ID_Card.Filename

	// Check file type
	if userDTO.CV.Header.Get("Content-Type") != "application/pdf" && userDTO.CV.Header.Get("Content-Type") != "image/png" && userDTO.CV.Header.Get("Content-Type") != "image/jpeg" && userDTO.CV.Header.Get("Content-Type") != "image/jpg" {
		return entity.User{}, errors.New("cv file type is not supported")
	}

	if userDTO.ID_Card.Header.Get("Content-Type") != "application/pdf" && userDTO.ID_Card.Header.Get("Content-Type") != "image/png" && userDTO.ID_Card.Header.Get("Content-Type") != "image/jpeg" && userDTO.CV.Header.Get("Content-Type") != "image/jpg" {
		return entity.User{}, errors.New("id_card file type is not supported")
	}

	// Check file name
	if userDTO.CV.Filename == "" {
		return entity.User{}, errors.New("cv is not valid")
	}

	if userDTO.ID_Card.Filename == "" {
		return entity.User{}, errors.New("id_card is not valid")
	}

	// CV upload
	cvFileName := fmt.Sprintf("%s/files/%s", user.ID, user.CV_ID)
	if err := utils.UploadFileUtility(userDTO.CV, cvFileName); err != nil {
		return entity.User{}, err
	}

	// ID Card upload
	idCardFileName := fmt.Sprintf("%s/files/%s", user.ID, user.ID_Card_ID)
	if err := utils.UploadFileUtility(userDTO.ID_Card, idCardFileName); err != nil {
		return entity.User{}, err
	}

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

func (u *userService) IdentityUpdateUser(ctx context.Context, userDto dto.UserIdentityUpdateDto) (entity.User, error) {
	var user entity.User

	user.ID = userDto.ID
	user.Name_AES = userDto.Name
	user.Name_DEC = userDto.Name
	user.Name_RC4 = userDto.Name
	user.Number_AES = userDto.Number
	user.Number_DEC = userDto.Number
	user.Number_RC4 = userDto.Number

	// WIP
	user.ID_Card_AES = userDto.ID_Card.Header.Get("")
	user.ID_Card_DEC = userDto.ID_Card.Header.Get("")
	user.ID_Card_RC4 = userDto.ID_Card.Header.Get("")
	user.CV_AES = userDto.CV.Header.Get("")
	user.CV_DEC = userDto.CV.Header.Get("")
	user.CV_RC4 = userDto.CV.Header.Get("")

	// WIP
	result, err := u.UserRepository.UpdateUser(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil

}

func (u *userService) CredentialUpdateUser(ctx context.Context, userDTO dto.UserCredentialUpdateDto) (entity.User, error) {
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

	// Credential
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

	// Identity
	decryptedNameAES, err := utils.DecryptAES(user.Name_AES)
	if err == nil {
		user.Name_AES = decryptedNameAES
	}

	decryptedNameDES, err := utils.DecryptDES(user.Name_DEC)
	if err == nil {
		user.Name_DEC = decryptedNameDES
	}

	decryptedNameRC4, err := utils.DecryptRC4(user.Name_RC4)
	if err == nil {
		user.Name_RC4 = decryptedNameRC4
	}

	decryptedNumberAES, err := utils.DecryptAES(user.Number_AES)
	if err == nil {
		user.Number_AES = decryptedNumberAES
	}

	decryptedNumberDES, err := utils.DecryptDES(user.Number_DEC)
	if err == nil {
		user.Number_DEC = decryptedNumberDES
	}

	decryptedNumberRC4, err := utils.DecryptRC4(user.Number_RC4)
	if err == nil {
		user.Number_RC4 = decryptedNumberRC4
	}

	return user, nil
}
