package searchController

import (
	"api/logger"
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchService interface {
	ContentPath() string
	SearchFiles(search models.SearchContext) []models.FileDTO
}

type SearchController struct {
	searchService SearchService
	Path          string
}

func NewSearchController(searchService SearchService, path string) *SearchController {
	return &SearchController{searchService: searchService, Path: path}
}

func (s *SearchController) PostSearch(c *gin.Context) {
	var search models.SearchContext
	err := c.BindJSON(&search)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logger.Println(logger.API, "error: %v", err.Error())
		return
	}

	results := s.searchService.SearchFiles(search)
	c.JSON(http.StatusOK, results)
}
