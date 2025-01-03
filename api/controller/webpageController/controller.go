package webpageController

import (
	params "api/parameters"
	"github.com/gin-gonic/gin"
)

type FileService interface {
	Exists(path string) bool
	GetArticle(path string) (string, error)
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

	fullPath, err := ctrl.getFullPath(filePath)
	if err != nil {
		c.File(params.WWWRoot + notFoundFile)
		return
	}

	c.File(fullPath)
}

func (ctrl *Controller) getFullPath(path string) (fullPath string, err error) {
	fullPath = params.WWWRoot + path
	if ctrl.fileService.Exists(fullPath) {
		return
	}

	fullPath = params.ContentPath + path
	if ctrl.fileService.Exists(fullPath) {
		return
	}

	fullPath = params.WWWRoot + articleFile
	if _, err = ctrl.fileService.GetArticle(path); err != nil {
		fullPath = ""
	}

	return
}
