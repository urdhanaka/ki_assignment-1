package controllers

import (
	"ki_assignment-1/dto"
	"ki_assignment-1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController interface {
	UploadFile(c *gin.Context)
	GetAllFiles(c *gin.Context)
}

type fileController struct {
	FileService service.FileService
}

func NewFileController(fileService service.FileService) FileController {
	return &fileController{
		FileService: fileService,
	}
}

func (f *fileController) UploadFile(c *gin.Context) {
	var fileDTO dto.FileCreateDto

	if err := c.ShouldBind(&fileDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	file, err := f.FileService.UploadFile(c, fileDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	c.JSON(http.StatusOK, file)
}

func (f *fileController) GetAllFiles(c *gin.Context) {
	files, err := f.FileService.GetAllFiles(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}