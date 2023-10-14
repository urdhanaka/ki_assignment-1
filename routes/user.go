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
		user.GET("/:id", UserController.GetUserByID)
		user.PUT("/update/:id", UserController.UpdateUser)
		user.DELETE("/delete/:id", UserController.DeleteUser)
	}
}
