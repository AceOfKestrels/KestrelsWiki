package main

import (
	"api/controller/fileController"
	"api/ent"
	"api/logger"
	"api/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"

	_ "github.com/xiaoqidun/entps"
)

func main() {
	dbPath := "./data.db"
	apiPort := 8080

	log.SetOutput(os.Stdout)

	logger.Println(logger.DB, "attempting to open database at %v", dbPath)
	dbClient, err := ent.Open("sqlite3", "file:"+dbPath)
	dbClient = dbClient.Debug()
	if err != nil {
		log.Fatal(err)
	}
	defer func(dbClient *ent.Client) {
		logger.Println(logger.DB, "closing database")
		err := dbClient.Close()
		if err != nil {

		}
	}(dbClient)

	logger.Println(logger.DB, "updating database schema")
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	searchService := service.NewSearchService(dbClient, "../testFiles/")

	background := context.Background()
	ctx, cancel := context.WithTimeout(background, 5*time.Second)
	defer cancel()
	logger.Println(logger.INIT, "updating file index")
	err = searchService.UpdateIndex(ctx)
	if err != nil {
		logger.Println(logger.INIT, err.Error())
	}

	logger.Println(logger.API, "initializing gin engine")
	engine := gin.Default()

	controller := fileController.NewFileController(searchService, "/api/file")

	engine.GET(controller.Path+"/*filepath", controller.GetFile)

	logger.Println(logger.API, "starting web server on port %v", apiPort)
	err = engine.Run(fmt.Sprintf("localhost:%v", apiPort))
	if err != nil {
		log.Fatal(err)
	}
}
