package webpageController

import (
	"api/models"
	params "api/parameters"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type FileService interface {
	Exists(path string) bool
	GetFileDto(filePath string) (models.FileDTO, error)
}

type Controller struct {
	fileService   FileService
	PathParamName string
}

func New(fileService FileService) *Controller {
	return &Controller{fileService: fileService}
}

const indexFile = "/index.html"
const articleFile = "/article.html"
const notFoundFile = "/not-found.html"

func (ctrl *Controller) GetPage(c *gin.Context) {
	filePath := c.Request.URL.Path

	if len(filePath) == 0 {
		c.File(indexFile)
		return
	}

	fullPath, redirect := ctrl.getFullPath(filePath)
	if len(fullPath) == 0 {
		c.File(params.WWWRoot + notFoundFile)
		return
	}
	if redirect {
		fullPath = strings.ToLower(fullPath)
		if strings.HasSuffix(fullPath, ".md") {
			fullPath = fullPath[0 : len(fullPath)-3]
		}
		c.Redirect(http.StatusPermanentRedirect, fullPath)
		return
	}

	c.File(fullPath)
}

func (ctrl *Controller) getFullPath(path string) (fullPath string, redirect bool) {
	redirect = false
	fullPath = params.WWWRoot + path
	if ctrl.fileService.Exists(fullPath) {
		return
	}

	fullPath = params.ContentPath + path
	if ctrl.fileService.Exists(fullPath) {
		return
	}

	dto, err := ctrl.fileService.GetFileDto(path)
	if err != nil {
		return "", false
	}
	if len(dto.MirrorOf) != 0 {
		return dto.MirrorOf, true
	}

	return params.WWWRoot + articleFile, false
}
