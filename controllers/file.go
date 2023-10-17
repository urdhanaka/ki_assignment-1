package controllers

import (
	"fmt"
	"ki_assignment-1/dto"
	"ki_assignment-1/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileController interface {
	UploadFile(c *gin.Context)
	GetFile(c *gin.Context)
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

func (f *fileController) GetFile(c *gin.Context) {
	fileName := c.Query("filename")
	encryptionMethod := c.Query("encryption_method")

	fileDecrypt, err := f.FileService.DecryptFile(fileName, encryptionMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(encryptionMethod, fileDecrypt)

	// Get the uuid from db
	fileID := c.Query("file_id")

	filePath := fmt.Sprintf("uploads/96052b2b-02a8-4747-8210-6d4820804dd5/files/%s", fileID)
	fmt.Println(filePath)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		res := gin.H{
			"status":  "error",
			"message": "file not found",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.File(filePath)
}