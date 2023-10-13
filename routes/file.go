package routes

import (
	"ki_assignment-1/controllers"

	"github.com/gin-gonic/gin"
)

func FileRoutes(router *gin.Engine, FileController controllers.FileController) {
	file := router.Group("/file")
	{
		file.POST("/upload", FileController.UploadFile)
		file.GET("/all", FileController.GetAllFiles)
	}
}