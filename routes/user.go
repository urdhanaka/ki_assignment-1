package routes

import (
	"ki_assignment-1/controllers"
	"ki_assignment-1/middleware"
	"ki_assignment-1/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, UserController controllers.UserController, jwtService service.JWTService) {
	user := router.Group("/user")
	{
		user.POST("/register", UserController.RegisterUser)
		user.POST("/login", UserController.LoginUser)
		user.GET("/all", middleware.Authenticate(jwtService), UserController.GetAllUser)
		user.GET("/:id", middleware.Authenticate(jwtService), UserController.GetUserByID)
		user.PUT("/update/:id", middleware.Authenticate(jwtService), UserController.UpdateCredentialUser)
		user.PUT("/update-identity/:id", middleware.Authenticate(jwtService), UserController.UpdateIdentityUser)
		user.DELETE("/delete/:id", middleware.Authenticate(jwtService), UserController.DeleteUser)
		user.GET("/decrypted/all", middleware.Authenticate(jwtService), UserController.GetAllUserDecrypted)
		user.GET("/decrypted/:id", middleware.Authenticate(jwtService), UserController.GetUserByIDDecrypted)
		user.GET("/public", middleware.Authenticate(jwtService), UserController.GetUserPublicKeyByUsername)
		user.GET("/private", middleware.Authenticate(jwtService), UserController.GetUserPrivateKeyByUsername)
		user.GET("/symmetric", middleware.Authenticate(jwtService), UserController.GetUserSymmetricKeyByUsername)
	}
}
