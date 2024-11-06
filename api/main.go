package main

import (
	"api/controller/fileController"
	"api/controller/searchController"
	"api/controller/webhookController"
	"api/logger"
	"api/service"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/xiaoqidun/entps"
	"log"
)

var ApiPort int
var Debug bool

func main() {
	apiPort := flag.Int("apiPort", 8080, "the port to run the api on")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	ApiPort = *apiPort
	Debug = *debug

	logger.Init()

	searchService := service.NewSearchService("../testFiles/")

	err := searchService.RebuildIndex(true)

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

	logger.Println(logger.API, "starting web server on port %v", ApiPort)
	err = engine.Run(fmt.Sprintf("localhost:%v", ApiPort))
	if err != nil {
		log.Fatal(err)
	}
}
