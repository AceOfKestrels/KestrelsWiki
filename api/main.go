package main

import (
	"api/controller/fileController"
	"api/ent"
	"api/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	_ "github.com/xiaoqidun/entps"
)

func main() {
	dbClient, err := ent.Open("sqlite3", "file:./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(dbClient *ent.Client) {
		err := dbClient.Close()
		if err != nil {

		}
	}(dbClient)

	engine := gin.Default()

	fileService := service.NewFileService()
	searchService := service.NewSearchService(fileService, dbClient)

	background := context.Background()
	ctx, cancel := context.WithTimeout(background, 5*time.Second)
	defer cancel()
	err = searchService.UpdateIndex(ctx, "../testFiles/README.md")
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
