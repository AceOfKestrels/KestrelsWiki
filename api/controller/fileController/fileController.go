package fileController

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type SearchService interface {
	ContentPath() string
	GetFileDto(fileContent string) (models.FileDTO, error)
}

type FileController struct {
	searchService SearchService
	Path          string
}

func NewFileController(searchService SearchService, path string) *FileController {
	return &FileController{searchService: searchService, Path: path}
}

func (f *FileController) GetFile(context *gin.Context) {
	filePath := context.Param("filepath")
	if filePath == "" {
		context.Status(http.StatusBadRequest)
		return
	}

	if strings.Contains(filePath, ".") {
		context.File(f.searchService.ContentPath() + filePath)
		context.Status(http.StatusOK)
		return
	}

	dto, err := f.searchService.GetFileDto(filePath + ".md")
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
