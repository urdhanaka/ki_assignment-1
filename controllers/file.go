package controllers

import (
	"fmt"
	"ki_assignment-1/dto"
	"ki_assignment-1/service"
	"ki_assignment-1/utils"
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
	jwtService  service.JWTService
	UserService service.UserService
	FileService service.FileService
}

func NewFileController(fileService service.FileService, jwts service.JWTService, userService service.UserService) FileController {
	return &fileController{
		jwtService:  jwts,
		FileService: fileService,
		UserService: userService,
	}
}

func (f *fileController) UploadFile(c *gin.Context) {
	var fileDTO dto.FileCreateDto

	token := c.MustGet("token").(string)
	userID, err := f.jwtService.FindUserIDByToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&fileDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	fileDTO.UserID = userID
	file, err := f.FileService.UploadFile(c, fileDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	c.JSON(http.StatusOK, file)
}

func (f *fileController) GetFile(c *gin.Context) {
	token := c.MustGet("token").(string)
	userID, err := f.jwtService.FindUserIDByToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fileDTO dto.GetFileDto
	if err := c.ShouldBind(&fileDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("dto.Publickey:  " + fileDTO.PublicKey)

	rsaPublicKey, err := utils.ParsePublicKeyFromString(fileDTO.PublicKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filePath, err := f.FileService.GetFilePath(c, fileDTO.Filename, userID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	DecryptedFileContent, err := f.FileService.GetFile(c, filePath, fileDTO.Filename)
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

	signature, err := f.FileService.GetFileSignature(c, userID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verify the signature of the file
	isVerified := utils.VerifySignature([]byte(DecryptedFileContent), signature, rsaPublicKey)

	// c.File(DecryptedFileContent)
	c.JSON(http.StatusOK, gin.H{
		"file": DecryptedFileContent,
		"is_signature_verified": isVerified,
	})

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
