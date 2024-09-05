package main

import (
	"api/controller/fileController"
	"api/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	fileService := service.NewFileService()
	searchService := service.SearchServiceImpl{}
	err := searchService.UpdateIndex("../testFiles/README.md")
	if err != nil {
		log.Fatal(err)
	}

	controller := fileController.NewFileController(fileService, "/api/file", "../testFiles")

	engine.GET(controller.Path+"/*filepath", controller.GetFile)

	err = engine.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
