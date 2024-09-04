package fileController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileService interface {
	ReadFileContents(fileName string) (string, error)
}

type FileController struct {
	fileService FileService
}

func NewFileController(fileService FileService) *FileController {
	return &FileController{fileService: fileService}
}

func (f *FileController) GetFile(context *gin.Context) {
	filePath := "../testFiles/" + context.Param("filepath")
	if filePath == "" {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	content, error := f.fileService.ReadFileContents(filePath)
	if error != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	dto, error := NewFileDto(content)
	if error != nil {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	context.JSON(http.StatusOK, dto)
}
