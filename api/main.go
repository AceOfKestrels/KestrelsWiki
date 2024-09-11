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

	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	engine := gin.Default()

	searchService := service.NewSearchService(dbClient, "../testFiles/")

	background := context.Background()
	ctx, cancel := context.WithTimeout(background, 5*time.Second)
	defer cancel()
	err = searchService.AddFileToIndex(ctx, "README.md")
	if err != nil {
		log.Fatal(err)
	}

	controller := fileController.NewFileController(searchService, "/api/file")

	engine.GET(controller.Path+"/*filepath", controller.GetFile)

	err = engine.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
