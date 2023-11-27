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
	GetFileByUserID(c *gin.Context)
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

	filePath, err := f.FileService.GetFilePath(c, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	DecryptedFileContent, err := f.FileService.GetFile(c, filePath, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(DecryptedFileContent)
	_, err = os.Stat(DecryptedFileContent)
	if os.IsNotExist(err) {
		res := gin.H{
			"status":  "error",
			"message": "file not found",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.File(DecryptedFileContent)

	err = os.Remove(DecryptedFileContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// Get File by user id
func (f *fileController) GetFileByUserID(c *gin.Context) {
	userID := c.Query("user_id")

	fmt.Println(userID)

	files, err := f.FileService.GetFileByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}
