package searchController

import (
	"api/logger"
	"api/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type SearchService interface {
	ContentPath() string
	SearchFiles(ctx context.Context, search models.SearchContext) ([]models.FileDTO, error)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	results, err := s.searchService.SearchFiles(ctx, search)
	if err != nil {
		c.Status(http.StatusNotFound)
		logger.Println(logger.API, "error: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}
