package controllers

import (
	"fmt"
	"net/http"

	"ki_assignment-1/dto"
	"ki_assignment-1/service"

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
	GetUserPublicKeyByID(c *gin.Context)
}

type userController struct {
	UserService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		UserService: userService,
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

	fmt.Println(user.ID, user.Name_AES, user.Username_AES, user.Password_AES)

	token := u.jwtService.GenerateToken(user.ID)

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

func (u *userController) GetUserPublicKeyByID(c *gin.Context) {
	userID := c.Param("id")

	publicKey, err := u.UserService.GetUserPublicKeyByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, publicKey)
}
