package main

import (
	"api/controller/fileController"
	"api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	fileService := service.NewFileService()
	fileController := fileController.NewFileController(fileService)

	engine.GET("/api/file/:filepath", fileController.GetFile)

	engine.Run("localhost:8080")
}
