package main

import (
	"api/controller/fileController"
	"api/controller/searchController"
	"api/controller/webhookController"
	"api/logger"
	params "api/parameters"
	"api/service/fileService"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/xiaoqidun/entps"
	"log"
)

func main() {
	apiPort := flag.Int("apiPort", 8080, "the port to run the api on")
	debug := flag.Bool("debug", false, "debug mode")
	contentPath := flag.String("contentPath", "../testFiles/", "the content path")
	wwwroot := flag.String("wwwroot", "wwwroot/", "the web content root path")
	flag.Parse()

	params.ApiPort = *apiPort
	params.Debug = *debug
	params.ContentPath = *contentPath
	params.WWWRoot = *wwwroot

	logger.Init()

	fileServ := fileService.New(params.ContentPath)

	err := fileServ.RebuildIndex(true)

	logger.Println(logger.API, "initializing gin engine")
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()

	fileCtrl := fileController.New(fileServ, "/api/file")
	searchCtrl := searchController.New(fileServ, "/api/search")
	webhookCtrl := webhookController.New(fileServ, "/api/webhook")

	engine.GET(fileCtrl.Path+"/*filepath", fileCtrl.GetFile)
	engine.POST(searchCtrl.Path, searchCtrl.PostSearch)
	engine.POST(webhookCtrl.WebhookEndpoint, webhookCtrl.PostWebhook)

	logger.Println(logger.API, "starting web server on port %v", params.ApiPort)
	err = engine.Run(fmt.Sprintf("localhost:%v", params.ApiPort))
	if err != nil {
		log.Fatal(err)
	}
}
