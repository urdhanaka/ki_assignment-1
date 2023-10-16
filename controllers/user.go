package controllers

import (
	"ki_assignment-1/dto"
	"ki_assignment-1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUserDecrypted(c *gin.Context)
	GetUserByIDDecrypted(c *gin.Context)
}

type userController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		UserService: userService,
	}
}

func (u *userController) RegisterUser(c *gin.Context) {
	var userDTO dto.UserCreateDto

	if err := c.ShouldBindJSON(&userDTO); err != nil {
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

func (u *userController) UpdateUser(c *gin.Context) {
	var userDTO dto.UserUpdateDto

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.UserService.UpdateUser(c, userDTO)
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
