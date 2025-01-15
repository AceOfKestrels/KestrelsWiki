package searchController

import (
	"api/logger"
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileService interface {
	ContentPath() string
	SearchFiles(search models.SearchContext) []models.FileDTO
}

type Controller struct {
	fileService FileService
	Path        string
}

func New(fileService FileService, path string) *Controller {
	return &Controller{fileService: fileService, Path: path}
}

func (s *Controller) PostSearch(c *gin.Context) {
	var search models.SearchContext
	err := c.BindJSON(&search)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logger.Println(logger.API, "error: %v", err.Error())
		return
	}

	results := s.fileService.SearchFiles(search)
	resultTemplate := "<ul class=\"searchResults\">"

	for _, result := range results {
		path := result.Path[0 : len(result.Path)-3]
		resultTemplate += fmt.Sprintf("<li class=\"searchResult\"><a href=\"%s\">%s</a></li>", path, result.Title)
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(resultTemplate))
}
