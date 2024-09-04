package fileController

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileService interface {
	ReadFileContents(fileName string) (string, error)
	GetFileDto(fileContent string) (models.FileDTO, error)
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

	content, err := f.fileService.ReadFileContents(filePath)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	dto, err := f.fileService.GetFileDto(content)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if dto.Meta.MirrorOf != "" {
		context.Redirect(http.StatusPermanentRedirect, dto.Meta.MirrorOf)
		return
	}

	context.JSON(http.StatusOK, dto)
}
