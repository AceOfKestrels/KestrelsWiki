package fileController

import (
	"api/logger"
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type FileService interface {
	ContentPath() string
	GetFileDto(fileContent string) (models.FileDTO, error)
}

type Controller struct {
	fileService   FileService
	Path          string
	PathParamName string
}

func New(fileService FileService, path string, pathParamName string) *Controller {
	return &Controller{fileService: fileService, Path: path, PathParamName: pathParamName}
}

func (f *Controller) GetFile(c *gin.Context) {
	filePath := c.Param(f.PathParamName)
	if filePath == "" {
		c.Status(http.StatusBadRequest)
		logger.Println(logger.API, "error: filepath was empty")
		return
	}

	if strings.Contains(filePath, ".") {
		c.File(f.fileService.ContentPath() + filePath)
		c.Status(http.StatusOK)
		return
	}

	dto, err := f.fileService.GetFileDto(filePath[1:] + ".md")
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
