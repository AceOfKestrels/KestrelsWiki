package main

import (
	"api/controller/fileController"
	"api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	fileService := service.NewFileService()
	controller := fileController.NewFileController(fileService, "/api/file", "../testFiles")

	engine.GET(controller.Path+"/*filepath", controller.GetFile)

	engine.Run("localhost:8080")
}
