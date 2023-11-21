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
	VerifyUser(ctx context.Context, userDTO dto.UserLoginDto) (bool, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, userID string) (entity.User, error)
	CredentialUpdateUser(ctx context.Context, userDTO dto.UserCredentialUpdateDto) (entity.User, error)
	IdentityUpdateUser(ctx context.Context, userDTO dto.UserIdentityUpdateDto) (entity.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetAllUserDecrypted(ctx context.Context) ([]entity.User, error)
	GetUserByIDDecrypted(ctx context.Context, userID string) (entity.User, error)
	GetUserPublicKeyByID(ctx context.Context, id uuid.UUID) (string, error)
	GetUserPrivateKeyByID(ctx context.Context, id uuid.UUID) (string, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	GetAllowedUserByID(ctx context.Context, userID uuid.UUID, allowedUserID uuid.UUID) (entity.AllowedUser, error)
	GetUserSymmetricKeyByID(userID uuid.UUID) (string, error)
	EncryptSecretKey(symmetricKey string, publicKey string) (string, error)
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
	user.Username = userDTO.Username
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

	// Generate Secret Key (Symmetric Key)
	user.SecretKey = utils.GenerateSecretKey()
	user.SecretKey8Byte = utils.GenerateSecretKey8Byte()
	user.IV = utils.GenerateIV()
	user.IV8Byte = utils.Generate8Byte()

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
	if err := utils.UploadFileUtility(userDTO.CV, cvFileName, user.SecretKey, user.IV); err != nil {
		return entity.User{}, err
	}

	// ID Card upload
	idCardFileName := fmt.Sprintf("%s/files/%s", user.ID, user.ID_Card_ID)
	if err := utils.UploadFileUtility(userDTO.ID_Card, idCardFileName, user.SecretKey, user.IV); err != nil {
		return entity.User{}, err
	}

	err := utils.GenerateAsymmetricKeys(user.ID)
	if err != nil {
		return entity.User{}, nil
	}

	result, err := u.UserRepository.RegisterUser(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (u *userService) VerifyUser(ctx context.Context, userDTO dto.UserLoginDto) (bool, error) {
	user, err := u.UserRepository.GetUserByUsername(userDTO.Username)
	if err != nil {
		return false, err
	}

	encryptedPassword, err := utils.EncryptAES([]byte(userDTO.Password), user.SecretKey, user.IV)
	if err != nil {
		return false, err
	}

	if user.Password_AES != encryptedPassword {
		return false, errors.New("password is not valid")
	}

	return true, nil
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
		// AES encryption
		decryptedNameAES, err := utils.DecryptAES(users[i].Name_AES, users[i].SecretKey, users[i].IV)
		if err == nil {
			users[i].Name_AES = decryptedNameAES
		}
		decryptedNumberAES, err := utils.DecryptAES(users[i].Number_AES, users[i].SecretKey, users[i].IV)
		if err == nil {
			users[i].Number_AES = decryptedNumberAES
		}
		decryptedCVAES, err := utils.DecryptAES(users[i].CV_AES, users[i].SecretKey, users[i].IV)
		if err == nil {
			users[i].CV_AES = decryptedCVAES
		}
		decryptedIDCardAES, err := utils.DecryptAES(users[i].ID_Card_AES, users[i].SecretKey, users[i].IV)
		if err == nil {
			users[i].ID_Card_AES = decryptedIDCardAES
		}
		decryptedUsernameAES, err := utils.DecryptAES(users[i].Username_AES, users[i].SecretKey, users[i].IV)
		if err == nil {
			users[i].Username_AES = decryptedUsernameAES
		}
		decryptedPasswordAES, err := utils.DecryptAES(users[i].Password_AES, users[i].SecretKey, users[i].IV)
		if err == nil {
			users[i].Password_AES = decryptedPasswordAES
		}

		// DES
		decryptedNameDES, err := utils.DecryptDES(users[i].Name_DEC, users[i].SecretKey8Byte, users[i].IV8Byte)
		if err == nil {
			users[i].Name_DEC = decryptedNameDES
		}
		decryptedNumberDES, err := utils.DecryptDES(users[i].Number_DEC, users[i].SecretKey8Byte, users[i].IV8Byte)
		if err == nil {
			users[i].Number_DEC = decryptedNumberDES
		}
		decryptedCVDES, err := utils.DecryptDES(users[i].CV_DEC, users[i].SecretKey8Byte, users[i].IV8Byte)
		if err == nil {
			users[i].CV_DEC = decryptedCVDES
		}
		decryptedIDCardDES, err := utils.DecryptDES(users[i].ID_Card_DEC, users[i].SecretKey8Byte, users[i].IV8Byte)
		if err == nil {
			users[i].ID_Card_DEC = decryptedIDCardDES
		}
		decryptedUsernameDES, err := utils.DecryptDES(users[i].Username_DEC, users[i].SecretKey8Byte, users[i].IV8Byte)
		if err == nil {
			users[i].Username_DEC = decryptedUsernameDES
		}
		decryptedPasswordDES, err := utils.DecryptDES(users[i].Password_DEC, users[i].SecretKey8Byte, users[i].IV8Byte)
		if err == nil {
			users[i].Password_DEC = decryptedPasswordDES
		}

		// RC4
		decryptedNameRC4, err := utils.DecryptRC4(users[i].Name_RC4, users[i].SecretKey)
		if err == nil {
			users[i].Name_RC4 = decryptedNameRC4
		}
		decryptedNumberRC4, err := utils.DecryptRC4(users[i].Number_RC4, users[i].SecretKey)
		if err == nil {
			users[i].Number_RC4 = decryptedNumberRC4
		}
		decryptedCVRC4, err := utils.DecryptRC4(users[i].CV_RC4, users[i].SecretKey)
		if err == nil {
			users[i].CV_RC4 = decryptedCVRC4
		}
		decryptedIDCardRC4, err := utils.DecryptRC4(users[i].ID_Card_RC4, users[i].SecretKey)
		if err == nil {
			users[i].ID_Card_RC4 = decryptedIDCardRC4
		}
		decryptedUsernameRC4, err := utils.DecryptRC4(users[i].Username_RC4, users[i].SecretKey)
		if err == nil {
			users[i].Username_RC4 = decryptedUsernameRC4
		}
		decryptedPasswordRC4, err := utils.DecryptRC4(users[i].Password_RC4, users[i].SecretKey)
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
	// AES encryption
	decryptedNameAES, err := utils.DecryptAES(user.Name_AES, user.SecretKey, user.IV)
	if err == nil {
		user.Name_AES = decryptedNameAES
	}
	decryptedNumberAES, err := utils.DecryptAES(user.Number_AES, user.SecretKey, user.IV)
	if err == nil {
		user.Number_AES = decryptedNumberAES
	}
	decryptedCVAES, err := utils.DecryptAES(user.CV_AES, user.SecretKey, user.IV)
	if err == nil {
		user.CV_AES = decryptedCVAES
	}
	decryptedIDCardAES, err := utils.DecryptAES(user.ID_Card_AES, user.SecretKey, user.IV)
	if err == nil {
		user.ID_Card_AES = decryptedIDCardAES
	}
	decryptedUsernameAES, err := utils.DecryptAES(user.Username_AES, user.SecretKey, user.IV)
	if err == nil {
		user.Username_AES = decryptedUsernameAES
	}
	decryptedPasswordAES, err := utils.DecryptAES(user.Password_AES, user.SecretKey, user.IV)
	if err == nil {
		user.Password_AES = decryptedPasswordAES
	}

	// DES
	decryptedNameDES, err := utils.DecryptDES(user.Name_DEC, user.SecretKey8Byte, user.IV8Byte)
	if err == nil {
		user.Name_DEC = decryptedNameDES
	}
	decryptedNumberDES, err := utils.DecryptDES(user.Number_DEC, user.SecretKey8Byte, user.IV8Byte)
	if err == nil {
		user.Number_DEC = decryptedNumberDES
	}
	decryptedCVDES, err := utils.DecryptDES(user.CV_DEC, user.SecretKey8Byte, user.IV8Byte)
	if err == nil {
		user.CV_DEC = decryptedCVDES
	}
	decryptedIDCardDES, err := utils.DecryptDES(user.ID_Card_DEC, user.SecretKey8Byte, user.IV8Byte)
	if err == nil {
		user.ID_Card_DEC = decryptedIDCardDES
	}
	decryptedUsernameDES, err := utils.DecryptDES(user.Username_DEC, user.SecretKey, user.IV)
	if err == nil {
		user.Username_DEC = decryptedUsernameDES
	}
	decryptedPasswordDES, err := utils.DecryptDES(user.Password_DEC, user.SecretKey, user.IV)
	if err == nil {
		user.Password_DEC = decryptedPasswordDES
	}

	// RC4
	decryptedNameRC4, err := utils.DecryptRC4(user.Name_RC4, user.SecretKey)
	if err == nil {
		user.Name_RC4 = decryptedNameRC4
	}
	decryptedNumberRC4, err := utils.DecryptRC4(user.Number_RC4, user.SecretKey)
	if err == nil {
		user.Number_RC4 = decryptedNumberRC4
	}
	decryptedCVRC4, err := utils.DecryptRC4(user.CV_RC4, user.SecretKey)
	if err == nil {
		user.CV_RC4 = decryptedCVRC4
	}
	decryptedIDCardRC4, err := utils.DecryptRC4(user.ID_Card_RC4, user.SecretKey)
	if err == nil {
		user.ID_Card_RC4 = decryptedIDCardRC4
	}
	decryptedUsernameRC4, err := utils.DecryptRC4(user.Username_RC4, user.SecretKey)
	if err == nil {
		user.Username_RC4 = decryptedUsernameRC4
	}

	decryptedPasswordRC4, err := utils.DecryptRC4(user.Password_RC4, user.SecretKey)
	if err == nil {
		user.Password_RC4 = decryptedPasswordRC4
	}

	return user, nil
}

func (u *userService) GetUserPublicKeyByID(ctx context.Context, id uuid.UUID) (string, error) {
	user, err := u.UserRepository.GetUserByID(ctx, id.String())
	if err != nil {
		return "", err
	}

	res, err := utils.GetPublicKey(user.ID)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (u *userService) GetUserPrivateKeyByID(ctx context.Context, id uuid.UUID) (string, error) {
	user, err := u.UserRepository.GetUserByID(ctx, id.String())
	if err != nil {
		return "", err
	}

	res, err := utils.GetPrivateKey(user.ID)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (u *userService) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	result, err := u.UserRepository.GetUserByUsername(username)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (u *userService) GetAllowedUserByID(ctx context.Context, userID uuid.UUID, allowedUserID uuid.UUID) (entity.AllowedUser, error) {
	result, err := u.UserRepository.GetAllowedUserByID(userID, allowedUserID)
	if err != nil {
		return entity.AllowedUser{}, err
	}

	return result, nil
}

func (u *userService) GetUserSymmetricKeyByID(userID uuid.UUID) (string, error) {
	user, err := u.UserRepository.GetUserByID(context.Background(), userID.String())
	if err != nil {
		return "", err
	}

	return user.SecretKey, nil
}

func (u *userService) EncryptSecretKey(symmetricKey string, publicKey string) (string, error) {
	encryptedKey, err := utils.EncryptSymmetricKey(symmetricKey, publicKey)
	if err != nil {
		return "", err
	}

	return encryptedKey, nil
}

// func (u* userService) GetPrivateData(ctx context.Context, privateKey string) (string, error) {
// 	privateKey, err := []byte(privateKey), nil
// 	if err != nil {
// 		return "", err
// 	}

	


// }
