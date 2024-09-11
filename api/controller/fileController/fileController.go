package fileController

import (
	"api/logger"
	"api/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type SearchService interface {
	ContentPath() string
	GetFileDto(ctx context.Context, fileContent string) (models.FileDTO, error)
}

type FileController struct {
	searchService SearchService
	Path          string
}

func NewFileController(searchService SearchService, path string) *FileController {
	return &FileController{searchService: searchService, Path: path}
}

func (f *FileController) GetFile(c *gin.Context) {
	filePath := c.Param("filepath")
	if filePath == "" {
		c.Status(http.StatusBadRequest)
		logger.Println(logger.API, "error: filepath was empty")
		return
	}

	if strings.Contains(filePath, ".") {
		c.File(f.searchService.ContentPath() + filePath)
		c.Status(http.StatusOK)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	dto, err := f.searchService.GetFileDto(ctx, filePath[1:]+".md")
	if err != nil {
		c.Status(http.StatusNotFound)
		logger.Println(logger.API, "error finding file: %v", err)
		return
	}

	if dto.MirrorOf != "" {
		redirect := strings.ToLower(dto.MirrorOf)
		if strings.HasSuffix(redirect, ".md") {
			redirect = redirect[0 : len(redirect)-3]
		}
		c.Redirect(http.StatusPermanentRedirect, redirect)
		return
	}

	dto.Path = f.Path + "/" + dto.Path

	c.JSON(http.StatusOK, dto)
}
