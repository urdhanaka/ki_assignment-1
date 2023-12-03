package routes

import (
	"ki_assignment-1/controllers"
	"ki_assignment-1/middleware"
	"ki_assignment-1/service"

	"github.com/gin-gonic/gin"
)

func FileRoutes(router *gin.Engine, FileController controllers.FileController, jwtService service.JWTService) {
	file := router.Group("/file")
	{
		file.POST("/upload", middleware.Authenticate(jwtService), FileController.UploadFile)
		file.GET("/detail", FileController.GetFile)
		// Get file by user id with param user id
		file.GET("/user", FileController.GetFileByUserID)
	}
}
