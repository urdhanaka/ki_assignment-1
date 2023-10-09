package controllers

import (
	"ki_assignment-1/dto"
	"ki_assignment-1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	// GetUserByID(c *gin.Context)
	// UpdateUser(c *gin.Context)
	// DeleteUser(c *gin.Context)
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
	var userDto dto.UserCreateDto

	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.UserService.RegisterUser(c, userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}