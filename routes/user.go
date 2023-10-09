package routes

import (
	"ki_assignment-1/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, UserController controllers.UserController) {
	user := router.Group("/user")
	{
		user.POST("/register", UserController.RegisterUser)
		user.GET("/all", UserController.GetAllUser)
		// user.GET("/:id", userController.GetUserByID)
		// user.PUT("/:id", userController.UpdateUser)
		// user.DELETE("/:id", userController.DeleteUser)
	}
}
