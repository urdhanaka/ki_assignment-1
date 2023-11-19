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
		user.PUT("/update/:id", UserController.UpdateCredentialUser)
		user.PUT("/update-identity/:id", UserController.UpdateIdentityUser)
		user.DELETE("/delete/:id", UserController.DeleteUser)
		user.GET("/decrypted/all", UserController.GetAllUserDecrypted)
		user.GET("/decrypted/:id", UserController.GetUserByIDDecrypted)
		user.GET("/public/:id", UserController.GetUserPublicKeyByID)
	}
}
