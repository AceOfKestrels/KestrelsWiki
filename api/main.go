package main

import (
	"api/controller/fileController"
	"api/controller/searchController"
	"api/controller/webhookController"
	"api/db/ent"
	"api/logger"
	"api/service"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/xiaoqidun/entps"
	"log"
)

func main() {
	dbPath := flag.String("dbPath", "./data.db", "path to the database file")
	apiPort := flag.Int("apiPort", 8080, "the port to run the api on")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	logger.Init()

	logger.Println(logger.DB, "attempting to open database at %v", dbPath)
	dbClient, err := ent.Open("sqlite3", "file:"+*dbPath)
	if err != nil {
		log.Fatal(err)
	}
	if *debug {
		dbClient = dbClient.Debug()
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

	logger.Println(logger.INIT, "updating file index")
	err = searchService.UpdateIndex()
	if err != nil {
		logger.Println(logger.INIT, err.Error())
	}

	logger.Println(logger.API, "initializing gin engine")
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()

	fileCtrl := fileController.NewFileController(searchService, "/api/file")
	searchCtrl := searchController.NewSearchController(searchService, "/api/search")
	webhookCtrl := webhookController.NewWebhookController(searchService, "/api/webhook")

	engine.GET(fileCtrl.Path+"/*filepath", fileCtrl.GetFile)
	engine.POST(searchCtrl.Path, searchCtrl.PostSearch)
	engine.POST(webhookCtrl.WebhookEndpoint, webhookCtrl.PostWebhook)

	logger.Println(logger.API, "starting web server on port %v", *apiPort)
	err = engine.Run(fmt.Sprintf("localhost:%v", *apiPort))
	if err != nil {
		log.Fatal(err)
	}
}
