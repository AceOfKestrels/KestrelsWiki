package searchController

import (
	"api/logger"
	"api/models"
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
	c.JSON(http.StatusOK, results)
}
