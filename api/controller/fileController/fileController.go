package fileController

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type FileService interface {
	GetFileDto(fileContent string) (models.FileDTO, error)
}

type FileController struct {
	fileService FileService
	Path        string
	contentPath string
}

func NewFileController(fileService FileService, path string, contentPath string) *FileController {
	return &FileController{fileService: fileService, Path: path, contentPath: contentPath}
}

func (f *FileController) GetFile(context *gin.Context) {
	filePath := context.Param("filepath")
	if filePath == "" {
		context.Status(http.StatusBadRequest)
		return
	}

	if strings.Contains(filePath, ".") {
		context.File(f.contentPath + filePath)
		context.Status(http.StatusOK)
		return
	}
	fullPath := f.contentPath + filePath + ".md"

	dto, err := f.fileService.GetFileDto(fullPath)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	if dto.MirrorOf != "" {
		redirect := strings.ToLower(dto.MirrorOf)
		if strings.HasSuffix(redirect, ".md") {
			redirect = redirect[0 : len(redirect)-3]
		}
		context.Redirect(http.StatusPermanentRedirect, redirect)
		return
	}

	dto.Path = f.Path + filePath + ".md"

	context.JSON(http.StatusOK, dto)
}
