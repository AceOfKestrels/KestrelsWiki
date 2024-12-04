package main

import (
	"api/controller/fileController"
	"api/controller/searchController"
	"api/controller/webhookController"
	"api/controller/webpageController"
	"api/logger"
	params "api/parameters"
	"api/service/fileService"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	apiPort := flag.Int("apiPort", 8080, "the port to run the api on")
	debug := flag.Bool("debug", false, "debug mode")
	contentPath := flag.String("contentPath", "../testFiles/", "the content path")
	wwwroot := flag.String("wwwroot", "wwwroot", "the web content root path")
	logPath := flag.String("logPath", "", "the path where log files are saved. leave black to disable logging to file")
	flag.Parse()

	params.ApiPort = *apiPort
	params.Debug = *debug
	params.ContentPath = *contentPath
	params.WWWRoot = *wwwroot
	params.LogPath = *logPath

	logger.Init()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
		cleanup()
	}()

	fileServ := fileService.New(params.ContentPath)

	err := fileServ.RebuildIndex(true)

	logger.Println(logger.API, "initializing gin engine")
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()

	webCtrl := webpageController.New(fileServ)
	fileCtrl := fileController.New(fileServ, "/api/file", "path")
	searchCtrl := searchController.New(fileServ, "/api/search")
	webhookCtrl := webhookController.New(fileServ, "/api/webhook")

	engine.GET(fileCtrl.Path+"/*path", fileCtrl.GetFile)
	engine.POST(searchCtrl.Path, searchCtrl.PostSearch)
	engine.POST(webhookCtrl.WebhookEndpoint, webhookCtrl.PostWebhook)
	engine.NoRoute(webCtrl.GetPage)

	logger.Println(logger.API, "starting web server on port %v", params.ApiPort)
	err = engine.Run(fmt.Sprintf("localhost:%v", params.ApiPort))
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second)
		if logger.LogFile != nil {
			_ = logger.LogFile.Sync()
		}
	}
}

func cleanup() {
	if logger.LogFile != nil {
		_ = logger.LogFile.Close()
		_ = logger.RenameOldLogFile(params.LogPath)
	}
}
