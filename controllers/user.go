package controllers

import (
	"fmt"
	"ki_assignment-1/dto"
	"ki_assignment-1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateCredentialUser(c *gin.Context)
	UpdateIdentityUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUserDecrypted(c *gin.Context)
	GetUserByIDDecrypted(c *gin.Context)
	GetUserPublicKeyByUsername(c *gin.Context)
	GetUserPrivateKeyByUsername(c *gin.Context)
	GetUserSymmetricKeyByUsername(c *gin.Context)
}

type userController struct {
	UserService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		UserService: userService,
		jwtService:  jwtService,
	}
}

func (u *userController) RegisterUser(c *gin.Context) {
	var userDTO dto.UserCreateDto

	if err := c.ShouldBind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.UserService.RegisterUser(c, userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) LoginUser(c *gin.Context) {
	var userDTO dto.UserLoginDto

	if err := c.ShouldBind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verifyUser, _ := u.UserService.VerifyUser(c.Request.Context(), userDTO)
	if !verifyUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	user, err := u.UserService.GetUserByUsername(c.Request.Context(), userDTO.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token := u.jwtService.GenerateToken(user.ID)

	fmt.Println(token)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  verifyUser,
	},
	)
}

func (u *userController) GetAllUser(c *gin.Context) {
	users, err := u.UserService.GetAllUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *userController) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := u.UserService.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) UpdateCredentialUser(c *gin.Context) {
	var userDTO dto.UserCredentialUpdateDto

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.UserService.CredentialUpdateUser(c, userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) UpdateIdentityUser(c *gin.Context) {
	var userDTO dto.UserIdentityUpdateDto

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.UserService.IdentityUpdateUser(c, userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := u.UserService.DeleteUser(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (u *userController) GetAllUserDecrypted(c *gin.Context) {
	users, err := u.UserService.GetAllUserDecrypted(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *userController) GetUserByIDDecrypted(c *gin.Context) {
	id := c.Param("id")

	user, err := u.UserService.GetUserByIDDecrypted(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) GetUserPublicKeyByUsername(c *gin.Context) {
	token := c.MustGet("token").(string)
	fmt.Println(token)
	userID, err := u.jwtService.FindUserIDByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	publicKey, err := u.UserService.GetUserPublicKeyByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, publicKey)
}

func (u *userController) GetUserPrivateKeyByUsername(c *gin.Context) {
	token := c.MustGet("token").(string)

	userID, err := u.jwtService.FindUserIDByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	privateKey, err := u.UserService.GetUserPrivateKeyByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, privateKey)
}

func (u *userController) GetUserSymmetricKeyByUsername(c *gin.Context) {
	// Check if the user is allowed to access other user data
	token := c.MustGet("token").(string)

	userIDRequesting, err := u.jwtService.FindUserIDByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the requesting username from the url
	username := c.Query("username")
	userRequested, err := u.UserService.GetUserByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the requesting user's public key
	publicKey, err := u.UserService.GetUserPublicKeyByID(c, userIDRequesting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists in allowed users list
	// allowedUser, err := u.UserService.GetAllowedUserByID(c, userIDRequesting, userRequested.ID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// Find the symmetric key
	symmetricKey, err := u.UserService.GetUserSymmetricKeyByID(userRequested.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encrypt Secret Key and IV with public key

	encryptedKey, err := u.UserService.EncryptSecretKey(symmetricKey, publicKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"secret_key":        userRequested.SecretKey,
	// 	"secret_key_8bytes": userRequested.SecretKey8Byte,
	// 	"iv":                userRequested.IV,
	// 	"iv_8bytes":         userRequested.IV8Byte,
	// })

	c.JSON(http.StatusOK, encryptedKey)
}

