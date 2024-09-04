package fileController

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type FileService interface {
	ReadFileContents(fileName string) (string, error)
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

	content, err := f.fileService.ReadFileContents(fullPath)
	if err != nil {
		context.Status(http.StatusNotFound)
		return
	}

	dto, err := f.fileService.GetFileDto(content)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	if dto.Meta.MirrorOf != "" {
		redirect := strings.ToLower(dto.Meta.MirrorOf)
		if strings.HasSuffix(redirect, ".md") {
			redirect = redirect[0 : len(redirect)-3]
		}
		context.Redirect(http.StatusPermanentRedirect, redirect)
		return
	}

	dto.Meta.Path = f.Path + filePath + ".md"

	context.JSON(http.StatusOK, dto)
}
